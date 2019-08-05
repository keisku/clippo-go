package repository

import (
	"log"

	"github.com/kskumgk63/clippo-go/user/entity"
)

// Create create a new user
func Create(user *entity.User) error {
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

// Get get a user
func Get(user *entity.User) *entity.User {
	// connect with DB
	db := GormConnect()
	defer db.Close()
	// get the user by email
	err := db.Where("email = ?", user.Email).Find(&user).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Println(err)
		return nil
	}
	return user
}

// IsEmailExisted check if user is existed
func IsEmailExisted(user *entity.User) bool {
	// connect with DB
	db := GormConnect()
	defer db.Close()
	return db.Where("email = ?", user.Email).Find(&user).RecordNotFound()
}
