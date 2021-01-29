package services

import (
	"github.com/jinzhu/gorm"
	"github.com/kelvin-mai/go-anon-board/models"
	"github.com/kelvin-mai/go-anon-board/providers"
)

type ThreadService interface {
	List(offset int) (error, *[]models.Thread)
	GetByID(id string) (bool, *models.Thread)
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

func (ts *threadService) GetByID(id string) (bool, *models.Thread) {
	var t models.Thread
	result := ts.db.Where("id = ?", id).Preload("Replies").First(&t)
	return result.RecordNotFound(), &t
}

func (ts *threadService) Create(t models.Thread) (error, *models.Thread) {
	result := ts.db.Create(&t)
	return result.Error, &t
}

func (ts *threadService) Update(id string, t models.Thread) error {
	result := ts.db.Model(&t).Where("id = ?", id).Update(&t)
	return result.Error
}

func (ts *threadService) Delete(id string) error {
	result := ts.db.Where("id = ?", id).Delete(&models.Thread{})
	return result.Error
}
