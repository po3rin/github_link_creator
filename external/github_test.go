package external_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/po3rin/github_link_creator/entity"
	"github.com/po3rin/github_link_creator/external"
)

func TestGetRepoData(t *testing.T) {
	tests := []struct {
		userName  string
		repoName  string
		expecterr bool
		want      *entity.Repo
	}{
		{
			userName: "po3rin",
			repoName: "gotree",
			want: &entity.Repo{
				Name:        "po3rin/gotree",
				URL:         "https://github.com/po3rin/gotree",
				Description: "Package gotree lets you display tree with go package document. This helps to understand the role of the package. hidden directory and File is not yet supported.",
				Forks:       0,
				Stars:       3,
				Owner: entity.User{
					AvatarURL: "https://avatars.githubusercontent.com/u/29445112?v=4",
				},
			},
			expecterr: false,
		},
		{
			userName:  "po3rin",
			repoName:  "fdfdfdfdfdfdf",
			want:      &entity.Repo{},
			expecterr: true,
		},
		{
			userName:  "fsdlivhidhvsdiabdijbisdvsdvdvdvd",
			repoName:  "gotree",
			want:      &entity.Repo{},
			expecterr: true,
		},
	}
	ctx := context.Background()
	r := external.NewRepository()
	for _, c := range tests {
		got, err := r.GetRepoData(ctx, c.userName, c.repoName)
		if err != nil {
			if c.expecterr {
				continue
			}
			t.Errorf("unexpected error. err: %v", err.Error())
			continue
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("unexpected error. got: %+v, want: %+v", got, c.want)
		}
	}

}
