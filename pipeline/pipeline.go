package pipeline

import (
	"context"
	"image"
	"strconv"
	"strings"

	"github.com/po3rin/github_link_creator/config"
	"github.com/po3rin/github_link_creator/entity"
	l "github.com/po3rin/github_link_creator/lib/logger"
	"github.com/po3rin/github_link_creator/processing"
	"github.com/po3rin/img2circle"
)

// Repoitory is espipeline repository interface.
type Repoitory interface {
	GetRepoData(ctx context.Context, userName string, repoName string) (*entity.Repo, error)
	GetUserImage(ctx context.Context, avatarURL string) (image.Image, error)
}

// ProcessingImg processing Image for creating repository image.
func ProcessingImg(ctx context.Context, r Repoitory, userName string, repoName string) (image.Image, error) {
	// TODO - using gorutine
	repo, err := r.GetRepoData(ctx, userName, repoName)
	if err != nil {
		return nil, err
	}
	img, err := r.GetUserImage(ctx, repo.Owner.AvatarURL)
	if err != nil {
		return nil, err
	}
	cropper, err := img2circle.NewCropper(img2circle.Params{Src: processing.ResizeImg(img)})
	if err != nil {
		l.Error(err)
		return nil, err
	}
	synthesizedImg, err := processing.SynthesizeToBase(cropper.CropCircle())
	if err != nil {
		l.Error(err)
		return nil, err
	}

	config.Title.AdjustTitleFontSize(repo.Name)

	img, err = processing.DrawText(synthesizedImg, config.Title, repo.Name)
	if err != nil {
		l.Error(err)
		return nil, err
	}
	if len(repo.Description) < 45 {
		img, err = processing.DrawText(img, config.FirstDescription, repo.Description)
		if err != nil {
			l.Error(err)
			return nil, err
		}
	} else {
		desc := repo.Description
		if len(repo.Description) > 86 {
			desc = repo.Description[:86] + " ..."
		}
		words := strings.Split(desc, " ")
		var firstline, secondline string
		for _, w := range words {
			firstline += w + " "
			if len(firstline) >= 40 {
				img, err = processing.DrawText(img, config.FirstDescription, firstline)
				if err != nil {
					l.Error(err)
					return nil, err
				}
				secondline = strings.TrimPrefix(desc, firstline)
				img, err = processing.DrawText(img, config.SecondDescription, secondline)
				if err != nil {
					l.Error(err)
					return nil, err
				}
				break
			}
		}
	}
	img, err = processing.DrawText(img, config.Star, num2StringWithSINortion(repo.Stars))
	if err != nil {
		l.Error(err)
		return nil, err
	}
	img, err = processing.DrawText(img, config.Fork, num2StringWithSINortion(repo.Forks))
	if err != nil {
		l.Error(err)
		return nil, err
	}

	return img, nil
}

func num2StringWithSINortion(i int) string {
	if i > 999 {
		i = i / 1000
		return strconv.Itoa(i) + "k"
	}
	return strconv.Itoa(i)
}
