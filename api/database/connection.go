package database

import (
	"github.com/schidstorm/ffmpeg-jobs/api/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) Connect() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(domain.Index()...)
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

func (c *Connection) DB() *gorm.DB {
	return c.db
}
