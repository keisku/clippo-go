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
	err := repository.Create(req)
	if err != nil {
		return &userpb.CreateUserResponse{Message: "***** FAIL SAVE USER *****"}, err
	}

	return &userpb.CreateUserResponse{Message: "***** SAVE USER *****"}, nil
}

// GetUser ユーザー取得
func (*UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	log.Println("GetUser RUN")
	dbUser, err := repository.Get(req)
	if err != nil {
		return &userpb.GetUserResponse{
			Id: strconv.FormatUint(uint64(dbUser.ID), 10),
			User: &userpb.User{
				Email:    "",
				Password: "",
			}}, err
	}
	return &userpb.GetUserResponse{
		Id: strconv.FormatUint(uint64(dbUser.ID), 10),
		User: &userpb.User{
			Email:    dbUser.Email,
			Password: dbUser.Password,
		}}, nil
}

// IsUserByEmailExisted メールアドレスに紐づくユーザーが存在するか判定
func (*UserServer) IsUserByEmailExisted(ctx context.Context, req *userpb.IsUserByEmailExistedRequest) (*userpb.IsUserByEmailExistedResponse, error) {
	log.Println("IsUserByEmailExisted RUN")
	flag := repository.IsEmailExisted(req)
	if flag {
		return &userpb.IsUserByEmailExistedResponse{
			Flag: true,
		}, nil
	}
	return &userpb.IsUserByEmailExistedResponse{
		Flag: false,
	}, nil
}
