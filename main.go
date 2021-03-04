package main

import (
	"fmt"

	"github.com/kelvin-mai/go-anon-board/config"
	"github.com/kelvin-mai/go-anon-board/database"
	"github.com/kelvin-mai/go-anon-board/routes"
)

func main() {
	c := config.NewConfig()

	r := routes.NewRouter(c)
	conn := database.NewDatabaseConnection(c)

	fmt.Println(conn)

	r.Serve()
}
