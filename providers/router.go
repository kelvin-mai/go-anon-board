package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/response"
	cors "github.com/rs/cors/wrapper/gin"
)

type Router interface {
	gin.IRouter
	Serve() error
}

type router struct {
	*gin.Engine
	port string
}

func NewRouter(c *Config) Router {
	config := c.Get()
	r := gin.New()
	r.Use(gin.Recovery())

	if config.GetString("ENVIRONMENT") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	if config.GetBool("app.log") {
		r.Use(gin.Logger())
	}

	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedHeaders: []string{"*"},
	}))

	r.GET("/", func(c *gin.Context) {
		response.OK(c, gin.H{"health": "OK"})
	})

	r.NoRoute(func(c *gin.Context) {
		response.ResourceNotFound(c, nil)
	})

	port := config.GetString("app.port")
	return &router{Engine: r, port: port}
}

func (r *router) Serve() error {
	return r.Run(":" + r.port)
}
