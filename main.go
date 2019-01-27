package main // import "github.com/po3rin/github_link_creator"

import (
	"context"
	_ "image/jpeg"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/po3rin/github_link_creator/external"
	"github.com/po3rin/github_link_creator/handler"
	"github.com/po3rin/github_link_creator/infrastructure"
	"github.com/po3rin/github_link_creator/lib/env"
	l "github.com/po3rin/github_link_creator/lib/logger"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	r := infrastructure.NewRouter()
	r.Handler = handler.Handler{
		Repo: external.NewRepository(),
	}
	router := r.InitRouter()

	srv := &http.Server{
		Addr:    env.Port,
		Handler: router,
	}
	go func() {
		l.Fatal(srv.ListenAndServe())
	}()

	// receive interrupt.
	killSignal := <-interrupt
	switch killSignal {
	case os.Kill:
		l.Info("Got SIGKILL...")
	case os.Interrupt:
		l.Info("Got SIGINT...")
	case syscall.SIGTERM:
		l.Info("Got SIGTERM...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	l.Info("The service is shutting down...")
	err := srv.Shutdown(ctx)
	if err != nil {
		l.Error("Failed to graceful shut down...")
	}
	l.Info("Done")
}
