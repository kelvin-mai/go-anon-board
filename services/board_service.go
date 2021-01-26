package services

import (
	"github.com/kelvin-mai/go-anon-board/models"
	"github.com/kelvin-mai/go-anon-board/providers"
)

type BoardService interface {
	List() []models.Board
	Get(id string) *models.Board
	Create(b models.Board) models.Board
}

type boardService struct {
	conn providers.DatabaseConnection
}

func NewBoardService(conn providers.DatabaseConnection) BoardService {
	return &boardService{conn: conn}
}

func (bs *boardService) List() []models.Board {
	db := bs.conn.GetDB()
	var b []models.Board
	db.Find(&b)
	return b
}

func (bs *boardService) Get(id string) *models.Board {
	db := bs.conn.GetDB()
	var b models.Board
	result := db.Where("id = ?", id).First(&b)
	if result.RecordNotFound() {
		return nil
	}
	return &b
}

func (bs *boardService) Create(b models.Board) models.Board {
	db := bs.conn.GetDB()
	db.Create(&b)
	return b
}
