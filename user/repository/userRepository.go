package repository

import (
	"log"

	"github.com/kskumgk63/clippo-go/user/entity"
)

// Create create a new user
func Create(email, password string) error {
	var user entity.User
	// connect with DB
	db := GormConnect()
	defer db.Close()

	// Create
	err := db.Create(&user).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatalln(err)
		return err
	}

	return nil
}

// Get ユーザーの取得
func Get(email string) (entity.User, error) {
	var user entity.User
	// connect with DB
	db := GormConnect()
	defer db.Close()
	// get the user by email
	err := db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Println(err)
		return user, err
	}
	return user, nil
}

// IsEmailExisted メールアドレスに紐づくユーザーは登録されているか精査
func IsEmailExisted(email string) bool {
	var user entity.User
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
