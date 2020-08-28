package models

import (
	"time"
)

type Thread struct {
	ID             uint      `json:"_id" gorm:"primary_key"`
	Text           string    `json:"text"`
	CreatedOn      time.Time `json:"created_on" gorm:"default:now()"`
	BumpedOn       time.Time `json:"bumped_on" gorm:"default:now()"`
	Reported       bool      `json:"reported" gorm:"default:false"`
	DeletePassword string    `json:"-" gorm:"not null"`
	BoardID        uint      `json:"-"`
	Replies        []Reply   `json:"replies"`
}
