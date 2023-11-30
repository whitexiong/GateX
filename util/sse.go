package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SSEWriter struct {
	w       http.ResponseWriter
	flusher http.Flusher
}

func NewSSEWriter(w http.ResponseWriter) *SSEWriter {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	return &SSEWriter{
		w:       w,
		flusher: w.(http.Flusher),
	}
}

func (s *SSEWriter) SendEvent(data string) error {
	var parsedData interface{}
	if err := json.Unmarshal([]byte(data), &parsedData); err != nil {
		return err
	}

	formattedData, err := json.Marshal(parsedData)
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintf(s.w, "data: %s\n\n", string(formattedData)); err != nil {
		return err
	}
	s.flusher.Flush()

	return nil
}
