package loggers

import (
	"fmt"

	"github.com/golang/glog"
)

// GlogAdapter is a type that adapts a glog to the Logger interface.
type GlogAdapter struct{}

// Printf logs a formatted string at the Info level using glog.
func (g *GlogAdapter) Printf(format string, v ...interface{}) {
	glog.InfoDepth(1, fmt.Sprintf(format, v...))
}

// Errorf logs a formatted string at the Error level using glog.
func (g *GlogAdapter) Errorf(format string, v ...interface{}) {
	glog.ErrorDepth(1, fmt.Sprintf(format, v...))
}

// Fatalf logs a formatted string at the Fatal level using glog,
// then the process will exit.
func (g *GlogAdapter) Fatalf(format string, v ...interface{}) {
	glog.FatalDepth(1, fmt.Sprintf(format, v...))
}

// Glog takes no arguments and returns a Logger that uses glog.
func Glog() Logger {
	return &GlogAdapter{}
}
