package domain

import (
	"time"
)

type Job struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Started    bool
	Progress   float64
	InputFile  string
	OutputFile string
}
