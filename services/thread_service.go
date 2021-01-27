package services

import (
	"github.com/jinzhu/gorm"
	"github.com/kelvin-mai/go-anon-board/models"
	"github.com/kelvin-mai/go-anon-board/providers"
)

type ThreadService interface {
	List(offset int) (error, *[]models.Thread)
	GetById(id string) (bool, *models.Thread)
	Create(t models.Thread) (error, *models.Thread)
	Update(id string, t models.Thread) error
}

type threadService struct {
	db *gorm.DB
}

func NewThreadService(conn providers.DatabaseConnection) ThreadService {
	return &threadService{db: conn.GetDB()}
}

func (ts *threadService) List(offset int) (error, *[]models.Thread) {
	var t []models.Thread
	result := ts.db.Limit(10).Offset(offset).Find(&t)
	return result.Error, &t
}

func (ts *threadService) GetById(id string) (bool, *models.Thread) {
	var t models.Thread
	result := ts.db.Where("id = ?", id).First(&t)
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
