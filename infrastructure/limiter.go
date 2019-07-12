package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	l "github.com/po3rin/github_link_creator/lib/logger"
	"golang.org/x/time/rate"
)

func limiter(r *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := r.Wait(c); err != nil {
			l.Warnf("To many requests, %v", c.Request)
			c.JSON(http.StatusTooManyRequests, "To many requests, Please wait for a while")
			c.Abort()
		}
		c.Next()
	}
}
