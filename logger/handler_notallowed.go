package main

import "net/http"

func NotAllowed(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Header().Set("content-type", "text/plain")
	_, err := w.Write([]byte("method not allowed"))
	if err != nil {
		return err
	}

	return nil
}
