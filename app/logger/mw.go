package logger

import (
	"bytes"
	"encoding/json"
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
		headers, _ := json.Marshal(r.Header)
		l.Infof("Request [%s]: %s %s\nheaders:\n%s\nbody:\n%s", callID, r.Method, r.URL.Path, string(headers), string(rbody))

		rw := NewDeferredWriter(w)
		// Call the next handler
		next.ServeHTTP(rw, r)

		// Log the outgoing response
		buf := rw.Buffer()
		l.Infof("Response [%s]:\n%s", callID, buf.String())

		rw.Done()
	})
}
