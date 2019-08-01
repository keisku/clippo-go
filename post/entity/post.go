package entity

import (
	"github.com/jinzhu/gorm"
)

// Post for DB
type Post struct {
	gorm.Model
	URL, Title, Description, Image string `gorm:"not null"`
	UserID                         uint   `gorm:"not null"`
	Tags                           []Tag  `gorm:"many2many:post_tags"`
}
