package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// FunFact is the handler for generating fun fact about the user
// after they had done their coding test.
func (d *Dependency) FunFact(w http.ResponseWriter, r *http.Request) {
	var member Member

	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate for empty memberID first
	if member.ID == "" {
		http.Error(w, "member_id is empty", http.StatusBadRequest)
		return
	}

	// Read about buffered channel vs non-buffered channels
	wpm := make(chan int8, 1)
	deletionRate := make(chan float64, 1)
	attempt := make(chan int8, 1)

	// Run all the calculate function concurently
	go d.CalculateWordsPerMinute(r.Context(), member.ID, wpm)
	go d.CalculateDeletionRate(r.Context(), member.ID, deletionRate)
	go d.CalculateSubmissionAttempts(r.Context(), member.ID, attempt)

	var result = struct {
		Wpm          int8    `json:"wpm"`
		DeletionRate float64 `json:"deletion_rate"`
		Attempt      int8    `json:"attempt"`
	}{
		<-wpm,
		<-deletionRate,
		<-attempt,
	}

	res, err := json.Marshal(result)
	if err != nil {
		// handle the error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	headers := w.Header()
	headers.Set("content-type", "application/json")
	w.Write(res)
}

func (d *Dependency) CalculateWordsPerMinute(ctx context.Context, memberID string, result chan int8) {
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// TODO:  ini buat ngambil nganu, jangan lupa result
	result, err := queryAPI.Query(context.TODO(), `
		from(bucket: "spectator")
			|> range(start: -1d)
			|> filter(fn: (r) => r["_measurement"] == "coding_event")
			|> filter(fn: (r) => r["_event"] == "keystroke")
			|> filter(fn: (r) => r["_actor"] == "`+memberID+`")
			|> aggregateWindow(
					every: 1m,
					fn: (tables=<-, column) => tables |> count()
		  )
			|> yield(name: "count")
	`)
	if err != nil {
		// FIXME: seharusnya jangan panic
		panic(err)
	}

	var wpmTot, keyTot = 0, 0
	for result.Next() {
		keytotal := result.Record().Value().(int64)
		wpmTot += keytotal / 5
		keyTot += 1
	}

	// Cara calculate WPM:
	// SELECT semua KeystrokeEvent, group by TIME, each TIME itu 1 menit
	// for every 1 minute, hitung total keystroke event itu,
	// terus dibagi dengan 5
	//
	// Itu baru dapet WPM per 1 menit itu.
	// Nanti mungkin bisa di store data nya jadi slice (per 1 menit,
	// ngga perlu specify menit keberapanya, karena slice pasti urut)
	// terus return ke channel hasil average dari semua menit yang ada

	// Return the result here
	result <- int8(wpmTot / keyTot)
}

func (d *Dependency) CalculateSubmissionAttempts(ctx context.Context, memberID string, result chan int8) {
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// number of question submission attempts
	// TODO:  ini buat ngambil nganu, jangan lupa result
	// SELECT COUNT(_time) FROM spectator WHERE _type = "coding_attempted"
	result, err := queryAPI.Query(ctx, `from(bucket: "spectator")
		|> range(start: -1d)
		|> filter(fn: (r) => r["_measurement"] == "coding_event")
		|> filter(fn: (r) => r["_event"] == "coding_attempted")
		|> filter(fn: (r) => r["_actor"] == "`+memberID+`")
		|> group(columns: "question_id")
		|> group()
		|> count()
	`)
	if err != nil {
		// FIXME: seharusnya jangan panic
		panic(err)
	}

	// terus langsung return hasilnya
	// tapi bisa juga di group per question, jadi
	// misalnya untuk question #1, dia ada 5 attempt, question #2 ada 10 attempt
	// and so on so forth.

	// Return the result here
	result <- 42
}

func (d *Dependency) CalculateDeletionRate(ctx context.Context, memberID string, result chan float64) {
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// TODO:  ini buat ngambil nganu, jangan lupa result
	result, err := queryAPI.Query(context.TODO(), `from(bucket: "spectator")
  |> range(start: -1d)
  |> filter(fn: (r) => r["_measurement"] == "coding_event")
  |> filter(fn: (r) => r["_event"] == "keystroke")
  |> filter(fn: (r) => r["_actor"] == "`+memberID+`")
  |> filter(fn: (r) =>(r["_value"] == "backspace" or r["_value"] == "delete"))
  |> count()
  |> yield(name: "count")`)

	if err != nil {
		panic(err)
	}

	result.Next()
	delTot := result.Record().Value().(int64)

	result, err = queryAPI.Query(context.TODO(), `from(bucket: "spectator")
  |> range(start: -1d)
  |> filter(fn: (r) => r["_measurement"] == "coding_event")
  |> filter(fn: (r) => r["_event"] == "keystroke")
  |> filter(fn: (r) => r["_actor"] == "`+memberID+`")
	|> count()
  |> yield(name: "count")`)

	if err != nil {
		panic(err)
	}

	result.Next()
	tot := result.Record().Value().(int64)

	result <- (float64(delTot) / float64(tot))

	// SELECT semua KeystrokeEvent WHERE value = delete OR value = backspace
	// terus jumlahin
	// dah gitu doang.

	// Return the result here
}
