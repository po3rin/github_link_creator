package main // import "github.com/po3rin/github_link_creator"

import (
	_ "image/jpeg"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/po3rin/github_link_creator/external"
	"github.com/po3rin/github_link_creator/handler"
	"github.com/po3rin/github_link_creator/infrastructure"
	"github.com/po3rin/github_link_creator/lib/env"
	l "github.com/po3rin/github_link_creator/lib/logger"
)

func main() {
	r := infrastructure.NewRouter()
	r.Handler = handler.Handler{
		Repo: external.NewRepository(),
	}
	router := r.InitRouter()

	if os.Getenv("LAMBDA") == "true" {
		l.Fatal(gateway.ListenAndServe(env.Port, router))
	} else {
		l.Fatal(http.ListenAndServe(env.Port, router))
	}
}
