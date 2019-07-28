package entity

import (
	"github.com/jinzhu/gorm"
)

// Post for DB
type Post struct {
	gorm.Model
	URL, Title, Description, Image, Tag string
	UserID                              uint
}
