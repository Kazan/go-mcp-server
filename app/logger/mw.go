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
		l.Infof("Request [%s]: %s %s\nbody:\n%s", callID, r.Method, r.URL.Path, string(rbody))

		// capture response body for logging purposes
		rw := NewResponseCapturer(w)

		// Call the next handler
		next.ServeHTTP(rw, r)

		// Log the outgoing response
		buf := rw.Buffer()
		var j map[string]any
		_ = json.Unmarshal(buf.Bytes(), &j)
		outBody, _ := json.MarshalIndent(j, "", "  ")

		l.Infof("Response [%s]:\n%s", callID, outBody)
	})
}
