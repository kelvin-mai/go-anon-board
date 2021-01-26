package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/models"
	"github.com/kelvin-mai/go-anon-board/response"
	"github.com/kelvin-mai/go-anon-board/services"
)

type BoardController interface {
	ListBoards(c *gin.Context)
	GetBoard(c *gin.Context)
	CreateBoard(c *gin.Context)
}

type boardController struct {
	bs services.BoardService
}

func NewBoardController(bs services.BoardService) BoardController {
	return &boardController{bs: bs}
}

func (bc *boardController) ListBoards(c *gin.Context) {
	boards := bc.bs.List()
	if boards != nil {
		response.OK(c, boards)
		return
	}
	return
}

func (bc *boardController) GetBoard(c *gin.Context) {
	id := c.Param("id")
	board := bc.bs.Get(id)
	if board != nil {
		response.OK(c, board)
		return
	}
	response.ResourceNotFound(c, nil)
	return
}

func (bc *boardController) CreateBoard(c *gin.Context) {
	var b models.Board
	if err := c.ShouldBindJSON(&b); err != nil {
		response.BadRequest(c, errors.New("Invalid request body"))
		return
	}
	board := bc.bs.Create(b)
	response.Created(c, board)
	return
}
