package log

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// LogLevel represents a log severity level. Use the package variables as an
// enum.
type LogLevel = logrus.Level

var (
	LevelDebug  = LogLevel(logrus.DebugLevel)
	LevelInfo   = LogLevel(logrus.InfoLevel)
	LevelWarn   = LogLevel(logrus.WarnLevel)
	LevelError  = LogLevel(logrus.ErrorLevel)
	LevelDPanic = LogLevel(logrus.PanicLevel)
	LevelPanic  = LogLevel(logrus.PanicLevel)
	LevelFatal  = LogLevel(logrus.FatalLevel)
)

// LevelFromString parses a string-based level and returns the corresponding
// LogLevel.
//
// Supported strings are: DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL, and
// their lower-case forms.
//
// The returned LogLevel must be discarded if error is not nil.
func LevelFromString(level string) (LogLevel, error) {
	level = strings.ToLower(level)
	if level == "dpanic" {
		level = "panic"
	}
	return logrus.ParseLevel(level)
}
