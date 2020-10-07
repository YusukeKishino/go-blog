package model

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string `gorm:"size:255;not null;unique"`
	Password string
}
