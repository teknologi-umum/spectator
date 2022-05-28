package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
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

	if q.RequestID == "" {
		str.WriteString("|> group(columns: [\"request_id\", \"_time\"])\n")
	} else {
		str.WriteString("|> group(columns: [\"_time\"])\n")
	}

	if q.Level != "" {
		str.WriteString(`|> filter(fn: (r) => r["_measurement"] == "` + q.Level + `")` + "\n")
	}

	if q.Application != "" {
		str.WriteString(`|> filter(fn: (r) => r["application"] == "` + q.Application + `")` + "\n")
	}

	if q.RequestID != "" {
		str.WriteString(`|> filter(fn: (r) => r["request_id"] == "` + q.RequestID + `")` + "\n")
	}

	str.WriteString("|> yield()\n")

	return str.String()
}

func (d *Dependency) fetchLog(ctx context.Context, query queries) ([]LogData, error) {
	queryAPI := d.DB.QueryAPI(d.Org)
	// build query for influx
	queryStr := buildQuery(query)
	if d.Debug {
		log.Println(queryStr)
	}

	rows, err := queryAPI.Query(ctx, queryStr)
	if err != nil {
		return []LogData{}, fmt.Errorf("querying data: %v", err)
	}
	defer rows.Close()

	var output []LogData
	var temp LogData
	var tablePosition int64
	for rows.Next() {
		record := rows.Record()
		table, ok := rows.Record().ValueByKey("table").(int64)
		if !ok {
			table = 0
		}
		switch record.Field() {
		case "body":
			bodyJSON, ok := record.Value().(string)
			if !ok {
				bodyJSON = ""
			}
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
			temp.Language, ok = record.Value().(string)
			if !ok {
				temp.Language = ""
			}
		case "message":
			value, ok := record.Value().(string)
			if !ok {
				value = ""
			}
			messageBytes, err := hex.DecodeString(value)
			if err != nil {
				return []LogData{}, fmt.Errorf("decoding string: %v", err)
			}
			temp.Message = string(messageBytes)
		}

		if d.Debug {
			log.Println(rows.Record().String())
			log.Printf("table %d\n", rows.Record().Table())
		}

		if table != 0 && table > tablePosition {
			output = append(output, temp)
			tablePosition = table
		} else {
			var ok bool

			temp.Application, ok = record.ValueByKey("application").(string)
			if !ok {
				temp.Application = ""
			}

			temp.Environment, ok = record.ValueByKey("environment").(string)
			if !ok {
				temp.Environment = ""
			}

			temp.Level = record.Measurement()

			temp.RequestID, ok = record.ValueByKey("request_id").(string)
			if !ok {
				temp.RequestID = ""
			}

			temp.Timestamp = record.Time()
		}
	}
	// append the last temp, if the output length is more than zero
	if len(output) > 0 || temp.RequestID != "" {
		output = append(output, temp)
	}

	return output, nil
}
