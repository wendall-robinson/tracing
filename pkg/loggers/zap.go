package loggers

import (
	"go.uber.org/zap"
)

// ZapAdapter is a type that adapts a zap.SugaredLogger to the Logger interface.
type ZapAdapter struct {
	logger *zap.SugaredLogger
}

// Printf logs a formatted string at the Info level using a zap.SugaredLogger.
func (z *ZapAdapter) Printf(format string, v ...interface{}) {
	z.logger.Infof(format, v...)
}

// Errorf logs a formatted string at the Error level using a zap.SugaredLogger.
func (z *ZapAdapter) Errorf(format string, v ...interface{}) {
	z.logger.Errorf(format, v...)
}

// Fatalf logs a formatted string at the Fatal level using a zap.SugaredLogger,
// then the process will exit.
func (z *ZapAdapter) Fatalf(format string, v ...interface{}) {
	z.logger.Fatalf(format, v...)
}

// Zap takes a zap.SugaredLogger and returns a Logger that uses it.
func Zap(logger *zap.SugaredLogger) Logger {
	return &ZapAdapter{
		logger: logger,
	}
}
