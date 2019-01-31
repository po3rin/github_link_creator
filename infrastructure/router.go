package infrastructure

import (
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
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
	router.Use(static.Serve("/", static.LocalFile("client/dist", true)))
	router.GET("/api/v1/health", r.Handler.HealthCheck)
	router.GET("/api/v1/images/:user/:repo", r.Handler.GetCode)
	router.GET("/api/v1/fdceda", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"val": os.Getenv("GITHUB_CLIENT_ID"),
		})
	})
	return router
}
