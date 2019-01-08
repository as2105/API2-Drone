package logging

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// Fields is an alias of logrus.Fields
type Fields = logrus.Fields

// Logger is an alias of logrus.Logger
type Logger = logrus.Logger

// FieldLogger is an alias of logrus.FieldLogger
type FieldLogger = logrus.FieldLogger

// DefaultLogger is a global logger object
var DefaultLogger = logrus.StandardLogger()

// NewLogger returns the global logger object
func NewLogger() *Logger {
	return DefaultLogger
}

// SetLevel sets the global logging level
func SetLevel(log *Logger, level string) {
	if level, err := logrus.ParseLevel(level); err != nil {
		log.WithError(err).Warn("invalid log level provided, defaulting to INFO")
	} else {
		log.SetLevel(level)
	}
}

// SetFormatter sets the global logging format (only "json" is allowed, all other values result in the default logging format)
func SetFormatter(log *Logger, formatter string) {
	switch strings.ToLower(formatter) {
	case "json":
		log.SetFormatter(new(logrus.JSONFormatter))
	default:
		log.SetFormatter(new(logrus.TextFormatter))
	}
}
