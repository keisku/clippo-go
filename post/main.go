package main

import (
	"log"
	"net"

	"github.com/kskumgk63/clippo-go/post/repository"

	"github.com/kskumgk63/clippo-go/post/service"

	"github.com/kskumgk63/clippo-go/post/postpb"
	"google.golang.org/grpc"
)

func main() {
	repository.CreatePostsAndTagsTable()
	log.Println("***** POST SERVER RUNNING *****")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	postpb.RegisterPostServiceServer(s, &service.PostServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
