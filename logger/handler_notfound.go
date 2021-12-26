package main

import "net/http"

func NotFound(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("content-type", "text/plain")
	_, err := w.Write([]byte("not found"))
	if err != nil {
		return err
	}

	return nil
}
