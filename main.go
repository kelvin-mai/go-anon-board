package main

import (
	"github.com/kelvin-mai/go-anon-board/controllers"
	"github.com/kelvin-mai/go-anon-board/models"
	"github.com/kelvin-mai/go-anon-board/providers"
	"github.com/kelvin-mai/go-anon-board/services"
)

func main() {
	db, r := initializeProviders()
	bs := initializeServices(db)
	bc := initializeControllers(bs)

	api := r.Group("/api")
	api.GET("/boards", bc.ListBoards)
	api.POST("/boards", bc.CreateBoard)
	api.GET("/boards/:id", bc.GetBoard)

	r.Serve()
}

func initializeProviders() (providers.DatabaseConnection, providers.Router) {
	c := providers.NewConfig()
	db := providers.NewDatabaseConnection(c)
	r := providers.NewRouter(c)
	db.Sync(c, &models.Board{}, &models.Thread{}, &models.Reply{})
	return db, r
}

func initializeServices(db providers.DatabaseConnection) services.BoardService {
	bs := services.NewBoardService(db)
	return bs
}

func initializeControllers(bs services.BoardService) controllers.BoardController {
	bc := controllers.NewBoardController(bs)
	return bc
}
