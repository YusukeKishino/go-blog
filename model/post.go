package model

import "time"

type Post struct {
	ID        uint   `gorm:"primary_key"`
	Title     string `gorm:"size:255;not null"`
	Content   string `gorm:"type:text not null"`
	Status    string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
