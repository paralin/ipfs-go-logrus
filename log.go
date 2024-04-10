// Package log is the logging library used by IPFS & libp2p
// (https://github.com/ipfs/go-ipfs).
package log

import (
	"time"
)

// StandardLogger provides API compatibility with standard printf loggers
// eg. go-logging
type StandardLogger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
}

// EventLogger extends the StandardLogger interface to allow for log items
// containing structured metadata
type EventLogger interface {
	StandardLogger
}

// Logger retrieves an event logger by name
func Logger(system string) *ZapEventLogger {
	if len(system) == 0 {
		setuplog := getLogger("setup-logger")
		setuplog.Error("Missing name parameter")
		system = "undefined"
	}

	logger := getLogger(system)

	return &ZapEventLogger{
		system:         system,
		StandardLogger: logger,
	}
}

// ZapEventLogger implements the EventLogger and wraps a go-logging Logger
type ZapEventLogger struct {
	StandardLogger
	system string
}

// Warning is for compatibility
// Deprecated: use Warn(args ...interface{}) instead
func (logger *ZapEventLogger) Warning(args ...interface{}) {
	logger.Warn(args...)
}

// Warningf is for compatibility
// Deprecated: use Warnf(format string, args ...interface{}) instead
func (logger *ZapEventLogger) Warningf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// FormatRFC3339 returns the given time in UTC with RFC3999Nano format.
func FormatRFC3339(t time.Time) string {
	return t.UTC().Format(time.RFC3339Nano)
}

// WithSkip returns a new logger that skips the specified number of stack frames when reporting the
// line/file.
func WithSkip(l *ZapEventLogger, skip int) *ZapEventLogger {
	// no-op
	copyLogger := *l
	return &copyLogger
}
