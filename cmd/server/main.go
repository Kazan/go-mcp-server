package main

import (
	"net/http"

	"github.com/kazan/go-mcp-server/app/calculator"
	"github.com/kazan/go-mcp-server/app/logger"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Calculator Demo",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithRecovery(),
		server.WithLogging(),
	)

	calculator.Attach(s)

	// Start the server
	httpServer := server.NewStreamableHTTPServer(s)

	// Lets do some logging
	log := logger.NewLogger()
	handler := logger.LoggingMiddlewareFunc(log, httpServer)

	// Standard http server loop
	if err := http.ListenAndServe(":8080", handler); err != nil {
		if err != http.ErrServerClosed {
			log.Errorf("Failed to start server: %v\n", err)
			return
		}
		log.Infof("Server closed")
	}
}
