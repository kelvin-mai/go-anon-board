package models

import (
	"time"
)

type Reply struct {
	ID             uint `json:"_id" gorm:"primary_key"`
	Text           string
	CreatedOn      time.Time
	Reported       bool `gorm:"default:false"`
	DeletePassword string
}
