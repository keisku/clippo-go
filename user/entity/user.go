package entity

import (
	"github.com/jinzhu/gorm"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
User ユーザー
	HasMany Posts, UserID(外部キー)
*/
type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255);unique_index;not null"`
	Password string `gorm:"type:varchar(60);not null"`
	Posts    []Post
}
