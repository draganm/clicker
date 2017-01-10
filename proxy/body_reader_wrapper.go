package proxy

import (
	"io"
	"time"
)

type bodyReaderWrapper struct {
	io.ReadCloser
	captured       []byte
	capturedLength int
	bytesRead      int
	lastReadAt     time.Time
}

func newBodyReaderWrapper(bodyReader io.ReadCloser, captureSize int) *bodyReaderWrapper {
	return &bodyReaderWrapper{
		ReadCloser:     bodyReader,
		captured:       make([]byte, captureSize),
		capturedLength: 0,
	}
}

func (r *bodyReaderWrapper) Read(dest []byte) (int, error) {

	n, err := r.ReadCloser.Read(dest)
	r.lastReadAt = time.Now()
	toCapture := len(r.captured) - r.capturedLength
	if toCapture > 0 {
		if n < toCapture {
			toCapture = n
		}
		copy(r.captured[r.capturedLength:], dest[:toCapture])
		r.capturedLength += toCapture
	}
	r.bytesRead += n

	return n, err
}

func (r *bodyReaderWrapper) capturedData() []byte {
	return r.captured[:r.capturedLength]
}
