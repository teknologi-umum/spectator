package main

import (
	"encoding/json"
	"net/http"
)

func FunFact(w http.ResponseWriter, r *http.Request) {

	var x Member

	err := json.NewDecoder(r.Body).Decode(&x)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wpm := make(chan int8)
	delRate := make(chan float32)

	go func() {

	}()
	// aggregate WPM

	go func() {

	}()
	// aggregate Delete keys

	var result = struct {
		Wpm     int8
		DelRate float32
	}{
		<-wpm,
		<-delRate,
	}

	res, _ := json.Marshal(result)

	w.Write(res)

	return
}
