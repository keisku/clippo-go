package service

import (
	"context"
	"log"
	"strconv"

	"github.com/kskumgk63/clippo-go/user/entity"

	"github.com/kskumgk63/clippo-go/user/repository"

	"github.com/kskumgk63/clippo-go/user/userpb"
)

// UserServer user server
type UserServer struct{}

// CreateUser create a new user
func (*UserServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	log.Println("CreateUser RUN")

	var user entity.User

	user.Email = req.GetUser().GetEmail()
	user.Password = req.GetUser().GetPassword()

	err := repository.Create(&user)
	if err != nil {
		return &userpb.CreateUserResponse{Message: "***** FAIL SAVE USER *****"}, err
	}

	return &userpb.CreateUserResponse{Message: "***** SAVE USER *****"}, nil
}

// GetUser get a user
func (*UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	log.Println("GetUser RUN")

	var user entity.User
	user.Email = req.GetEmail()

	dbUser := repository.Get(&user)
	return &userpb.GetUserResponse{
		Id: strconv.FormatUint(uint64(dbUser.ID), 10),
		User: &userpb.User{
			Email:    dbUser.Email,
			Password: dbUser.Password,
		}}, nil
}

// IsUserByEmailExisted check if user is existed
func (*UserServer) IsUserByEmailExisted(ctx context.Context, req *userpb.IsUserByEmailExistedRequest) (*userpb.IsUserByEmailExistedResponse, error) {
	log.Println("IsUserByEmailExisted RUN")
	var user entity.User
	user.Email = req.GetEmail()

	flag := repository.IsEmailExisted(&user)
	if flag {
		return &userpb.IsUserByEmailExistedResponse{
			Flag: false,
		}, nil
	} else {
		return &userpb.IsUserByEmailExistedResponse{
			Flag: true,
		}, nil
	}
}
