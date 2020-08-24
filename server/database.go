package server

import (
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kelvin-mai/go-anon-board/models"
)

var DB *gorm.DB

func DbConnect() *gorm.DB {
	dsn := url.URL{
		User:     url.UserPassword("postgres", "postgres"),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", "localhost", 5432),
		Path:     "go_anon",
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open("postgres", dsn.String())
	if err != nil {
		panic("database connection failed")
	}
	db.LogMode(true)
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

func AutoMigrate() {
	db := GetDB()
	db.AutoMigrate(&models.Thread{})
	db.AutoMigrate(&models.Reply{})
}
