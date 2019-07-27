package entity

import (
	"github.com/jinzhu/gorm"
)

// Tag tag will be used for search
type Tag struct {
	gorm.Model
	TagName, TagID string
	UserID         uint
}
