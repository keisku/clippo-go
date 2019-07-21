package entity

import (
	"github.com/jinzhu/gorm"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Post 投稿モデル
type Post struct {
	gorm.Model
	URL, Title, Description, Image, Usecase, Genre string
	UserID                                         uint
}
