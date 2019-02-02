package logger

import (
	"os"

	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/raven-go"
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
		Log.Level = logrus.WarnLevel
	}

	Log.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: true}

	sentryDSN := os.Getenv("SENTRY_DSN")
	if sentryDSN == "" {
		return
	}
	client, err := raven.New(sentryDSN)
	if err != nil {
		Fatal(err)
	}
	hook, err := logrus_sentry.NewWithClientSentryHook(client, []logrus.Level{
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	})
	if err == nil {
		Log.Hooks.Add(hook)
	}
}
