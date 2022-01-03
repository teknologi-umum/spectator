package main

import (
	"context"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type LogPayload struct {
	AccessToken string    `json:"access_token" msgpack:"access_token"`
	Data        []LogData `json:"data" msgpack:"data"`
}

type LogData struct {
	RequestID   string            `json:"request_id" msgpack:"request_id"`
	Application string            `json:"application" msgpack:"application"`
	Message     string            `json:"message" msgpack:"message"`
	Body        map[string]string `json:"body" msgpack:"body"`
	Level       string            `json:"level" msgpack:"level"`
	Environment string            `json:"environment" msgpack:"environment"`
	Language    string            `json:"language" msgpack:"language"`
	Timestamp   time.Time         `json:"timestamp" msgpack:"timestamp"`
}

func (d *Dependency) writeIntoLog(ctx context.Context, p []LogData) error {
	// recovering here in case of any error
	defer func() {
		r := recover()
		if r != nil {
			log.Printf("panic: %v", r.(error))
		}
	}()

	writeAPI := d.DB.WriteAPI(d.Org, "log")

	for index, payload := range p {
		// write defaults first
		if payload.Environment == "" {
			payload.Environment = "UNSET"
		}

		if payload.Level == "" {
			payload.Level = "DEBUG"
		}

		if payload.Timestamp.IsZero() {
			payload.Timestamp = time.Now()
		}

		// convert body into json as influxdb doesnt accept map
		bodyBytes, err := json.Marshal(payload.Body)
		if err != nil {
			return fmt.Errorf("payload index: %d, marshalling json: %v", index, err)
		}

		point := influxdb2.NewPoint(
			payload.Level,
			map[string]string{
				"request_id":  payload.RequestID,
				"application": payload.Application,
				"environment": payload.Environment,
			},
			map[string]interface{}{
				"language": payload.Language,
				"message":  hex.EncodeToString([]byte(payload.Message)),
				"body":     hex.EncodeToString(bodyBytes),
			},
			payload.Timestamp,
		)

		writeAPI.WritePoint(point)
	}

	writeAPI.Flush()
	err := writeAPI.Errors()
	if err != nil {
		return fmt.Errorf("writing into log: %v", err)
	}
	return nil
}

func buildQuery(q queries) string {
	var str strings.Builder
	str.WriteString("from(bucket: \"log\")\n")
	// range query
	str.WriteString("|> range(")
	if !q.TimeFrom.IsZero() {
		str.WriteString("start: " + strconv.FormatInt(q.TimeFrom.Unix(), 10))
	} else {
		str.WriteString("start: 0")
	}

	if !q.TimeTo.IsZero() {
		str.WriteString(", stop: " + strconv.FormatInt(q.TimeTo.Unix(), 10))
	}

	str.WriteString(")\n")

	str.WriteString("|> sort(columns: [\"_time\"])\n")
	str.WriteString("|> group(columns: [\"request_id\"])\n")
	thereIsDataToBeFiltered := q.Level != "" || q.Application != "" || q.RequestID != ""
	if thereIsDataToBeFiltered {
		str.WriteString(`|> filter(fn: (r) => `)
	}

	var filtered []string

	if q.Level != "" {
		filtered = append(filtered, `r["_measurement"] == "`+q.Level+`"`)
	}

	if q.Application != "" {
		filtered = append(filtered, `r["application"] == "`+q.Application+`"`)
	}

	if q.RequestID != "" {
		filtered = append(filtered, `r["request_id"] == "`+q.RequestID+`"`)
	}

	str.WriteString(strings.Join(filtered, " and "))

	if thereIsDataToBeFiltered {
		str.WriteString(")\n")
	}
	return str.String()
}

func (d *Dependency) fetchLog(ctx context.Context, query queries) ([]LogData, error) {
	queryAPI := d.DB.QueryAPI(d.Org)
	// build query for influx
	queryStr := buildQuery(query)

	rows, err := queryAPI.Query(ctx, queryStr)
	if err != nil {
		return []LogData{}, err
	}

	defer rows.Close()

	var output []LogData
	var temp LogData
	var lastTableIndex int = -1
	for rows.Next() {
		unmarshaledRow, err := unmarshalInfluxRow(rows.Record().String())
		if err != nil {
			return []LogData{}, err
		}

		tableStr, ok := unmarshaledRow["table"].(string)
		if !ok {
			continue
		}

		table, err := strconv.Atoi(tableStr)
		if err != nil {
			return []LogData{}, err
		}
		if table == lastTableIndex {
			switch unmarshaledRow["_field"].(string) {
			case "body":
				bodyJSON := unmarshaledRow["_value"].(string)
				bodyBytes, err := hex.DecodeString(bodyJSON)
				if err != nil {
					return []LogData{}, err
				}
				body := make(map[string]string, 100)
				err = json.Unmarshal(bodyBytes, &body)
				if err != nil {
					return []LogData{}, err
				}
				temp.Body = body
			case "language":
				temp.Language = unmarshaledRow["_value"].(string)
			case "message":
				messageBytes, err := hex.DecodeString(unmarshaledRow["_value"].(string))
				if err != nil {
					return []LogData{}, err
				}
				temp.Message = string(messageBytes)
			}
		} else {
			// clear the last temp, but check if its less than zero
			if lastTableIndex >= 0 {
				output = append(output, temp)
			}
			// create a new one
			temp.Application = unmarshaledRow["application"].(string)
			temp.Environment = unmarshaledRow["environment"].(string)
			temp.Level = unmarshaledRow["_measurement"].(string)
			temp.RequestID = unmarshaledRow["request_id"].(string)
			temp.Timestamp = rows.Record().Time()
			lastTableIndex = table
		}
	}
	// append the last temp, if the output length is more than zero
	if len(output) > 0 || temp.RequestID != "" {
		output = append(output, temp)
	}

	return output, nil
}

func unmarshalInfluxRow(row string) (map[string]interface{}, error) {
	// because csv.NewReader() accepts io.Reader, we'll create one from strings pkg
	input := strings.NewReader(row)
	reader := csv.NewReader(input)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	records, err := reader.Read()
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("reading row value to csv: %v", err)
	}

	// find records length
	// because it's a jagged array, we'll do a nested one
	var recordsLength = len(records)

	output := make(map[string]interface{}, recordsLength)
	for _, rec := range records {
		kv := strings.Split(rec, ":")
		output[kv[0]] = kv[1]
	}

	return output, nil
}
