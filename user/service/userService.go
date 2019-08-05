package service

import (
	"context"
	"log"
	"strconv"

	"github.com/kskumgk63/clippo-go/user/repository"

	"github.com/kskumgk63/clippo-go/user/userpb"
)

// UserServer ユーザーサーバー
type UserServer struct{}

// CreateUser ユーザー作成
func (*UserServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	log.Println("CreateUser RUN")

	email := req.GetUser().GetEmail()
	password := req.GetUser().GetPassword()

	err := repository.Create(email, password)
	if err != nil {
		return &userpb.CreateUserResponse{Message: "***** FAIL SAVE USER *****"}, err
	}

	return &userpb.CreateUserResponse{Message: "***** SAVE USER *****"}, nil
}

// GetUser ユーザー取得
func (*UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	log.Println("GetUser RUN")
	email := req.GetEmail()

	user, err := repository.Get(email)
	if err != nil {
		return &userpb.GetUserResponse{
			Id: strconv.FormatUint(uint64(user.ID), 10),
			User: &userpb.User{
				Email:    "",
				Password: "",
			}}, err
	}
	return &userpb.GetUserResponse{
		Id: strconv.FormatUint(uint64(user.ID), 10),
		User: &userpb.User{
			Email:    user.Email,
			Password: user.Password,
		}}, nil
}

// IsUserByEmailExisted メールアドレスに紐づくユーザーが存在するか判定
func (*UserServer) IsUserByEmailExisted(ctx context.Context, req *userpb.IsUserByEmailExistedRequest) (*userpb.IsUserByEmailExistedResponse, error) {
	log.Println("IsUserByEmailExisted RUN")
	email := req.GetEmail()
	flag := repository.IsEmailExisted(email)
	if flag {
		return &userpb.IsUserByEmailExistedResponse{
			Flag: true,
		}, nil
	}
	return &userpb.IsUserByEmailExistedResponse{
		Flag: false,
	}, nil
}
