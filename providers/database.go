package providers

import (
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DatabaseConnection interface {
	GetDB() *gorm.DB
}

type databaseConnection struct {
	DB *gorm.DB
}

func NewDatabaseConnection(c *Config) DatabaseConnection {
	config := c.Get()
	user := config.GetString("db.username")
	password := config.GetString("db.password")
	database := config.GetString("db.database")
	host := config.GetString("db.host")
	port := config.GetInt("db.port")

	dsn := url.URL{
		User:     url.UserPassword(user, password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", host, port),
		Path:     database,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	conn, err := gorm.Open("postgres", dsn.String())
	if err != nil {
		panic("database connection failed")
	}
	db := conn

	if config.GetBool("db.log") {
		db.LogMode(true)
	}

	return &databaseConnection{DB: db}
}

func (d *databaseConnection) GetDB() *gorm.DB {
	return d.DB
}
