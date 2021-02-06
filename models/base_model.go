package models

import "time"

type BaseModel struct {
	ID      uint      `gorm:"primary_key" json:"id"`
	Created time.Time `json:"created" gorm:"default:now()"`
}
