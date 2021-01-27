package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/controllers"
	"github.com/kelvin-mai/go-anon-board/providers"
	"github.com/kelvin-mai/go-anon-board/services"
)

func main() {
	conn, r := initializeProviders()

	ts := services.NewThreadService(conn)

	tc := controllers.NewThreadController(ts)

	api := r.Group("/api")
	threadRoutes := api.Group("/thread")

	registerThreadRoutes(threadRoutes, tc)

	r.Serve()
}

func initializeProviders() (providers.DatabaseConnection, providers.Router) {
	c := providers.NewConfig()
	conn := providers.NewDatabaseConnection(c)
	r := providers.NewRouter(c)
	return conn, r
}

func registerThreadRoutes(rg *gin.RouterGroup, c controllers.ThreadController) {
	rg.GET("/", c.ListThreads)
	rg.POST("/", c.CreateThread)
	rg.GET("/:id", c.GetThread)
	rg.PUT("/:id", c.ReportThread)
	rg.DELETE("/:id", c.DeleteThread)
}
