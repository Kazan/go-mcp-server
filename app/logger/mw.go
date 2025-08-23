package logger

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/mark3labs/mcp-go/server"
)

func LoggingMiddlewareFunc(l *StdLogger, next *server.StreamableHTTPServer) http.HandlerFunc {
	calls := 0

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++

		rbody, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(rbody))

		callID := fmt.Sprintf("%08d", calls)
		l.Infof("Request [%s]: %s %s, body:\n%s", callID, r.Method, r.URL.Path, string(rbody))

		rw := NewDeferredWriter(w)
		// Call the next handler
		next.ServeHTTP(rw, r)

		// Log the outgoing response
		buf := rw.Buffer()
		l.Infof("Response [%s]:\n%s", callID, buf.String())

		rw.Done()
	})
}
