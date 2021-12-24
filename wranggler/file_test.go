package main_test

import (
	main "rori"
	"testing"
	"time"
)

func TestConvertDataToJSON(t *testing.T) {
	data := []main.SampleInput{
		{
			Time:  time.Now().Add(time.Second * 1),
			Actor: "James",
			X:     20,
			Y:     13,
		},
		{
			Time:  time.Now().Add(time.Second * 2),
			Actor: "James",
			X:     21,
			Y:     13,
		},
		{
			Time:  time.Now().Add(time.Second * 3),
			Actor: "James",
			X:     22,
			Y:     14,
		},
	}
	_, err := main.ConvertDataToJSON(data)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}
}

func TestConvertDataToCSV(t *testing.T) {
	data := []main.SampleInput{
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

	res, err := main.ConvertDataToCSV(data)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	t.Log(string(res))
}
