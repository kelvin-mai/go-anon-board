package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/controllers"
	"github.com/kelvin-mai/go-anon-board/utils"
)

func setupDefaults(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "OK"})
	})
	r.NoRoute(func(c *gin.Context) {
		utils.CreateApiError(http.StatusNotFound, errors.New("no route found"))
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
	rg.Use(ApiKey("API-KEY", apiKey))

	rg.DELETE("/thread/:id", c.DeleteThread)
	rg.DELETE("/replies/:id", c.DeleteReply)
}
