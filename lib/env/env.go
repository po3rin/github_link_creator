package env

import (
	"os"
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
	}
	Port = port

	githubClientID := os.Getenv("GITHUB_CLIENT_ID")
	if githubClientID == "" {
	}
	GithubClientID = githubClientID

	githubSecret := os.Getenv("GITHUB_SECRET")
	if githubSecret == "" {
	}
	GithubSecret = githubSecret

	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	if s3BucketName == "" {
		s3BucketName = defaultS3BucketName
	}
	S3BucketName = s3BucketName
}
