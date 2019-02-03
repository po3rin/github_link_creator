package external

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"net/http"

	"github.com/po3rin/github_link_creator/entity"
	"github.com/po3rin/github_link_creator/lib/env"
	l "github.com/po3rin/github_link_creator/lib/logger"

	"github.com/pkg/errors"
)

var requestParams string

func init() {
	if env.GithubClientID == "" || env.GithubSecret == "" {
		l.Warn("not setted github client id or secret")
		requestParams = ""
		return
	}
	requestParams = fmt.Sprintf("?client_id=%v&client_secret=%v", env.GithubClientID, env.GithubSecret)
}

// GetRepoData request repository data.
func (r *Repository) GetRepoData(ctx context.Context, userName string, repoName string) (*entity.Repo, error) {
	uri := fmt.Sprintf("https://api.github.com/repos/%v/%v", userName, repoName)
	req, err := http.NewRequest("GET", uri+requestParams, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to new request, url: %v", uri))
	}
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to request url: %v, Please make sure repository exists", uri)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	repo := &entity.Repo{}
	err = json.Unmarshal(body, repo)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

// GetUserImage get image from github.
func (r *Repository) GetUserImage(ctx context.Context, avatarURL string) (image.Image, error) {
	req, err := http.NewRequest("GET", avatarURL+requestParams, nil)
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
