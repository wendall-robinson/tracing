package loggers

import (
	"log"
)

// StdLogger is a type that adapts the standard log package to the Logger interface.
type StdLogger struct{}

// Printf logs a message.
func (s *StdLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Errorf logs an error message.
func (s *StdLogger) Errorf(format string, v ...interface{}) {
	log.Printf("ERROR: "+format, v...)
}

// Fatalf logs a fatal error message, then the process will exit.
func (s *StdLogger) Fatalf(format string, v ...interface{}) {
	log.Fatalf("FATAL: "+format, v...)
}

// Standard returns a Logger that uses the standard library's log package.
func Standard() Logger {
	return &StdLogger{}
}
