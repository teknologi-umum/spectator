package main

import (
	"log"
	"net/http"
)

func HandleError(e error, w http.ResponseWriter, r *http.Request) {
	log.Println(e)
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("content-type", "text/plain")
	_, err := w.Write([]byte(e.Error()))
	if err != nil {
		log.Printf("error writing data into server: %v", err)
	}
}
