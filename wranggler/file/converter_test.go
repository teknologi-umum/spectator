package file_test

import (
	"testing"
	"time"
	"worker/file"
)

type SampleInput struct {
	Time  time.Time
	Actor string
	X     int
	Y     int
}

func TestConvertDataToCSV(t *testing.T) {
	data := []SampleInput{
		{
			Time:  time.Now().Add(time.Second * 1),
			Actor: "James",
			X:     20,
			Y:     13,
		},
		{
			Time:  time.Now().Add(time.Second * 2),
			Actor: "James, Riyadi",
			X:     21,
			Y:     13,
		},
		{
			Time:  time.Now().Add(time.Second * 3),
			Actor: "James Riyadi",
			X:     22,
			Y:     14,
		},
	}

	res, err := file.ConvertDataToCSV(data)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	t.Log(string(res))
}
