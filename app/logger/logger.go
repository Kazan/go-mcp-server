package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

type StdLogger struct {
	l *slog.Logger
}

func NewLogger() *StdLogger {
	w := os.Stdout

	// Create a new logger
	logger := slog.New(tint.NewHandler(w, nil))

	return &StdLogger{l: logger}
}

func (l *StdLogger) Infof(format string, v ...any) {
	l.l.Info(fmt.Sprintf(format, v...))
}

func (l *StdLogger) Errorf(format string, v ...any) {
	l.l.Error(fmt.Sprintf(format, v...))
}
