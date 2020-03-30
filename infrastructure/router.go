package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/po3rin/github_link_creator/handler"
	"golang.org/x/time/rate"
)

type Infrastructure interface {
	InitRouter() *gin.Engine
}

type Router struct {
	Handler handler.Handler
}

func NewRouter() *Router {
	return &Router{}
}

// InitRouter provide initialized　router
func (r *Router) InitRouter() *gin.Engine {
	router := gin.Default()

	l := rate.NewLimiter(rate.Limit(1), 50)
	router.Use(limiter(l))

	router.GET("/api/v1/health", r.Handler.HealthCheck)
	router.GET("/api/v1/images/:user/:repo", r.Handler.GetCode)
	return router
}
