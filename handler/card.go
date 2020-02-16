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
	type resp struct {
		code int
		obj  interface{}
	}
	doneCh := make(chan resp)

	go func() {
		userName := c.Param("user")
		repoName := c.Param("repo")

		img, err := pipeline.ProcessingImg(ctx, h.Repo, userName, repoName)
		if err != nil {
			doneCh <- resp{
				code: http.StatusInternalServerError,
				obj: ErrorResponse{
					Message:          err.Error(),
					DocumentationURL: documentationURL,
				},
			}
			return
		}
		location, err := h.Repo.UploadImg(img, userName+"/"+repoName)
		if err != nil {
			l.Error(err)
			doneCh <- resp{
				code: http.StatusInternalServerError,
				obj: ErrorResponse{
					Message:          err.Error(),
					DocumentationURL: documentationURL,
				},
			}
			return
		}
		url := fmt.Sprintf("https://github.com/%v/%v", userName, repoName)
		result := fmt.Sprintf(`<a href="%v"><img src="%v" width="460px"></a>`, url, location)
		doneCh <- resp{
			code: http.StatusOK,
			obj: Response{
				Value:         result,
				RepositoryURL: url,
				CardURL:       location,
			},
		}
	}()

	select {
	case res := <-doneCh:
		c.JSON(res.code, res.obj)
	case <-ctx.Done():
		msg := fmt.Sprintf("Processing timed out in %d seconds", env.Timeout)
		l.Error(msg)
		c.JSON(http.StatusRequestTimeout, gin.H{
			"message": msg,
		})
	}
}
