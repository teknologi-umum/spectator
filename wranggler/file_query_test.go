package main_test

import (
	"context"
	"math/rand"
	"testing"
	"time"
	worker "worker"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestQueryKeystrokes(t *testing.T) {
	// TODO:
	// 1. insert some fake data into the influx db
	// 2. query the data with the function
	// 3. compare the length of the result and the length of fake data
	// add another test (maybe a subtest, or another test function)
	// that checks if there is no data to be queried.
	// we must check if that (rare and edgy) event happen,
	// so what would the software react?

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

	min := time.Date(2019, 5, 2, 1, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 5, 2, 1, 4, 0, 0, time.UTC).Unix()
	delta := max - min

	for i := 0; i < 50; i++ {
		p := influxdb2.NewPoint(
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
	}

	readInputAPI := db.QueryAPI(deps.DBOrganization)
	result, err := deps.QueryKeystrokes(ctx, readInputAPI, id)
	if err != nil {
		t.Fatal("Test Query Keystroke", err)
		return
	}

	if len(result) == 50 {
		t.Log("Test Query Keystorke done")
		return
	} else {
		t.Fatal("Data not 50")
	}
}

func TestQueryMouseClick(t *testing.T) {

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

	min := time.Date(2019, 5, 2, 1, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 5, 2, 1, 4, 0, 0, time.UTC).Unix()
	delta := max - min

	for i := 0; i < 50; i++ {
		p := influxdb2.NewPoint(
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
	}

	readInputAPI := db.QueryAPI(deps.DBOrganization)
	result, err := deps.QueryMouseClick(ctx, readInputAPI, id)
	if err != nil {
		t.Fatal("Test Query Mouse Click", err)
		return
	}

	if len(result) == 50 {
		t.Log("Test Query Mouse Click done")
	} else {
		t.Fatal("Data not 50")
	}
	return
}

func TestQueryMouseMove(t *testing.T) {

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

	min := time.Date(2019, 5, 2, 1, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 5, 2, 1, 4, 0, 0, time.UTC).Unix()
	delta := max - min

	for i := 0; i < 50; i++ {
		p := influxdb2.NewPoint(
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

	readInputAPI := db.QueryAPI(deps.DBOrganization)
	result, err := deps.QueryMouseMove(ctx, readInputAPI, id)
	if err != nil {
		t.Fatal("Test Query Mouse Move", err)
		return
	}

	if len(result) == 50 {
		t.Log("Test Query Mouse Move done")
	} else {
		t.Fatal("Data not 50")
	}
	return
}

func TestQueryPersonalInfo(t *testing.T) {

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
	}

	readInputAPI := db.QueryAPI(deps.DBOrganization)
	result, err := deps.QueryPersonalInfo(ctx, readInputAPI, id)
	if err != nil {
		t.Fatal("Test Query Personal Info", err)
		return
	}

	if len(result) == 50 {
		t.Log("Test Query Personal Info")
	} else {
		t.Fatal("Data not 50")
	}
	return
}

func TestQuerySamTest(t *testing.T) {

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
	writeSessionAPI := db.WriteAPIBlocking(deps.DBOrganization, worker.BucketSessionEvents)

	min := time.Date(2019, 5, 2, 1, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 5, 2, 1, 4, 0, 0, time.UTC).Unix()
	delta := max - min

	for i := 0; i < 50; i++ {
		p := influxdb2.NewPoint(
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
	}

	readInputAPI := db.QueryAPI(deps.DBOrganization)
	result, err := deps.QuerySAMTest(ctx, readInputAPI, id)
	if err != nil {
		t.Fatal("Test Query Sam Test", err)
		return
	}

	if len(result) == 50 {
		t.Log("Test Query Sam Test")
	} else {
		t.Fatal("Data not 50")
	}
	return
}
