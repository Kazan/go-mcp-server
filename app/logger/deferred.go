package logger

import (
	"bytes"
	"net/http"
)

// ResponseCapturer implements required methods from http.ResponseWriter in order
// to swap it effortlessly, while providing the capability for middlewares to
// add Header values even if the response Body handling has already started.
type ResponseCapturer struct {
	w    http.ResponseWriter
	buf  *bytes.Buffer
	code int
}

// NewResponseCapturer creates a writer wrapping the provided http.ResponseWriter
// allowing us to write to the stream, as well as capturing intermediate output.
func NewResponseCapturer(w http.ResponseWriter) *ResponseCapturer {
	return &ResponseCapturer{
		w:   w,
		buf: new(bytes.Buffer),
	}
}

// Header implements http.ResponseWriter.Header
func (rw *ResponseCapturer) Header() http.Header {
	return rw.w.Header()
}

// Write implements http.ResponseWriter.Write
func (rw *ResponseCapturer) Write(data []byte) (int, error) {
	rw.w.Write(data)

	return rw.buf.Write(data)
}

// WriteHeader implements http.ResponseWriter.WriteHeader
func (rw *ResponseCapturer) WriteHeader(statusCode int) {
	rw.code = statusCode
}

func (rw *ResponseCapturer) Buffer() *bytes.Buffer {
	return rw.buf
}
