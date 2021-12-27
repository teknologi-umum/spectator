package main_test

import (
	"context"
	"fmt"
	//"log"
	//"os"
	"testing"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	// "github.com/minio/minio-go/v7"
	// "github.com/minio/minio-go/v7/pkg/credentials"
	"encoding/json"
	//rori "rori"
	"time"
)

type Point struct {
	Type  string `json:"t"`
	Event string `json:"e"`
	Actor string `json:"a"`
	Value string `json:"v"`
}

type Submission struct {
	Type           string `json:"t"`
	Event          string `json:"e"`
	Actor          string `json:"a"`
	QuestionNumber string `json:"q"`
	Value          string `json:"v"`
}

func TestShit(t *testing.T) {
	x := `
[
	{"t": "coding_event","e": "keystroke","a":"3","v":"e"},
	{"t": "coding_event","e": "keystroke","a":"2","v":"a"},
	{"t": "coding_event","e": "keystroke","a":"2","v":"d"},
	{"t": "coding_event","e": "keystroke","a":"5","v":";"},
	{"t": "coding_event","e": "keystroke","a":"3","v":"-"},
	{"t": "coding_event","e": "keystroke","a":"3","v":"delete"},
	{"t": "coding_event","e": "keystroke","a":"2","v":"backspace"},
	{"t": "coding_event","e": "keystroke","a":"2","v":"x"},
	{"t": "coding_event","e": "keystroke","a":"2","v":"a"},
	{"t": "coding_event","e": "keystroke","a":"2","v":"c"},
	{"t": "coding_event","e": "keystroke","a":"4","v":"delete"},
	{"t": "coding_event","e": "keystroke","a":"3","v":"backspace"},
	{"t": "coding_event","e": "mouse_movement","a":"2","v":"u"},
	{"t": "coding_event","e": "mouse_movement","a":"3","v":"4"},
	{"t": "coding_event","e": "mouse_movement","a":"4","v":"2"}
]`

	y := `[ 
	{"t": "test_event","e": "submission_attempt","a":"2","q":"1","v": "false"}, 
	{"t": "test_event","e": "submission_attempt","a":"3","q":"1","v": "false"}, 
	{"t": "test_event","e": "submission_attempt","a":"4","q":"1","v": "false"}, 
	{"t": "test_event","e": "submission_attempt","a":"1","q":"1","v": "false"}, 
	{"t": "test_event","e": "submission_attempt","a":"1","q":"1","v": "false"}, 
	{"t": "test_event","e": "submission_attempt","a":"1","q":"2","v": "false"}, 
	{"t": "test_event","e": "submission_attempt","a":"1","q":"3","v": "false"}, 
	{"t": "test_event","e": "submission_attempt","a":"2","q":"3","v": "false"}, 
	{"t": "test_event","e": "submission_attempt","a":"2","q":"1","v": "false"}, 
	{"t": "test_event","e": "submission_attempt","a":"2","q":"1","v": "false"}, 
	{"t": "test_event","e": "submission_attempt","a":"2","q":"1","v": "false"}, 
	{"t": "test_event","e": "submission_attempt","a":"2","q":"1","v": "false"}, 
	{"t": "test_event","e": "submission_attempt","a":"2","q":"1","v": "false"}
]
`

	batch := []Point{}
	err := json.Unmarshal([]byte(x), &batch)
	if err != nil {
		t.Error(err)
	}

	if len(batch) == 0 {
		t.Errorf("batch empty")
	}

	batchY := []Submission{}
	err = json.Unmarshal([]byte(y), &batchY)
	if err != nil {
		t.Error(err)
	}

	if len(batchY) == 0 {
		t.Errorf("BatchY Empty")
	}

	const token = "l5fi5YuUJlBjq0EeLYLYqf-7jz7tGCq-QNGWGbWA8sME5pEFtCW1RUHk4oVChYcGNA9g92BQVLeL6pCk-NAmkA=="
	const bucket = "spectator"
	const org = "teknum"

	//	minioHost, ok := os.LookupEnv("MINIO_HOST")
	//	if !ok {
	//		log.Fatalln("MINIO_HOST envar missing")
	//	}
	//
	//	minioID, ok := os.LookupEnv("MINIO_ACCESS_ID")
	//	if !ok {
	//		log.Fatalln("MINIO_ACCESS_ID envar missing")
	//	}
	//
	//	minioSecret, ok := os.LookupEnv("MINIO_SECRET_KEY")
	//	if !ok {
	//		log.Fatalln("MINIO_SECRET_KEY envar missing")
	//	}
	//
	// Create InfluxDB instance
	influxConn := influxdb2.NewClient("http://localhost:8086", token)
	defer influxConn.Close()

	// Create Minio instance
	//	minioConn, err := minio.New(minioHost, &minio.Options{
	//		Creds:  credentials.NewStaticV4(minioID, minioSecret, ""),
	//		Secure: true,
	//	})
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	// Initialize dependency injection str
	writeAPI := influxConn.WriteAPI(org, bucket)

	queryAPI := influxConn.QueryAPI(org)

	/*
	   _time: Date,
	   _type: "coding_event",
	   _question_number: string,
	   _event: "keystroke",
	   key_code: string,
	   unrelated_key: boolean,
	   modifier: string,
	   _actor: string
	*/
	// write line protocol

	for _, item := range batch {

		writeAPI.WritePoint(influxdb2.NewPointWithMeasurement(item.Type).AddTag("_event", item.Event).AddTag("_actor", item.Actor).AddField("value", item.Value).SetTime(time.Now()))
	}
	for _, item := range batchY {
		writeAPI.WritePoint(influxdb2.NewPointWithMeasurement(item.Type).AddTag("_event", item.Event).AddTag("_actor", item.Actor).AddTag("question_number", item.QuestionNumber).AddField("value", item.Value).SetTime(time.Now()))
	}

	// Flush writes
	writeAPI.Flush()

	result, err := queryAPI.Query(context.TODO(), `from(bucket: "spectator")
  |> range(start: -1d)
  |> filter(fn: (r) => r["_measurement"] == "coding_event")
  |> filter(fn: (r) => r["_event"] == "keystroke")
  |> filter(fn: (r) => r["_actor"] == "2")
  |> filter(fn: (r) =>(r["_value"] == "backspace" or r["_value"] == "delete"))
  |> count()
  |> yield(name: "count")`)

	if err != nil || result == nil {
		t.Error(err)
	}

	result.Next()
	delTot := result.Record().Value().(int64)

	result, err = queryAPI.Query(context.TODO(), `from(bucket: "spectator")
  |> range(start: -1d)
  |> filter(fn: (r) => r["_measurement"] == "coding_event")
  |> filter(fn: (r) => r["_event"] == "keystroke")
  |> filter(fn: (r) => r["_actor"] == "2")
	|> count()
  |> yield(name: "count")`)

	if err != nil {
		t.Error(err)
	}

	result.Next()
	tot := result.Record().Value().(int64)

	fmt.Println((float64(delTot) / float64(tot)))

	result, err = queryAPI.Query(context.TODO(), `from(bucket: "spectator")
		|> range(start: -1d)
		|> filter(fn: (r) => r["_measurement"] == "test_event")
		|> filter(fn: (r) => r["_event"] == "submission_attempt")
		|> filter(fn: (r) => r["_actor"] == "2")
		|> group(columns: ["question_id"])
		|> count()
	`)
	if err != nil {
		// FIXME: seharusnya jangan panic
		panic(err)
	}

	// why the result isn't array
	for result.Next() {
		fmt.Println(result.Record())
	}

}
