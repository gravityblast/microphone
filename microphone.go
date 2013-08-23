package main

import (
  "github.com/pilu/lyricfind"
  "github.com/pilu/traffic"
  "fmt"
  "os"
  "net/http"
  "log"
  "database/sql"
)

const VERSION = "0.1.0"

type InvalidGid struct {
  Gid string
}

func (err InvalidGid) Error() string {
  return fmt.Sprintf("Invalid GID `%s`", err.Gid)
}

type MissingApiKeys struct {}

func (err MissingApiKeys) Error() string {
  return "Missing api keys"
}

type ResponseError struct {
  Message string
}

var DB *sql.DB
var lyricfindClient *lyricfind.Client
var router *traffic.Router

func init() {
  var err error
  dbString := os.Getenv("DB")
  DB, err = sql.Open("postgres", dbString)
  if err != nil {
    log.Fatal(err)
  }
  lyricfindClient = lyricfind.NewClient()

  router = traffic.New()
  router.AddBeforeFilter(SetResponseHeaders)
  router.Get("/", RootHandler)
  router.Get("/recordings/:gid", RecordingHandler)
  router.NotFoundHandler = NotFoundHandler
}

func main() {
  port := os.Getenv("PORT")
  host := os.Getenv("HOST")

  fmt.Printf("Starting on %s:%s\n", host, port)

  err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
