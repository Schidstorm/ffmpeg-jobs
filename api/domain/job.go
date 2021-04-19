package domain

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Started    bool
	Progress   float64
	InputFile  string
	OutputFile string
}
