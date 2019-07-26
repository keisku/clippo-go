package entity

import (
	"github.com/jinzhu/gorm"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Post 投稿
type Post struct {
	gorm.Model
	URL, Title, Description, Image, TagID string
	UserID                                uint
}