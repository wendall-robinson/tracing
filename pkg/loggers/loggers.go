package loggers

// Logger is an interface that defines methods for logging with different levels.
// It is designed to be satisfied by various logging packages.
type Logger interface {
	// Printf logs a formatted string at the Info level.
	Printf(format string, v ...interface{})

	// Errorf logs a formatted string at the Error level.
	Errorf(format string, v ...interface{})

	// Fatalf logs a formatted string at the Fatal level, then the process will exit.
	Fatalf(format string, v ...interface{})
}
