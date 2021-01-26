package main

import (
	"github.com/kelvin-mai/go-anon-board/models"
	"github.com/kelvin-mai/go-anon-board/providers"
)

func main() {
	c := providers.NewConfig()
	db := providers.NewDatabaseConnection(c)
	r := providers.NewRouter(c)

	db.Sync(c, &models.Board{}, &models.Thread{}, &models.Reply{})

	r.Serve()
}
