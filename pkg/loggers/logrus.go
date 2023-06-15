package loggers

import (
	"github.com/sirupsen/logrus"
)

// LogrusAdapter is a type that adapts a logrus.Logger to the Logger interface.
type LogrusAdapter struct {
	logger *logrus.Logger
}

// Printf logs a formatted string at the Info level using a logrus.Logger.
func (l *LogrusAdapter) Printf(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

// Errorf logs a formatted string at the Error level using a logrus.Logger.
func (l *LogrusAdapter) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

// Fatalf logs a formatted string at the Fatal level using a logrus.Logger,
// then the process will exit.
func (l *LogrusAdapter) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}

// Logrus takes a logrus.Logger and returns a Logger that uses it.
func Logrus(logger *logrus.Logger) Logger {
	return &LogrusAdapter{
		logger: logger,
	}
}
