package database

import (
	"errors"
	"github.com/schidstorm/ffmpeg-jobs/api/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) Connect(dialectorName, dsn string) error {
	var dialector gorm.Dialector

	switch dialectorName {
	case "sqlite":
		dialector = sqlite.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	default:
		return errors.New("dialectorName must be sqlite or postgres")
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
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
