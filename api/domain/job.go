package domain

import (
	"time"
)

type Job struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Claimable  bool
	Claimed    bool
	Progress   float64
	InputFile  string
	OutputFile string
	Failed     bool
	Error      string
	Estimation time.Duration
}
