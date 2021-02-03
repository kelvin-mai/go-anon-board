package providers

import (
	"fmt"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	var enableLogging logger.Interface
	if config.GetBool("db.log") {
		enableLogging = logger.Default
	}

	dsn := url.URL{
		User:     url.UserPassword(user, password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", host, port),
		Path:     database,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		Logger: enableLogging,
	})
	if err != nil {
		panic("database connection failed")
	}

	return &databaseConnection{DB: db}
}

func (d *databaseConnection) GetDB() *gorm.DB {
	return d.DB
}
