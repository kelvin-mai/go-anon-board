package models

type Board struct {
	BaseModel
	Title   string   `gorm:"unique" json:"title"`
	Threads []Thread `json:"threads"`
}
