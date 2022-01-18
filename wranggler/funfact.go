package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"worker/logger"
	pb "worker/proto"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// FunFact is the handler for generating fun fact about the user
// after they had done their coding test.
func (d *Dependency) FunFact(ctx context.Context, in *pb.Member) (*pb.FunFactResponse, error) {
	// Parse UUID
	sessionID, err := uuid.Parse(in.GetSessionId())
	if err != nil {
		defer d.Log(
			err.Error(),
			logger.Level_ERROR.Enum(),
			in.RequestId,
			map[string]string{
				"session_id": in.GetSessionId(),
				"function":   "funfact",
				"info":       "parsing uuid",
			},
		)
		return &pb.FunFactResponse{}, fmt.Errorf("parsing uuid: %v", err)
	}

	// Read about buffered channel vs non-buffered channels
	wpm := make(chan uint32, 1)
	deletionRate := make(chan float32, 1)
	attempt := make(chan uint32, 1)

	// Run all the calculate function concurently
	errs, gctx := errgroup.WithContext(ctx)
	errs.Go(func() error {
		return d.CalculateWordsPerMinute(gctx, sessionID, wpm)
	})
	errs.Go(func() error {
		return d.CalculateDeletionRate(gctx, sessionID, deletionRate)
	})
	errs.Go(func() error {
		return d.CalculateSubmissionAttempts(gctx, sessionID, attempt)
	})

	err = errs.Wait()
	if err != nil {
		defer d.Log(
			err.Error(),
			logger.Level_ERROR.Enum(),
			in.RequestId,
			map[string]string{
				"session_id": in.GetSessionId(),
				"function":   "funfact",
				"info":       "calculating fun fact",
			},
		)
		return &pb.FunFactResponse{}, fmt.Errorf("calculating fun fact: %v", err)
	}

	var result = struct {
		Wpm          uint32  `json:"wpm"`
		DeletionRate float32 `json:"deletion_rate"`
		Attempt      uint32  `json:"attempt"`
	}{
		<-wpm,
		<-deletionRate,
		<-attempt,
	}

	return &pb.FunFactResponse{
		WordsPerMinute:     result.Wpm,
		DeletionRate:       result.DeletionRate,
		SubmissionAttempts: result.Attempt,
	}, nil
}

func (d *Dependency) CalculateWordsPerMinute(ctx context.Context, sessionID uuid.UUID, result chan uint32) error {
	// Cara calculate WPM:
	// SELECT semua KeystrokeEvent, group by TIME, each TIME itu 1 menit
	// for every 1 minute, hitung total keystroke event itu,
	// terus dibagi dengan 5
	//
	// Itu baru dapet WPM per 1 menit itu.
	// Nanti mungkin bisa di store data nya jadi slice (per 1 menit,
	// ngga perlu specify menit keberapanya, karena slice pasti urut)
	// terus return ke channel hasil average dari semua menit yang ada

	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	rows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "coding_event_keystroke")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> window(every: 1m)
		|> filter(fn: (r) => r["_field"] == "key_char")
		|> sort(columns: ["_time"])`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	var wordsPerMinute []uint32
	var tablePosition int64
	var temporaryWords uint32
	for rows.Next() {
		record := rows.Record()
		table, ok := record.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		temporaryWords += 1

		if table != 0 && table > tablePosition {
			wordsPerMinute = append(wordsPerMinute, temporaryWords)
			temporaryWords = 0
		}
	}

	// Append the last value
	if len(wordsPerMinute) > 0 && temporaryWords != 0 {
		wordsPerMinute = append(wordsPerMinute, temporaryWords)
	}

	var averageWpm uint32
	var wordsSum uint32
	for _, wpm := range wordsPerMinute {
		wordsSum += wpm / 5
	}

	averageWpm = wordsSum / uint32(len(wordsPerMinute))

	// Return the result here
	result <- averageWpm
	return nil
}

func (d *Dependency) CalculateSubmissionAttempts(ctx context.Context, sessionID uuid.UUID, result chan uint32) error {
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// number of question submission attempts
	// TODO:  ini buat ngambil nganu, jangan lupa result
	// SELECT COUNT(_time) FROM spectator WHERE _type = "coding_attempted"
	rows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "code_test_attempt")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> group(columns: ["question_id"])
		|> count()`,
	)
	if err != nil {
		return err
	}

	// terus langsung return hasilnya
	// tapi bisa juga di group per question, jadi
	// misalnya untuk question #1, dia ada 5 attempt, question #2 ada 10 attempt
	// and so on so forth.

	// Return the result here
	if rows.Record() == nil {
		result <- uint32(0)
		return nil
	}
	value, ok := rows.Record().Value().(int64)
	if !ok {
		log.Println("[ERROR] casting value to int64")
		result <- uint32(0)
		return nil
	}

	result <- uint32(value)

	return nil
}

func (d *Dependency) CalculateDeletionRate(ctx context.Context, sessionID uuid.UUID, result chan float32) error {
	var deletionTotal int64 = 0
	var totalKeystrokes int64 = 0
	var ok bool

	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// TODO:  ini buat ngambil nganu, jangan lupa result
	deletionRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "coding_event_keystroke")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn: (r) => (r["key_char"] == "backspace" or r["key_char"] == "delete"))
		|> count()
		|> yield(name: "count")`,
	)
	if err != nil {
		return err
	}
	defer deletionRows.Close()

	for deletionRows.Next() {
		deletionTotal, ok = deletionRows.Record().Value().(int64)
		if !ok {
			return errors.New("fail to infer deletion Total")
		}
	}

	keystrokeTotalRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> range(start: -1d)
		|> filter(fn: (r) => r["_measurement"] == "coding_event_keystroke")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn: (r) => (r._field == "key_char" and r._value != ""))
		|> count()
		|> yield(name: "count")`,
	)
	if err != nil {
		return (err)
	}
	defer keystrokeTotalRows.Close()

	for keystrokeTotalRows.Next() {
		value, ok := keystrokeTotalRows.Record().Value().(int64)
		if !ok {
			return errors.New("fail to infer keystroke Total")
		}

		totalKeystrokes = value
	}

	result <- (float32(deletionTotal) / float32(totalKeystrokes))

	// SELECT semua KeystrokeEvent WHERE value = delete OR value = backspace
	// terus jumlahin
	// dah gitu doang.

	// Return the result here
	return nil
}
