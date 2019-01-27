package pipeline

import (
	"context"
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/po3rin/github_link_creator/config"
	"github.com/po3rin/github_link_creator/entity"
	"github.com/po3rin/img2circle"
)

// Repoitory is espipeline repository interface.
type Repoitory interface {
	GetRepoData(ctx context.Context, userName string, repoName string) (*entity.Repo, error)
	GetUserData(ctx context.Context, userName string) (*entity.User, error)
	GetUserImage(ctx context.Context, avatarURL string) (image.Image, error)
}

// ProcessingImg processing Image for creating repository image.
func ProcessingImg(ctx context.Context, r Repoitory, userName string, repoName string) (image.Image, error) {
	// TODO - using gorutine
	repo, err := r.GetRepoData(ctx, userName, repoName)
	if err != nil {
		return nil, err
	}
	user, err := r.GetUserData(ctx, userName)
	if err != nil {
		return nil, err
	}
	img, err := r.GetUserImage(ctx, user.AvatarURL)
	if err != nil {
		return nil, err
	}
	cropper, err := img2circle.NewCropper(img2circle.Params{Src: ResizeImg(img)})
	if err != nil {
		return nil, err
	}
	synthesizedImg, err := SynthesizeToBase(cropper.CropCircle())
	if err != nil {
		return nil, err
	}

	img = DrawText(synthesizedImg, config.Title, repo.Name)

	if len(repo.Description) < 45 {
		img = DrawText(img, config.FirstDescription, repo.Description)
	} else {
		desc := repo.Description
		if len(repo.Description) > 86 {
			desc = repo.Description[:86] + " ..."
		}
		words := strings.Split(desc, " ")
		var firstline, secondline string
		for _, w := range words {
			firstline += w + " "
			fmt.Println(firstline)
			if len(firstline) >= 40 {
				img = DrawText(img, config.FirstDescription, firstline)
				secondline = strings.TrimPrefix(desc, firstline)
				img = DrawText(img, config.SecondDescription, secondline)
				break
			}
		}
	}

	img = DrawText(img, config.Star, strconv.Itoa(repo.Stars))
	img = DrawText(img, config.Fork, strconv.Itoa(repo.Forks))

	return img, nil
}
