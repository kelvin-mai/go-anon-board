package router

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/controllers"
)

func notFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{"error": "route not found"})
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"ping": "pong"})
}

func newRouter() *gin.Engine {
	r := gin.Default()
	r.Use(helmet.DNSPrefetchControl())
	r.Use(helmet.FrameGuard())
	r.Use(helmet.Referrer("same-origin"))
	r.Static("/assets", "./assets")
	r.NoRoute(notFoundHandler)
	return r
}

func setUpRoutes(r *gin.Engine) {
	views := r.Group("/")
	controllers.RegisterHtmlRoutes(views)

	api := r.Group("/api/")
	api.GET("/", pingHandler)
	controllers.RegisterThreadRoutes(api.Group("/threads/:board"))
	controllers.RegisterReplyRoutes(api.Group("/replies/:board"))
}

func Init() {
	r := newRouter()
	r.LoadHTMLGlob("templates/*")
	setUpRoutes(r)
	r.Run()
}
