package comm

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"time"
)

// Event represents HTTP request
type Event struct {
	Type         string
	UUID         string
	Time         time.Time
	RequestURI   string
	Header       http.Header
	CapturedBody []byte
	Method       string
}

// Encode encodes event to bytes
func (r Event) Encode() ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(r)
	return buf.Bytes(), err
}

// Decode decodes event data to Event object
func Decode(data []byte) (Event, error) {
	evt := Event{}
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&evt)
	return evt, err
}
