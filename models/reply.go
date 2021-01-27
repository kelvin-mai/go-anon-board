package models

type Reply struct {
	BaseModel
	Text           string `json:"text" binding:"required"`
	Reported       bool   `gorm:"default:false" json:"reported"`
	DeletePassword string `json:"delete_password" gorm:"not null" binding:"required"`
	ThreadID       uint   `json:"-"`
}
