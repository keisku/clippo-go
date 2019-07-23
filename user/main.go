package main

import (
	"github.com/kskumgk63/clippo-go/user/repository"
	"fmt"
	"log"
	"net"

	"github.com/kskumgk63/clippo-go/user/service"

	"github.com/kskumgk63/clippo-go/user/userpb"
	"google.golang.org/grpc"
)

func main() {
	log.Println(repository.GormConnect())
	fmt.Println("***** USER SERVER RUNNING *****")

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &service.UserServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
