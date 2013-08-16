package main

import (
  "net/http"
)

func ExtractApAiKeysFromRequest(r *http.Request) (searchApiKey string, displayApiKey string, err error) {
  values := r.URL.Query()

  searchApiKey = values.Get("search_api_key")
  displayApiKey = values.Get("display_api_key")

  if len(searchApiKey) == 0 || len(displayApiKey) == 0 {
    err = MissingApiKeys{}
  }

  return
}

func ExtractGidFromRequest(r *http.Request) (string, error) {
  params := r.URL.Query()
  gid := params.Get("gid")
  if len(gid) != 36 {
    return gid, InvalidGid{ gid }
  }

  return gid, nil
}
