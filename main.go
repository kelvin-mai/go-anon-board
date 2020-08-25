package main

import (
	"github.com/kelvin-mai/go-anon-board/models"
	"github.com/kelvin-mai/go-anon-board/router"
)

func main() {
	db := models.Init()
	defer db.Close()
	router.Init()
}
