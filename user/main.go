package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/kskumgk63/clippo-go/database"

	"github.com/kskumgk63/clippo-go/user/userpb"
	"google.golang.org/grpc"
)

type userServer struct{}

func (*userServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	email := req.GetUser().GetEmail()
	password := req.GetUser().GetPassword()

	// MySQLと接続
	db := database.GormConnect()
	defer db.Close()
	user := database.User{
		Email:    email,
		Password: password,
	}
	db.Create(&user)
	db.Model(&user).Update("CreatedAt", time.Now().Add(9*time.Hour))
	db.Model(&user).Update("UpdatedAt", time.Now().Add(9*time.Hour))

	return &userpb.CreateUserResponse{Message: "***** SAVE USER *****"}, nil
}
func (*userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	var dbUser database.User
	email := req.GetEmail()
	// MySQLからユーザーの取得
	db := database.GormConnect()
	defer db.Close()
	db.Find(&dbUser, "email=?", email)
	return &userpb.GetUserResponse{
		Id: strconv.FormatUint(uint64(dbUser.ID), 10),
		User: &userpb.User{
			Email:    dbUser.Email,
			Password: dbUser.Password,
		}}, nil
}

func (*userServer) IsUserByEmailExisted(ctx context.Context, req *userpb.IsUserByEmailExistedRequest) (*userpb.IsUserByEmailExistedResponse, error) {
	var dbUser database.User
	email := req.GetEmail()
	// MySQLからユーザーの取得
	db := database.GormConnect()
	defer db.Close()
	err := db.Find(&dbUser, "email=?", email).Error
	if err != nil {
		log.Println(err)
		return &userpb.IsUserByEmailExistedResponse{
			Flag: false,
		}, nil
	}
	return &userpb.IsUserByEmailExistedResponse{
		Flag: true,
	}, nil
}

func main() {
	fmt.Println("***** USER SERVER RUNNING *****")

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &userServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
