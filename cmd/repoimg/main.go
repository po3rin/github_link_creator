package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
	"time"

	"github.com/po3rin/github_link_creator/external"
	"github.com/po3rin/github_link_creator/pipeline"
)

const timeout = 30

var output = flag.String("o", "repo.png", "path to the output image")
var repo = flag.String("n", "", "specified full name of github repository")

func handleArg(arg string) (string, string, error) {
	if arg == "" {
		return "", "", errors.New("[-n] flag is required. format is 'user/repo'")
	}
	args := strings.Split(arg, "/")
	if len(args) != 2 {
		return "", "", fmt.Errorf("invalid args : %v", arg)
	}
	return args[0], args[1], nil
}
func main() {
	flag.Parse()
	user, repo, err := handleArg(*repo)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	doneCh := make(chan struct{})

	go func() {
		r := external.NewRepository()
		img, err := pipeline.ProcessingImg(ctx, r, user, repo)
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create(*output)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		err = png.Encode(file, img)
		if err != nil {
			log.Fatal(err)
		}
		doneCh <- struct{}{}
	}()

	select {
	case <-doneCh:
		return
	case <-ctx.Done():
		fmt.Printf("Processing timed out in %d seconds\n", timeout)
	}
}
