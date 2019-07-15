package main

import (
	"context"
	"fmt"
	"log"
	"net"

	goose "github.com/advancedlogic/GoOse"
	"github.com/kskumgk63/clippo-go/post/postpb"
	"google.golang.org/grpc"
)

type postServer struct{}

func (*postServer) GetPostDetail(ctx context.Context, req *postpb.PostURLRequest) (*postpb.PostResponse, error) {
	fmt.Printf("GetPostDetail was invoked with %v\n", req)

	// リクエストURLのタイトルとディスクリプションをスクレイピング
	url := req.GetUrl()

	g := goose.New()
	article, _ := g.ExtractFromURL(url)

	// gRPCレスポンスの作成
	res := &postpb.PostResponse{
		Url:         url,
		Title:       article.Title,
		Description: article.MetaDescription,
		Image:       article.TopImage,
	}

	return res, nil
}

func main() {
	fmt.Println("***** SERVER RUNNING *****")

	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	postpb.RegisterPostServiceServer(s, &postServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
