package main

import (
	"context"
	"encoding/json"
	"net/http"
	"runtime"
)

func FunFact(w http.ResponseWriter, r *http.Request) {
  runtime.GOMAXPROCS(2)

  type Member struct {
    MemId string "json:`member_id`"
  }

  var x Member

  err := json.NewDecoder(r.Body).Decode(&x)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  queryAPI := influxConn.QueryAPI(influxOrg)

  wpm := make(chan int8)
  delRate := make(chan float32)
  attempt := make(chan int8)

  go func() {
    // TODO:  ini buat ngambil nganu, jangan lupa result
    _, err := queryAPI.Query(context.Background(), "from()")
    if err != nil {
      panic(err)
    }

  }()
  // aggregate WPM

  go func() {
    // TODO:  ini buat ngambil nganu, jangan lupa result
    _, err := queryAPI.Query(context.Background(), "from()")
    if err != nil {
      panic(err)
    }

  }()
  // aggregate Delete keys

  go func() {
    // TODO:  ini buat ngambil nganu, jangan lupa result
    _, err := queryAPI.Query(context.Background(), "from()")
    if err != nil {
      panic(err)
    }

  }()
  // question attempt

  var result = struct {
    Wpm     int8
    DelRate float32
    Attempt int8
  }{
    <-wpm,
    <-delRate,
    <-attempt,
  }

  res, _ := json.Marshal(result)

  w.Write(res)

  return
}
