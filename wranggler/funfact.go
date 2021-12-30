package main

import (
	"context"
	"encoding/json"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strconv"
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

	if _, err := strconv.ParseInt(member.ID, 10, 64); err != nil {
		http.Error(w, "member_id is empty", http.StatusBadRequest)
		return
	}

	// Read about buffered channel vs non-buffered channels
	wpm := make(chan int8, 1)
	deletionRate := make(chan float64, 1)
	attempt := make(chan []int8, 1)

	// Run all the calculate function concurently
	errs, ctx := errgroup.WithContext(r.Context())
	errs.Go(func() error {
		return d.CalculateWordsPerMinute(ctx, member.ID, wpm)
	})
	errs.Go(func() error {
		return d.CalculateDeletionRate(ctx, member.ID, deletionRate)
	})
	errs.Go(func() error {
		return d.CalculateSubmissionAttempts(ctx, member.ID, attempt)
	})

	err = errs.Wait()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result = struct {
		Wpm          int8    `json:"wpm"`
		DeletionRate float64 `json:"deletion_rate"`
		Attempt      []int8  `json:"attempt"`
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

func (d *Dependency) CalculateWordsPerMinute(ctx context.Context, memberID string, result chan int8) error {
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	res, err := queryAPI.Query(ctx, `
		from(bucket: "spectator")
			|> range(start: -1d)
			|> filter(fn: (r) => r["event"] == "coding_event_keystroke")
			|> filter(fn: (r) => r["session_id"] == "`+memberID+`")
			|> aggregateWindow(
					every: 1m,
					fn: (tables=<-, column) => tables |> count()
		  )
			|> yield(name: "count")
	`)
	if err != nil {
		return err
	}

	var wpmTotal, keyTotal = 0, 0
	for res.Next() {
		keytotal := res.Record().Value().(int64)
		wpmTotal += int(keytotal) / 5
		keyTotal += 1
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
	result <- int8(wpmTotal / keyTotal)
	return nil
}

func (d *Dependency) CalculateSubmissionAttempts(ctx context.Context, memberID string, result chan []int8) error {
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// number of question submission attempts
	// TODO:  ini buat ngambil nganu, jangan lupa result
	// SELECT COUNT(_time) FROM spectator WHERE _type = "coding_attempted"
	_, err := queryAPI.Query(ctx, `from(bucket: "spectator")
		|> range(start: -1d)
		|> filter(fn: (r) => r["type"] == "code_test_attempt")
		|> filter(fn: (r) => r["session_id"] == "`+memberID+`")
		|> group(columns: ["question_id"])
		|> count()
	`)
	if err != nil {
		return err
	}

	// FIXME: the result not array , the reasou UNKNOW

	// terus langsung return hasilnya
	// tapi bisa juga di group per question, jadi
	// misalnya untuk question #1, dia ada 5 attempt, question #2 ada 10 attempt
	// and so on so forth.

	// Return the result here
	result <- []int8{}
	return nil
}

func (d *Dependency) CalculateDeletionRate(ctx context.Context, memberID string, result chan float64) error {
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// TODO:  ini buat ngambil nganu, jangan lupa result
	res, err := queryAPI.Query(context.TODO(), `from(bucket: "spectator")
  |> range(start: -1d)
  |> filter(fn: (r) => r["type"] == "coding_event_keystroke")
  |> filter(fn: (r) => r["session_id"] == "`+memberID+`")
  |> filter(fn: (r) => (r["key_char"] == "backspace" or r["key_char"] == "delete"))
	|> filter(fn: (r) => not(
		(r["_field"] == "shift" and r["_value"] == true) or 
		(r["_field"] == "alt" and r["_value"] == true) or
		(r["_field"] == "control" and r["_value"] == true) or
		(r["_field"] == "meta" and r["_value"] == true) or 
		(r["_field"] == "unrelatedkey" and r["_value"] == true)
	)
  |> count()
  |> yield(name: "count")`)

	if err != nil {
		return (err)
	}

	res.Next()
	delTot := res.Record().Value().(int64)

	res, err = queryAPI.Query(context.TODO(), `from(bucket: "spectator")
  |> range(start: -1d)
  |> filter(fn: (r) => r["type"] == "coding_event_keystroke")
  |> filter(fn: (r) => r["session_id"] == "`+memberID+`")
	|> filter(fn: (r) => not(
		(r["_field"] == "shift" and r["_value"] == true) or 
		(r["_field"] == "alt" and r["_value"] == true) or
		(r["_field"] == "control" and r["_value"] == true) or
		(r["_field"] == "meta" and r["_value"] == true) or 
		(r["_field"] == "unrelatedkey" and r["_value"] == true)
	)
	|> count()
  |> yield(name: "count")`)

	if err != nil {
		return (err)
	}

	res.Next()
	tot := res.Record().Value().(int64)

	result <- (float64(delTot) / float64(tot))

	// SELECT semua KeystrokeEvent WHERE value = delete OR value = backspace
	// terus jumlahin
	// dah gitu doang.

	// Return the result here
	return nil
}
