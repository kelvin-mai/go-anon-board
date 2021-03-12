package main

import (
	"github.com/kelvin-mai/go-anon-board/config"
	"github.com/kelvin-mai/go-anon-board/controllers"
	"github.com/kelvin-mai/go-anon-board/database"
	"github.com/kelvin-mai/go-anon-board/routes"
	"github.com/kelvin-mai/go-anon-board/services"
)

func main() {
	c := config.NewConfig()

	r := routes.NewRouter(c)
	conn := database.NewDatabaseConnection(c)

	ts := services.NewThreadService(conn)

	tc := controllers.NewThreadController(ts)

	r.RegisterThreadRoutes(tc)

	r.Serve()
}
