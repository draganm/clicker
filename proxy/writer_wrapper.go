package proxy

import "net/http"

type writerWrapper struct {
	http.ResponseWriter
	capturedBuffer []byte
	captureSize    int
	captured       int
	statusCode     int
}

func newWriterWrapper(w http.ResponseWriter, captureSize int) *writerWrapper {
	return &writerWrapper{
		ResponseWriter: w,
		captureSize:    captureSize,
		capturedBuffer: make([]byte, captureSize),
		statusCode:     http.StatusOK,
	}
}

func (w *writerWrapper) Write(data []byte) (int, error) {
	if w.captured < w.captureSize {
		toCapture := len(data)
		canCapture := w.captureSize - w.captured
		if toCapture > canCapture {
			toCapture = canCapture
		}
		copy(w.capturedBuffer[w.captured:], data[:toCapture])
		w.captured += toCapture
	}
	return w.ResponseWriter.Write(data)
}

func (w *writerWrapper) capturedData() []byte {
	return w.capturedBuffer[:w.captured]
}

func (w *writerWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
