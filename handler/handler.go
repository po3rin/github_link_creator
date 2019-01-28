package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/po3rin/github_link_creator/lib/env"
	"github.com/po3rin/github_link_creator/pipeline"

	"github.com/gin-gonic/gin"
)

// GetImage create GitHub Card Image.
func (h *Handler) GetCode(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), env.Timeout*time.Second)
	defer cancel()
	doneCh := make(chan struct{})

	go func() {
		userName := c.Param("user")
		repoName := c.Param("repo")

		img, err := pipeline.ProcessingImg(ctx, h.Repo, userName, repoName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		location, err := h.Repo.UploadImg(img, userName+"/"+repoName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		result := fmt.Sprintf(`<a href="https://github.com/%v/%v"><img src="%v" width="460px"></a>`, userName, repoName, location)
		c.JSON(http.StatusOK, gin.H{
			"value": result,
		})
		doneCh <- struct{}{}
	}()

	select {
	case <-doneCh:
		return
	case <-ctx.Done():
		c.JSON(http.StatusRequestTimeout, gin.H{
			"message": fmt.Sprintf("Processing timed out in %d seconds", env.Timeout),
		})
	}
}
