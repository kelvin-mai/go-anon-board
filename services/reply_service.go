package services

import (
	"github.com/kelvin-mai/go-anon-board/database"
	"github.com/kelvin-mai/go-anon-board/models"
	"gorm.io/gorm"
)

type ReplyService interface {
	Delete(id string) error
}

type replyService struct {
	db *gorm.DB
}

func NewReplyService(conn database.DatabaseConnection) ReplyService {
	return &replyService{db: conn.GetDB()}
}

func (rs *replyService) Delete(id string) error {
	result := rs.db.Where("id = ?", id).Delete(&models.Reply{})
	return result.Error
}
