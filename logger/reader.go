package main

import (
	"context"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	pb "logger/proto"
	"strconv"
	"strings"

	"time"
)

type queries struct {
	Level       string
	RequestID   string
	Application string
	TimeFrom    time.Time
	TimeTo      time.Time
}

func (d *Dependency) ReadLog(ctx context.Context, r *pb.ReadLogRequest) (*pb.ReadLogResponse, error) {
	var timestampFrom time.Time
	var timestampTo time.Time
	if r.GetTimestampFrom() == 0 {
		timestampFrom = time.Time{}
	} else {
		timestampFrom = time.UnixMilli(r.GetTimestampFrom())
	}

	if r.GetTimestampTo() == 0 {
		timestampTo = time.Time{}
	} else {
		timestampTo = time.UnixMilli(r.GetTimestampTo())
	}

	var query = queries{
		Level:       r.GetLevel().String(),
		RequestID:   r.GetRequestId(),
		Application: r.GetApplication(),
		TimeFrom:    timestampFrom,
		TimeTo:      timestampTo,
	}

	logs, err := d.fetchLog(ctx, query)
	if err != nil {
		return &pb.ReadLogResponse{}, fmt.Errorf("reading log: %v", err)
	}

	return &pb.ReadLogResponse{
		Data: d.convertIntoProtoData(logs),
	}, nil
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
		return []LogData{}, fmt.Errorf("querying data: %v", err)
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
					return []LogData{}, fmt.Errorf("decoding string: %v", err)
				}
				body := make(map[string]string, 100)
				err = json.Unmarshal(bodyBytes, &body)
				if err != nil {
					return []LogData{}, fmt.Errorf("unmarshaling json: %v", err)
				}
				temp.Body = body
			case "language":
				temp.Language = unmarshaledRow["_value"].(string)
			case "message":
				messageBytes, err := hex.DecodeString(unmarshaledRow["_value"].(string))
				if err != nil {
					return []LogData{}, fmt.Errorf("decoding string: %v", err)
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
