package main

import (
	"context"
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
	defer func(){
		r := recover()
		if r != nil {
			log.Printf("panic: %v", r.(error))
		}
	}()
	writeAPI := d.DB.WriteAPI(d.Org, "log")
	point := influxdb2.NewPoint(
		p.Data.Level,
		map[string]string{
			"request_id": p.Data.RequestID,
			"application": p.Data.Application,
			"environment": p.Data.Environment,
		},
		map[string]interface{}{
			"language": p.Data.Language,
			"message": p.Data.Message,
			"body": p.Data.Body,
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
	return str.String()
}

func (d *Dependency) fetchLog(ctx context.Context, query queries) ([]Data, error) {
	queryAPI := d.DB.QueryAPI(d.Org)
	rows, err := queryAPI.Query(
		ctx,
		`from(bucket: "log")
		|> range(start: 0)
		|> sort(columns: ["_time"])
		|> group(columns: ["_time"])
		|> filter(fn: (r) => r["_measurement"] == "error")`,
	)
	if err != nil {
		return []Data{}, err
	}

	defer rows.Close()

	var output []Data
	for rows.Next() {
		var temp Data
		temp.RequestID = rows.Record().ValueByKey("request_id").(string)
		temp.Application = rows.Record().ValueByKey("application").(string)
		temp.Timestamp = rows.Record().Time()
		temp.Level = rows.Record().Measurement()
		log.Printf("current rows: %s", rows.Record().String())
		log.Printf("row field: %s", rows.Record().Field())
		output = append(output, temp)
	}

	return output, nil
}