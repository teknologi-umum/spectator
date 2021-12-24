package main

import (
	"net/http"
)

type Member struct {
	ID string "json:`member_id`"
}

func GenerateFile(w http.ResponseWriter, r *http.Request) {
	return
}
