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
	// GithubClientID - github client id
	GithubClientID string
	// GithubSecret - github client id
	GithubSecret string
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

	githubClientID := os.Getenv("GITHUB_CLIRNT_ID")
	if githubClientID == "" {
		l.Warnf("GITHUB_CLIRNT_ID environment is not exists.")
	}
	GithubClientID = githubClientID

	githubSecret := os.Getenv("GITHUB_SECRET")
	if githubSecret == "" {
		l.Warnf("GITHUB_SECRET environment is not exists.")
	}
	GithubSecret = githubSecret
}
