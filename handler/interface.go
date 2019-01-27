package handler

import (
	"context"
	"image"

	"github.com/po3rin/github_link_creator/entity"
)

// Handler - handle API request.
type Handler struct {
	Repo Repo
}

// Repo - has repository methods to separate dependencies.
type Repo interface {
	GetRepoData(ctx context.Context, userName string, repoName string) (*entity.Repo, error)
	GetUserData(ctx context.Context, userName string) (*entity.User, error)
	GetUserImage(ctx context.Context, avatarURL string) (image.Image, error)
}
