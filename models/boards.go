package models

type Board struct {
	ID      uint
	Title   string `gorm:"unique"`
	Threads []Thread
}

func FindOrCreateBoard(title string, b *Board) {
	if db.Where("title = ?", title).First(&b).RecordNotFound() {
		b = &Board{Title: title}
		db.Create(&b)
	}
}
