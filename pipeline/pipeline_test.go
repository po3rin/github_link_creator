package pipeline_test

import (
	"context"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"testing"

	"github.com/po3rin/github_link_creator/entity"
	"github.com/po3rin/github_link_creator/pipeline"
)

type testRepoitory struct{}

func (t *testRepoitory) GetRepoData(ctx context.Context, userName string, repoName string) (*entity.Repo, error) {
	return &entity.Repo{
		Name:        "po3rin",
		URL:         "example.com",
		Description: "this is test",
		Forks:       10,
		Stars:       10,
		Owner: entity.User{
			AvatarURL: "image.example.com",
		},
	}, nil
}
func (t *testRepoitory) GetUserImage(ctx context.Context, avatarURL string) (image.Image, error) {
	img, err := os.Open("../images/gopher.jpg")
	if err != nil {
		return nil, err
	}
	defer img.Close()
	src, _, err := image.Decode(img)
	if err != nil {
		return nil, err
	}
	return src, nil
}

func TestProcessingImg(t *testing.T) {
	r := &testRepoitory{}
	ctx := context.Background()
	userName := "po3rin"
	repoName := "testrepo"

	_, err := pipeline.ProcessingImg(ctx, r, userName, repoName)
	if err != nil {
		t.Fatalf("unexpected error, %v", err.Error())
	}
}
