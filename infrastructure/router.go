package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/po3rin/github_link_creator/handler"
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
	router.GET("/v1/health", r.Handler.HealthCheck)
	router.GET("/v1/images/:user/:repo", r.Handler.GetCode)
	return router
}
