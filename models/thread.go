package models

import (
	"time"
)

type Thread struct {
	ID             uint `json:"_id" gorm:"primary_key"`
	Text           string
	CreatedOn      time.Time
	BumpedOn       time.Time
	Reported       bool `gorm:"default:false"`
	DeletePassword string
}
