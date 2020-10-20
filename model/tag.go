package model

import (
	"time"
)

type Tag struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"size:255;not null;unique"`
	Posts     []Post `gorm:"many2many:post_tags"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
