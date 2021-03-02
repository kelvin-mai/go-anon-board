package main

import (
	"github.com/kelvin-mai/go-anon-board/config"
	"github.com/kelvin-mai/go-anon-board/routes"
)

func main() {
	c := config.NewConfig()
	r := routes.NewRouter(c)

	r.Serve()
}
