package main

import (
  "net/http"
  "encoding/json"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request){
  payload := jsonResponse{
    Error: false,
    Message: "Hit the broker",
  }

  _ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) sayHello(w http.ResponseWriter, r *http.Request){
  payload := jsonResponse{
    Error: false,
    Message: "Hello World From Go",
  }
  out, _ := json.MarshalIndent(payload, "", "\t")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusAccepted)
  w.Write(out)
}
