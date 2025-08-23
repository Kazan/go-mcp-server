package logger

import (
	"bytes"
	"io"
	"net/http"
)

// DeferredWriter implements required methods from http.ResponseWriter in order
// to swap it effortlessly, while providing the capability for middlewares to
// add Header values even if the response Body handling has already started.
type DeferredWriter struct {
	w    http.ResponseWriter
	buf  bytes.Buffer
	code int
}

// NewDeferredWriter creates a deferred writer wrapping the provided http.ResponseWriter
func NewDeferredWriter(w http.ResponseWriter) *DeferredWriter {
	return &DeferredWriter{w: w}
}

// Header implements http.ResponseWriter.Header
func (rw *DeferredWriter) Header() http.Header {
	return rw.w.Header()
}

// Write implements http.ResponseWriter.Write
func (rw *DeferredWriter) Write(data []byte) (int, error) {
	return rw.buf.Write(data)
}

// WriteHeader implements http.ResponseWriter.WriteHeader
func (rw *DeferredWriter) WriteHeader(statusCode int) {
	rw.code = statusCode
}

func (rw *DeferredWriter) Buffer() bytes.Buffer {
	return rw.buf
}

// Done should be called to flush the response down the original ResponseWriter
// after all Headers have been set, and the response code determined
func (rw *DeferredWriter) Done() (int64, error) {
	if rw.code > 0 {
		rw.w.WriteHeader(rw.code)
	}

	return io.Copy(rw.w, &rw.buf)
}
