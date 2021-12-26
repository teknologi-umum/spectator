package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func (d *Dependency) writeIntoLog(ctx context.Context, p Payload) error {
	// write defaults first
	if p.Data.Environment == "" {
		p.Data.Environment = "unset"
	}

	if p.Data.Level == "" {
		p.Data.Level = "debug"
	}

	if p.Data.Timestamp.IsZero() {
		p.Data.Timestamp = time.Now()
	}

	// recovering here in case of any error
	defer func() {
		r := recover()
		if r != nil {
			log.Printf("panic: %v", r.(error))
		}
	}()
	writeAPI := d.DB.WriteAPI(d.Org, "log")
	point := influxdb2.NewPoint(
		p.Data.Level,
		map[string]string{
			"request_id":  p.Data.RequestID,
			"application": p.Data.Application,
			"environment": p.Data.Environment,
		},
		map[string]interface{}{
			"language": p.Data.Language,
			"message":  p.Data.Message,
			"body":     p.Data.Body,
		},
		p.Data.Timestamp,
	)

	writeAPI.WritePoint(point)
	writeAPI.Flush()
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

func (d *Dependency) fetchLog(ctx context.Context, query queries) ([]Data, error) {
	queryAPI := d.DB.QueryAPI(d.Org)
	// build query for influx
	queryStr := buildQuery(query)

	rows, err := queryAPI.Query(ctx, queryStr)
	if err != nil {
		return []Data{}, err
	}

	defer rows.Close()

	var output []Data
	var temp Data
	var lastTableIndex int = -1
	for rows.Next() {
		unmarshaledRow, err := unmarshalInfluxRow(rows.Record().String())
		if err != nil {
			return []Data{}, err
		}

		table, err := strconv.Atoi(unmarshaledRow["table"].(string))
		if err != nil {
			return []Data{}, err
		}
		if table == lastTableIndex {
			switch unmarshaledRow["_field"].(string) {
			case "body":
				temp.Body = unmarshaledRow["_value"].(map[string]interface{})
			case "language":
				temp.Language = unmarshaledRow["_value"].(string)
			case "message":
				temp.Message = unmarshaledRow["_value"].(string)
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
