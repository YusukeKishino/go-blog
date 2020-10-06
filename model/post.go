package model

import "time"

type PostStatus string

const (
	Draft       PostStatus = "draft"
	Published   PostStatus = "published"
	Unpublished PostStatus = "unpublished"
)

type Post struct {
	ID        uint       `gorm:"primary_key"`
	Title     string     `gorm:"size:255;not null"`
	Content   string     `gorm:"type:text not null"`
	Status    PostStatus `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
