package main_test

import (
	"context"
	"math/rand"
	"testing"
	"time"
	"worker/proto"

	worker "worker"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func TestConvertDataToJSON(t *testing.T) {
	t.Cleanup(cleanup)
	rand.Seed(time.Now().Unix())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	deps := worker.Dependency{
		DB:             db,
		DBOrganization: dbOrganization,
		Bucket:         bucket,
	}

	id, err := uuid.NewUUID()
	if err != nil {
		t.Error(err)
	}

	writeInputAPI := db.WriteAPIBlocking(deps.DBOrganization, worker.BucketInputEvents)
	writeSessionAPI := db.WriteAPIBlocking(deps.DBOrganization, worker.BucketSessionEvents)
	min := time.Date(2019, 5, 2, 1, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 5, 2, 1, 4, 0, 0, time.UTC).Unix()
	delta := max - min

	var sessionPoints = []*write.Point{
		influxdb2.NewPoint(
			"personal_info",
			map[string]string{
				"session_id": id.String(),
			},
			map[string]interface{}{
				"student_number":      "",
				"hours_of_practice":   0,
				"years_of_experience": 0,
				"familiar_languages":  "",
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		),
		influxdb2.NewPoint(
			"sam_test",
			map[string]string{
				"session_id": id.String(),
			},
			map[string]interface{}{
				"aroused_level": 0,
				"pleased_level": 0,
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		),
	}
	
	for _, p := range sessionPoints {
		writeSessionAPI.WritePoint(ctx, p)
	}
	var inputPoints = []*write.Point{
		influxdb2.NewPoint(
			"code_submission",
			map[string]string{
				"session_id":      id.String(),
				"question_number": "1",
			},
			map[string]interface{}{
				"code":     "let () = print_endline \"UwU\"",
				"language": "OCaml",
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		),
		influxdb2.NewPoint(
			"coding_event_keystroke",
			map[string]string{
				"session_id": id.String(),
			},
			map[string]interface{}{
				"key_char": "a",
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		),
		influxdb2.NewPoint(
			"coding_event_mouseclick",
			map[string]string{
				"session_id":      id.String(),
				"question_number": "1",
			},
			map[string]interface{}{
				"key_char":     "a",
				"right_click":  false,
				"left_click":   false,
				"middle_click": false,
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		),
		influxdb2.NewPoint(
			"coding_event_mousemove",
			map[string]string{
				"session_id":      id.String(),
				"question_number": "1",
			},
			map[string]interface{}{
				"direction":     "right",
				"x_position":    0,
				"y_position":    0,
				"window_width":  0,
				"window_height": 0,
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		),
	}

	for _, p := range inputPoints {
		writeInputAPI.WritePoint(ctx, p)
	}

	deps.GenerateFiles(ctx, &proto.Member{SessionId: id.String()})

	// TODO: DELETE THE FILE HASBEN GENERATED, BUT NOT TODAY MAY HAD AGAINST ME WILL NOW
}

func TestConvertDataToCSV(t *testing.T) {
	// data := []main.SampleInput{
	// 	{
	// 		Time:  time.Now().Add(time.Second * 1),
	// 		Actor: "James",
	// 		X:     20,
	// 		Y:     13,
	// 	},
	// 	{
	// 		Time:  time.Now().Add(time.Second * 2),
	// 		Actor: "James, Riyadi",
	// 		X:     21,
	// 		Y:     13,
	// 	},
	// 	{
	// 		Time:  time.Now().Add(time.Second * 3),
	// 		Actor: "James Riyadi",
	// 		X:     22,
	// 		Y:     14,
	// 	},
	// }

	// res, err := main.ConvertDataToCSV(data)
	// if err != nil {
	// 	t.Errorf("an error was thrown: %v", err)
	// }

	// t.Log(string(res))
}
