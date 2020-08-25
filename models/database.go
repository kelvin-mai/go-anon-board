package models

import (
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&Board{},
		&Thread{},
		&Reply{},
	)
}

func Init() *gorm.DB {
	dsn := url.URL{
		User:     url.UserPassword("postgres", "postgres"),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", "localhost", 5432),
		Path:     "go_anon",
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	conn, err := gorm.Open("postgres", dsn.String())
	if err != nil {
		panic("database connection failed")
	}
	db = conn
	db.LogMode(true)
	AutoMigrate(db)
	return db
}
