package env

import (
	"os"

	l "github.com/po3rin/github_link_creator/lib/logger"
)

var (
	// Port - API port.
	Port string
	// LogLevel - logging level.
	LogLevel string
)

const (
	// Timeout second of request timeout.
	Timeout = 30
)

const (
	defaultPort     = ":8080"
	defaultLogLevel = "DEBUG"
)

func init() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = defaultPort
		l.Warnf("PORT environment is not exists. use default value %s", defaultPort)
	}
	Port = port

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = defaultLogLevel
		l.Warnf("LOG_LEVEL environment is not exists. use default value %s", defaultLogLevel)
	}
	LogLevel = logLevel
}
