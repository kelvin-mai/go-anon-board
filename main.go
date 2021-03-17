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
	rs := services.NewReplyService(conn)

	tc := controllers.NewThreadController(ts)
	rc := controllers.NewReplyController(rs)
	ac := controllers.NewAdminController(ts, rs)

	r.RegisterThreadRoutes(tc)
	r.RegisterReplyRoutes(rc)
	r.RegisterAdminRoutes(ac)

	r.Serve()
}
