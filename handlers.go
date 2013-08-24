package main

import (
  "net/http"
  "encoding/json"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
  RenderNotFound(w)
}

func SetResponseHeaders(w http.ResponseWriter, r *http.Request) bool {
  w.Header().Set("Content-Type", "application/json")

  return true
}

func RenderNotFound(w http.ResponseWriter) {
  err := ResponseError{ "Sorry, that page does not exist" }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusNotFound)
  json.NewEncoder(w).Encode(err)
}
