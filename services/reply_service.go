package services

import (
	"github.com/kelvin-mai/go-anon-board/models"
	"github.com/kelvin-mai/go-anon-board/providers"
	"gorm.io/gorm"
)

type ReplyService interface {
	List(offset int) (error, *[]models.Reply)
	GetByID(id string) (error, *models.Reply)
	Create(t models.Reply) (error, *models.Reply)
	Update(id string, t models.Reply) error
	Delete(id string) error
}

type replyService struct {
	db *gorm.DB
}

func NewReplyService(conn providers.DatabaseConnection) ReplyService {
	return &replyService{db: conn.GetDB()}
}

func (rs *replyService) List(offset int) (error, *[]models.Reply) {
	var r []models.Reply
	result := rs.db.Limit(10).Offset(offset).Find(&r)
	return result.Error, &r
}

func (rs *replyService) GetByID(id string) (error, *models.Reply) {
	var r models.Reply
	result := rs.db.Where("id = ?", id).First(&r)
	return result.Error, &r
}

func (rs *replyService) Create(r models.Reply) (error, *models.Reply) {
	result := rs.db.Create(&r)
	return result.Error, &r
}

func (rs *replyService) Update(id string, r models.Reply) error {
	result := rs.db.Model(&r).Where("id = ?", id).Updates(&r)
	return result.Error
}

func (rs *replyService) Delete(id string) error {
	result := rs.db.Where("id = ?", id).Delete(&models.Reply{})
	return result.Error
}
