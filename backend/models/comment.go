package models

import (
	"time"
)

type Comment struct {
	ID         string    `gorm:"type:uuid;primaryKey" json:"id"`
	ArticleID  string    `gorm:"type:uuid;not null;index" json:"article_id"`
	AuthorName string    `gorm:"not null" json:"author_name"`
	Content    string    `gorm:"type:text;not null" json:"content"`
	Approved   bool      `gorm:"default:false" json:"approved"`
	CreatedAt  time.Time `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt  time.Time `gorm:"not null;default:now()" json:"modified_at"`

	Article Article `gorm:"foreignKey:ArticleID"`
}
