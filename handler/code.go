package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/po3rin/github_link_creator/lib/env"
	l "github.com/po3rin/github_link_creator/lib/logger"
	"github.com/po3rin/github_link_creator/pipeline"

	"github.com/gin-gonic/gin"
)

// GetCode create GitHub Card Image.
func (h *Handler) GetCode(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), env.Timeout*time.Second)
	defer cancel()
	doneCh := make(chan struct{})

	go func() {
		userName := c.Param("user")
		repoName := c.Param("repo")

		img, err := pipeline.ProcessingImg(ctx, h.Repo, userName, repoName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message:          err.Error(),
				DocumentationURL: documentationURL,
			})
			l.Error(err)
			doneCh <- struct{}{}
			return
		}
		location, err := h.Repo.UploadImg(img, userName+"/"+repoName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message:          err.Error(),
				DocumentationURL: documentationURL,
			})
			l.Error(err)
			doneCh <- struct{}{}
			return
		}
		url := fmt.Sprintf("https://github.com/%v/%v", userName, repoName)
		result := fmt.Sprintf(`<a href="%v"><img src="%v" width="460px"></a>`, url, location)
		c.JSON(http.StatusOK, Response{
			Value:         result,
			RepositoryURL: url,
			CardURL:       location,
		})
		doneCh <- struct{}{}
	}()

	select {
	case <-doneCh:
		return
	case <-ctx.Done():
		msg := fmt.Sprintf("Processing timed out in %d seconds", env.Timeout)
		l.Error(msg)
		c.JSON(http.StatusRequestTimeout, gin.H{
			"message": msg,
		})
	}
}
