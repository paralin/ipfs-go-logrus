package log

import (
	"errors"
	"io"
	"sync"

	"github.com/sirupsen/logrus"
)

// ErrNoSuchLogger is returned when the util pkg is asked for a non existent logger
var ErrNoSuchLogger = errors.New("error: No such logger")

var loggerMutex sync.Mutex // guards access to global logger state

// loggerCore is the primary logging core
var loggerCore *logrus.Entry

type nullFormatter struct{}

func (f *nullFormatter) Format(*logrus.Entry) ([]byte, error) {
	return nil, nil
}

func NewNoopLogger() *logrus.Entry {
	log := logrus.New()
	log.Out = io.Discard
	log.Formatter = &nullFormatter{}
	return logrus.NewEntry(log)
}

// GetConfig returns a copy of the saved config. It can be inspected, modified,
// and re-applied using a subsequent call to SetupLogging().
func GetLogger() *logrus.Entry {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()

	if loggerCore == nil {
		loggerCore = NewNoopLogger()
	}

	return loggerCore
}

// SetupLogging will initialize the logger backend and set the flags.
// TODO calling this in `init` pushes all configuration to env variables
// - move it out of `init`? then we need to change all the code (js-ipfs, go-ipfs) to call this explicitly
// - have it look for a config file? need to define what that is
func SetupLogging(le *logrus.Entry) {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()

	if le == nil {
		le = NewNoopLogger()
	}

	loggerCore = le
}

// SetDebugLogging calls SetAllLoggers with logging.DEBUG
func SetDebugLogging() {
	SetAllLoggers(LevelDebug)
}

// SetAllLoggers changes the logging level of all loggers to lvl
func SetAllLoggers(lvl LogLevel) {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()

	if loggerCore == nil {
		loggerCore = NewNoopLogger()
	}

	loggerCore.Level = lvl
}

func getLogger(name string) *logrus.Entry {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()

	if loggerCore == nil {
		loggerCore = NewNoopLogger()
	}

	return loggerCore.WithField("logger-name", name)
}
