// Package log is the logging library used by IPFS & libp2p
// (https://github.com/ipfs/go-ipfs).
package log

import (
	"time"

	"github.com/sirupsen/logrus"
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
		system: system,
		Entry:  logger,
	}
}

// ZapEventLogger implements the EventLogger and wraps a go-logging Logger
type ZapEventLogger struct {
	*logrus.Entry
	system string
}

// Warnw is for compatibility
// Deprecated: use Warn(args ...interface{}) instead
func (logger *ZapEventLogger) Warnw(args ...interface{}) {
	logger.Warn(args...)
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

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func (logger *ZapEventLogger) Debugw(msg string, keysAndValues ...interface{}) {
	// TODO
	logger.Debug(msg, keysAndValues)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (logger *ZapEventLogger) Infow(msg string, keysAndValues ...interface{}) {
	// TODO
	logger.Info(msg, keysAndValues)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (logger *ZapEventLogger) Errorw(msg string, keysAndValues ...interface{}) {
	logger.Error(msg, keysAndValues)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (logger *ZapEventLogger) DPanicw(msg string, keysAndValues ...interface{}) {
	logger.Panic(msg, keysAndValues)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (logger *ZapEventLogger) Panicw(msg string, keysAndValues ...interface{}) {
	logger.Panic(msg, keysAndValues)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (logger *ZapEventLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	logger.Fatal(msg, keysAndValues)
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
