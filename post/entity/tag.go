package entity

import (
	"github.com/jinzhu/gorm"
)

// Tag 投稿
type Tag struct {
	gorm.Model
	TagName, TagID string
	UserID         uint
}
