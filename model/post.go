package model

import (
	"database/sql"
	"time"
)

type PostStatus string

const (
	Draft     PostStatus = "draft"
	Published PostStatus = "published"
)

type Post struct {
	ID          uint       `gorm:"primary_key"`
	Title       string     `gorm:"size:255;not null"`
	Content     string     `gorm:"type:text not null"`
	Status      PostStatus `gorm:"size:255;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt sql.NullTime
	Tags        []Tag `gorm:"many2many:post_tags"`
}

func (p *Post) IsPublished() bool {
	return p.Status == Published
}
