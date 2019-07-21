package repository

import (
	"log"

	"github.com/kskumgk63/clippo-go/front/entity"
)

// IsUserByEmailExisted 入力されたメールアドレスで登録されたユーザーが存在するか判定
// true : 存在している
// false : 存在していない
func IsUserByEmailExisted(user entity.User, email string) bool {
	db := GormConnect()
	err := db.Find(&user, "email=?", email).Error
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}
