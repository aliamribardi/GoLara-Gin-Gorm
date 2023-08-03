package Models

import (
	"gorm.io/gorm"
)

type User struct {
	ID 			int
	Name 		string 		`gorm:"required"`
	KelasID 	int			`gorm:"required"`
	Kelas 		Kelas 		
	gorm.Model
}