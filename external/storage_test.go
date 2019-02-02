package external_test

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"testing"

	"github.com/po3rin/github_link_creator/external"
)

func TestUploadImg(t *testing.T) {
	r := external.NewRepository()
	img, err := os.Open("../images/gopher.jpg")
	if err != nil {
		t.Fatalf("unexpected error, err: %v", err.Error())
	}
	defer img.Close()
	src, _, err := image.Decode(img)
	if err != nil {
		t.Fatalf("unexpected error, err: %v", err.Error())
	}
	repoName := "po3rin/testtesttest"

	_, err = r.UploadImg(src, repoName)
	if err != nil {
		t.Fatalf("unexpected error, err: %v", err.Error())
	}
}
