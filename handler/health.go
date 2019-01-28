package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/po3rin/github_link_creator/lib/env"

	"github.com/gin-gonic/gin"
)

// HealthCheck checks health.
func (h *Handler) HealthCheck(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), env.Timeout*time.Second)
	defer cancel()
	doneCh := make(chan struct{})

	go func() {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok!!",
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
