package external

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"net/http"

	"github.com/po3rin/github_link_creator/entity"

	"github.com/pkg/errors"
)

// GetRepoData request repository data.
func (r *Repository) GetRepoData(ctx context.Context, userName string, repoName string) (*entity.Repo, error) {
	uri := fmt.Sprintf("https://api.github.com/repos/%v/%v", userName, repoName)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to new request, url: %v", uri))
	}
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	repo := &entity.Repo{}
	json.Unmarshal(body, repo)
	return repo, nil
}

// GetUserData request user data.
func (r *Repository) GetUserData(ctx context.Context, userName string) (*entity.User, error) {
	uri := fmt.Sprintf("https://api.github.com/users/%v", userName)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to new request, url: %v", uri))
	}
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	user := &entity.User{}
	json.Unmarshal(body, user)
	return user, nil
}

// GetUserImage get image from github.
func (r *Repository) GetUserImage(ctx context.Context, avatarURL string) (image.Image, error) {
	req, err := http.NewRequest("GET", avatarURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to create NewRequest, url: %v", avatarURL))
	}
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to request, url: %v", avatarURL))
	}
	defer res.Body.Close()
	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to decode image, url: %v", avatarURL))
	}
	return img, nil
}
