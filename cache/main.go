package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/kskumgk63/clippo-go/cache/cachepb"
	gocache "github.com/pmylund/go-cache"
	"google.golang.org/grpc"
)

type cacheServer struct{}

var cache = gocache.New(1*time.Hour, 2*time.Hour)

func (*cacheServer) SetToken(ctx context.Context, req *cachepb.SetTokenRequest) (*cachepb.SetTokenResponse, error) {
	log.Println("================================================")
	log.Println("SetToken is invoked")
	log.Printf("Key = %v Token = %v\n", req.Key, req.Token)
	token := req.GetToken()
	key := req.GetKey()
	cache.Set(key, token, gocache.DefaultExpiration)
	res := &cachepb.SetTokenResponse{
		Message: "Set Token",
	}
	return res, nil
}

func (*cacheServer) GetToken(ctx context.Context, req *cachepb.GetTokenRequest) (*cachepb.GetTokenResponse, error) {
	log.Println("================================================")
	log.Println("GetToken is invoked")
	log.Printf("Key = %v\n", req.Key)
	key := req.GetKey()
	// キャッシュを取り出す
	cached, found := cache.Get(key)
	log.Println(cached)
	// 見つからなければリダイレクト
	if !found {
		res := &cachepb.GetTokenResponse{
			Token: "",
		}
		return res, nil
	}
	res := &cachepb.GetTokenResponse{
		Token: cached.(string),
	}
	return res, nil
}

func (*cacheServer) DeleteToken(ctx context.Context, req *cachepb.DeleteTokenRequest) (*cachepb.DeleteTokenResponse, error) {
	log.Println("================================================")
	log.Println("DeleteToken is invoked")
	log.Printf("Key = %v\n", req.Key)
	key := req.GetKey()
	cache.Delete(key)
	res := &cachepb.DeleteTokenResponse{
		Message: "Delete cache",
	}
	return res, nil
}

func (*cacheServer) SetID(ctx context.Context, req *cachepb.SetIDRequest) (*cachepb.SetIDResponse, error) {
	log.Println("================================================")
	log.Println("SetID is invoked")
	log.Printf("Key = %v ID = %v\n", req.Key, req.Id)
	id := req.GetId()
	key := req.GetKey()
	cache.Set(key, id, gocache.DefaultExpiration)
	res := &cachepb.SetIDResponse{
		Message: "Set ID",
	}
	return res, nil
}

func (*cacheServer) GetID(ctx context.Context, req *cachepb.GetIDRequest) (*cachepb.GetIDResponse, error) {
	log.Println("================================================")
	log.Println("GetID is invoked")
	log.Printf("Key = %v\n", req.Key)
	key := req.GetKey()
	// キャッシュを取り出す
	cached, found := cache.Get(key)
	// 見つからなければリダイレクト
	if !found {
		res := &cachepb.GetIDResponse{
			Id: "",
		}
		return res, nil
	}
	res := &cachepb.GetIDResponse{
		Id: cached.(string),
	}
	return res, nil
}

func (*cacheServer) DeleteID(ctx context.Context, req *cachepb.DeleteIDRequest) (*cachepb.DeleteIDResponse, error) {
	log.Println("================================================")
	log.Println("DelteID is invoked")
	log.Printf("Key = %v\n", req.Key)
	key := req.GetKey()
	cache.Delete(key)
	res := &cachepb.DeleteIDResponse{
		Message: "Delete cache",
	}
	return res, nil
}

func main() {
	fmt.Println("***** SERVER RUNNING *****")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	cachepb.RegisterCacheServiceServer(s, &cacheServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
