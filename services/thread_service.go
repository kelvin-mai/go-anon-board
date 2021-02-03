package services

import (
	"github.com/kelvin-mai/go-anon-board/models"
	"github.com/kelvin-mai/go-anon-board/providers"
	"gorm.io/gorm"
)

type ThreadService interface {
	List(offset int) (error, *[]models.Thread)
	GetByID(id string) (error, *models.Thread)
	Create(t models.Thread) (error, *models.Thread)
	Update(id string, t models.Thread) error
	Delete(id string) error
}

type threadService struct {
	db *gorm.DB
}

func NewThreadService(conn providers.DatabaseConnection) ThreadService {
	return &threadService{db: conn.GetDB()}
}

func (ts *threadService) List(offset int) (error, *[]models.Thread) {
	var t []models.Thread
	result := ts.db.Preload("Replies").Limit(10).Offset(offset).Find(&t)
	return result.Error, &t
}

func (ts *threadService) GetByID(id string) (error, *models.Thread) {
	var t models.Thread
	result := ts.db.Where("id = ?", id).Preload("Replies").First(&t)
	return result.Error, &t
}

func (ts *threadService) Create(t models.Thread) (error, *models.Thread) {
	result := ts.db.Create(&t)
	return result.Error, &t
}

func (ts *threadService) Update(id string, t models.Thread) error {
	result := ts.db.Model(&t).Where("id = ?", id).Updates(&t)
	return result.Error
}

func (ts *threadService) Delete(id string) error {
	result := ts.db.Where("id = ?", id).Delete(&models.Thread{})
	return result.Error
}
