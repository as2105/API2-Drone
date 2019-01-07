package logging

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// Fields ...
type Fields = logrus.Fields

// Logger ...
type Logger = logrus.Logger

// FieldLogger is an interface to abstract usage against a Logger or Entry
type FieldLogger = logrus.FieldLogger

// DefaultLogger ...
var DefaultLogger = logrus.StandardLogger()

// NewLogger ...
func NewLogger() *Logger {
	return DefaultLogger
}

// SetLevel ...
func SetLevel(log *Logger, level string) {
	if level, err := logrus.ParseLevel(level); err != nil {
		log.WithError(err).Warn("invalid log level provided, defaulting to INFO")
	} else {
		log.SetLevel(level)
	}
}

// SetFormatter ...
func SetFormatter(log *Logger, formatter string) {
	switch strings.ToLower(formatter) {
	case "json":
		log.SetFormatter(new(logrus.JSONFormatter))
	default:
		log.SetFormatter(new(logrus.TextFormatter))
	}
}
