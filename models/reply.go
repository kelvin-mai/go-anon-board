package models

import (
	"time"
)

type Reply struct {
	ID             uint      `json:"_id" gorm:"primary_key"`
	Text           string    `json:"text"`
	CreatedOn      time.Time `json:"created_on" gorm:"default:now()"`
	Reported       bool      `json:"reported" gorm:"default:false"`
	DeletePassword string    `json:"-" gorm:"not null"`
	ThreadID       uint      `json:"-"`
}
