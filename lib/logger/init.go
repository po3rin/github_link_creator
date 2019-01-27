package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log - logrus struct
var Log *logrus.Logger

func init() {
	Log = logrus.New()
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		Log.Level = logrus.DebugLevel
	case "INFO":
		Log.Level = logrus.InfoLevel
	case "WARN":
		Log.Level = logrus.WarnLevel
	case "ERROR":
		Log.Level = logrus.ErrorLevel
	default:
		Log.Level = logrus.ErrorLevel
	}

	Log.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: true}
}