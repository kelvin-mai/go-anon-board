package models

import "github.com/jinzhu/gorm"

type Reply struct {
	gorm.Model
	Text           string `json:"text"`
	Reported       bool   `json:"reported" gorm:"default:false"`
	DeletePassword string `json:"-" gorm:"not null"`
	ThreadID       uint   `json:"-"`
}
