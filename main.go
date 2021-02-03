package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/controllers"
	"github.com/kelvin-mai/go-anon-board/middlewares"
	"github.com/kelvin-mai/go-anon-board/providers"
	"github.com/kelvin-mai/go-anon-board/services"
)

func main() {
	c, conn, r := initializeProviders()

	ts := services.NewThreadService(conn)
	rs := services.NewReplyService(conn)

	tc := controllers.NewThreadController(ts)
	rc := controllers.NewReplyController(rs)
	ac := controllers.NewAdminController(ts, rs)

	api := r.Group("/api")
	threadRoutes := api.Group("/threads")
	replyRoutes := api.Group("/threads/:thread_id/replies")
	adminRoutes := api.Group("/admin")

	registerThreadRoutes(threadRoutes, tc)
	registerReplyRoutes(replyRoutes, rc)
	registerAdminRoutes(c, adminRoutes, ac)

	r.Serve()
}

func initializeProviders() (*providers.Config, providers.DatabaseConnection, providers.Router) {
	c := providers.NewConfig()
	conn := providers.NewDatabaseConnection(c)
	r := providers.NewRouter(c)

	return c, conn, r
}

func registerThreadRoutes(rg *gin.RouterGroup, c controllers.ThreadController) {
	rg.GET("/", c.ListThreads)
	rg.POST("/", c.CreateThread)
	rg.GET("/:id", c.GetThread)
	rg.PUT("/:id", c.ReportThread)
	rg.DELETE("/:id", c.DeleteThread)
}

func registerReplyRoutes(rg *gin.RouterGroup, c controllers.ReplyController) {
}

func registerAdminRoutes(c *providers.Config, rg *gin.RouterGroup, ac controllers.AdminController) {
	config := c.Get()
	apiKey := config.GetString("admin.api_key")
	rg.Use(middlewares.ApiKey("API-KEY", apiKey))

	rg.DELETE("/thread/:id", ac.DeleteThread)
}
