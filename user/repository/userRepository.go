package repository

import (
	"log"

	"github.com/kskumgk63/clippo-go/user/entity"
	"github.com/kskumgk63/clippo-go/user/userpb"
)

// Create ユーザーの作成
func Create(req *userpb.CreateUserRequest) error {
	var user entity.User
	email := req.GetUser().GetEmail()
	password := req.GetUser().GetPassword()

	// MySQLと接続
	db := GormConnect()
	defer db.Close()
	user.Email = email
	user.Password = password
	err := db.Create(&user).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatalln(err)
		return err
	}

	return nil
}

// Get ユーザーの取得
func Get(req *userpb.GetUserRequest) (entity.User, error) {
	var user entity.User
	email := req.GetEmail()
	// MySQLからユーザーの取得
	db := GormConnect()
	defer db.Close()
	err := db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Println(err)
		return user, err
	}
	return user, nil
}

// IsEmailExisted メールアドレスに紐づくユーザーは登録されているか精査
func IsEmailExisted(req *userpb.IsUserByEmailExistedRequest) bool {
	var user entity.User
	email := req.GetEmail()
	// MySQLからユーザーの取得
	db := GormConnect()
	defer db.Close()
	err := db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Println(err)
		return false
	}
	return true
}
