package logger

import "fmt"

type StdLogger struct {
}

func (l *StdLogger) Infof(format string, v ...any) {
	fmt.Printf("INFO: "+format+"\n", v...)
}

func (l *StdLogger) Errorf(format string, v ...any) {
	fmt.Printf("ERROR: "+format+"\n", v...)
}
