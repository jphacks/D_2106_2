package database

import (
	"os"

	"github.com/jphacks/D_2106_2/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlHandler struct {
	Conn *gorm.DB
}

// NewSqlClient initialize a new sql client.
func NewSqlClient(config *config.Config) (*SqlHandler, error) {
	if os.Getenv("MODE") == "test" {
		return &SqlHandler{nil}, nil
	}

	USER := config.DB_USER
	PASS := config.DB_PASS
	HOST := config.DB_HOST
	DBNAME := config.DB_NAME
	DB_PORT := config.DB_PORT
	dsn := USER + ":" + PASS + "@tcp(" + HOST + ":" + DB_PORT + ")/" + DBNAME
	dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		return &SqlHandler{nil}, err
	}

	return &SqlHandler{db}, nil
}
