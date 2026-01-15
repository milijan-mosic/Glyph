package models

import (
	"time"
)

type Article struct {
	ID          string    `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Content     string    `gorm:"type:text" json:"content"`
	Author      string    `gorm:"not null" json:"author"`
	Published   bool      `gorm:"default:false" json:"published"`
	CreatedAt   time.Time `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null;default:now()" json:"modified_at"`

	Comments []Comment `gorm:"foreignKey:ArticleID;constraint:OnDelete:CASCADE"`
}
