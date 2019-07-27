package entity

import (
	"github.com/jinzhu/gorm"
)

// Post 投稿
type Post struct {
	gorm.Model
	URL, Title, Description, Image, TagID string
	UserID                                uint
}
