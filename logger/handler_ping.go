package main

import (
	"fmt"
	"net/http"
)

func (d *Dependency) Ping(w http.ResponseWriter, r *http.Request) error {
	health, err := d.DB.Health(r.Context())
	if err != nil {
		return fmt.Errorf("health check call: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "text/plain")
	_, err = w.Write([]byte(string(health.Status)))
	if err != nil {
		return fmt.Errorf("error writing data into http: %v", err)
	}

	return nil
}
