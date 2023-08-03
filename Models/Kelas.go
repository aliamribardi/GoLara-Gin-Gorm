package Models

import (
	"gorm.io/gorm"
)

type Kelas struct {
	ID 		int
	Kelas 	string		`gorm:"required"`
	gorm.Model
}