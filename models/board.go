package models

import "github.com/jinzhu/gorm"

type Board struct {
	gorm.Model
	Title   string `gorm:"unique"`
	Threads []Thread
}
