package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kskumgk63/clippo-go/cache/service"

	"github.com/kskumgk63/clippo-go/cache/cachepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("***** CACHE SERVER RUNNING *****")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	cachepb.RegisterCacheServiceServer(s, &service.CacheServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
