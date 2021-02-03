package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/config"
	"github.com/kelvin-mai/go-anon-board/controllers"
	cors "github.com/rs/cors/wrapper/gin"
)

type Router interface {
	gin.IRouter
	Serve() error
	RegisterThreadRoutes(c controllers.ThreadController)
	RegisterAdminRoutes(c controllers.AdminController)
}

type router struct {
	*gin.Engine
	c *config.Config
}

func NewRouter(c *config.Config) Router {
	config := c.Get()
	r := gin.New()
	if config.GetString("ENVIRONMENT") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	if config.GetBool("app.log") {
		r.Use(gin.Logger())
	}
	r.Use(cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedHeaders: []string{"*"},
	}))

	setupDefaults(r)
	return &router{Engine: r, c: c}
}

func (r *router) Serve() error {
	port := r.c.Get().GetString("app.port")
	return r.Run(":" + port)
}
