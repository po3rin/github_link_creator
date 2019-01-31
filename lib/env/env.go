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
	// GithubClientID - github client id.
	GithubClientID string
	// GithubSecret - github client id.
	GithubSecret string
	// S3BucketName - set target s3 bucket name.
	S3BucketName string
)

const (
	// Timeout second of request timeout.
	Timeout = 30
)

const (
	defaultPort         = ":8080"
	defaultLogLevel     = "DEBUG"
	defaultS3BucketName = "github-link-card"
)

func init() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = defaultPort
		l.Warnf("PORT environment is not exists. use default value %s", defaultPort)
	}
	Port = port

	githubClientID := os.Getenv("GITHUB_CLIENT_ID")
	if githubClientID == "" {
		l.Warnf("GITHUB_CLIENT_ID environment is not exists.")
	}
	GithubClientID = githubClientID

	githubSecret := os.Getenv("GITHUB_SECRET")
	if githubSecret == "" {
		l.Warnf("GITHUB_SECRET environment is not exists.")
	}
	GithubSecret = githubSecret

	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	if s3BucketName == "" {
		s3BucketName = defaultS3BucketName
		l.Warnf("S3_BUCKET_Name environment is not exists. use default value %s", defaultS3BucketName)
	}
	S3BucketName = s3BucketName
}
