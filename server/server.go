package server

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/controller"
)

func newRouter() *gin.Engine {
	r := gin.Default()
	r.Use(helmet.DNSPrefetchControl())
	r.Use(helmet.FrameGuard())
	r.Use(helmet.Referrer("same-origin"))
	r.NoRoute(notFoundHandler)
	return r
}

func notFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{"error": "route not found"})
}

func setupRoutes(r *gin.Engine) {
	r.GET("/health", controller.HealthCheckHandler)
}

func Init() {
	r := newRouter()
	setupRoutes(r)
	r.Run()
}
