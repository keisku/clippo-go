package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"github.com/kskumgk63/Clippo-api/proto/post"
	"github.com/advancedlogic/GoOse"
	"google.golang.org/grpc"
)

type postServer struct{}

func (*postServer) GetPostDetail(ctx context.Context, req *post.PostURLRequest) (*post.PostResponse, error) {
	fmt.Printf("GetPostDetail was invoked with %v\n", req)

	// リクエストURLのタイトルとディスクリプションをスクレイピング
	url := req.GetUrl()

	g := goose.New()
	article, _ := g.ExtractFromURL(url)

	// gRPCレスポンスの作成
	res := &post.PostResponse{
		Url:         url,
		Title:       article.Title,
		Description: article.MetaDescription,
		Image:       article.TopImage,
	}

	return res, nil
}

func main() {
	fmt.Println("Starting the server....")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	post.RegisterPostServiceServer(s, &postServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
