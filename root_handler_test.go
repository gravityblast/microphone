package main

import (
  "testing"
  "fmt"
  assert "github.com/pilu/miniassert"
)

func TestRootHandler(t *testing.T) {
  recorder := newTestRequest("GET", "/")
  expectedBody := fmt.Sprintf(`{"version":"%s"}`, VERSION) + "\n"

  assert.Equal(t, []string{"application/json"}, recorder.HeaderMap["Content-Type"])
  assert.Equal(t, expectedBody, string(recorder.Body.Bytes()))
  assert.Equal(t, 200, recorder.Code)
}
