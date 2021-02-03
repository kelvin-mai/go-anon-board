package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/controllers"
	"github.com/kelvin-mai/go-anon-board/middlewares"
	"github.com/kelvin-mai/go-anon-board/response"
)

func setupDefaults(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		response.OK(c, gin.H{"health": "OK"})
	})
	r.NoRoute(func(c *gin.Context) {
		response.ResourceNotFound(c, nil)
	})
}

func (r *router) RegisterThreadRoutes(c controllers.ThreadController) {
	rg := r.Group("/api/threads")
	rg.GET("/", c.ListThreads)
	rg.POST("/", c.CreateThread)
	rg.GET("/:id", c.GetThread)
	rg.PUT("/:id", c.ReportThread)
	rg.DELETE("/:id", c.DeleteThread)
}

func (r *router) RegisterAdminRoutes(c controllers.AdminController) {
	apiKey := r.c.Get().GetString("admin.api_key")
	rg := r.Group("/api/admin")
	rg.Use(middlewares.ApiKey("API-KEY", apiKey))

	rg.DELETE("/thread/:id", c.DeleteThread)
	rg.DELETE("/replies/:id", c.DeleteReply)
}
