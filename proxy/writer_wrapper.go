package proxy

import "net/http"

type writerWrapper struct {
	http.ResponseWriter
	capturedBuffer []byte
	captured       int
	statusCode     int
	bytesWritten   int
}

func newWriterWrapper(w http.ResponseWriter, captureSize int) *writerWrapper {
	return &writerWrapper{
		ResponseWriter: w,
		capturedBuffer: make([]byte, captureSize),
		statusCode:     http.StatusOK,
	}
}

func (w *writerWrapper) Write(data []byte) (int, error) {

	if w.captured < len(w.capturedBuffer) {
		toCapture := len(data)
		canCapture := len(w.capturedBuffer) - w.captured
		if toCapture > canCapture {
			toCapture = canCapture
		}
		copy(w.capturedBuffer[w.captured:], data[:toCapture])
		w.captured += toCapture
	}
	w.bytesWritten += len(data)
	return w.ResponseWriter.Write(data)
}

func (w *writerWrapper) capturedData() []byte {
	return w.capturedBuffer[:w.captured]
}

func (w *writerWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
