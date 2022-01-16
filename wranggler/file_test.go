package main_test

import (
	"context"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	worker "worker"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
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

	for i := 0; i < 50; i++ {

		p := influxdb2.NewPoint(
			"personal_info",
			map[string]string{
				"session_id": id.String(),
			},
			map[string]interface{}{
				"student_number":      "",
				"hours_of_practice":   rand.Int31n(666),
				"years_of_experience": rand.Int31n(5),
				"familiar_languages":  "",
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		)

		writeSessionAPI.WritePoint(ctx, p)

		p = influxdb2.NewPoint(
			"sam_test",
			map[string]string{
				"session_id": id.String(),
			},
			map[string]interface{}{
				"aroused_level": rand.Int31n(3),
				"pleased_level": rand.Int31n(3),
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		)

		writeSessionAPI.WritePoint(ctx, p)

		// code_event_keystroke
		p = influxdb2.NewPoint(
			"coding_event_keystroke",
			map[string]string{
				"session_id": id.String(),
			},
			map[string]interface{}{
				"key_char": "a",
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		)

		writeInputAPI.WritePoint(ctx, p)

		// code_event_mouseclick
		p = influxdb2.NewPoint(
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
		)

		writeInputAPI.WritePoint(ctx, p)

		// code_event_mouseclick
		p = influxdb2.NewPoint(
			"coding_event_mousemove",
			map[string]string{
				"session_id":      id.String(),
				"question_number": "1",
			},
			map[string]interface{}{
				"direction":     "right",
				"x_position":    rand.Int31n(1337),
				"y_position":    rand.Int31n(768),
				"window_width":  rand.Int31n(1337),
				"window_height": rand.Int31n(768),
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		)

		writeInputAPI.WritePoint(ctx, p)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func(w *sync.WaitGroup) {
		deps.CreateFile(w, id)
	}(&wg)

	wg.Wait()

	filesJson, err := filepath.Glob("./*.json")
	if err != nil {
		t.Fatal(err)
	}
	filesCSV, err := filepath.Glob("./*.csv")
	if err != nil {
		t.Fatal(err)
	}

	result := append(filesJson, filesCSV...)

	if len(result) == 0 {
		t.Fail()
	}

	for _, f := range result {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}

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
