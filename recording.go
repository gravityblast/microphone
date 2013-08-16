package main

import (
  "database/sql"
  _ "github.com/bmizerany/pq"
  "github.com/pilu/lyricfind"
)

type Recording struct {
  Artist string
  Track string
}

type RecordingResponse struct {
  Amg int           `json:"amg"`
  Instrumental bool `json:"instrumental"`
  Viewable bool     `json:"viewable"`
  Has_lrc bool      `json:"has_lrc"`
  Title string      `json:"title"`
  ArtistName string `json:"artist_name"`
  Lyrics string     `json:"lyrics"`
  Copyright string  `json:"copyright"`
  Writer string     `json:"writer"`
}

func FindRecording(db *sql.DB, gid string) (*Recording, error) {
  query := `select AN.name as artist_name, TN.name as track_name from
  recording R, artist_credit AC, artist_name AN, track_name TN
  where AC.id = R.artist_credit and AC.name = AN.id and R.name = TN.id and
  R.gid = $1 limit 1;`

  recording := &Recording{}

  err := db.QueryRow(query, gid).Scan(&recording.Artist, &recording.Track)

  return recording, err
}

func BuildRecordingResponse(lyricsResponse lyricfind.LyricsResponse) *RecordingResponse {
  response := &RecordingResponse{}
  response.Amg = lyricsResponse.Track.Amg
  response.Instrumental = lyricsResponse.Track.Instrumental
  response.Viewable = lyricsResponse.Track.Viewable
  response.Has_lrc = lyricsResponse.Track.Has_lrc
  response.Title = lyricsResponse.Track.Title
  response.ArtistName = lyricsResponse.Track.Artist.Name
  response.Lyrics = lyricsResponse.Track.Lyrics
  response.Copyright = lyricsResponse.Track.Copyright
  response.Writer = lyricsResponse.Track.Writer

  return response
}
