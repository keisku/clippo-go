package main

import (
	"context"
	"fmt"
	"log"
	"net"
	clippopb "projects/Clippo-api/proto"

	goose "github.com/advancedlogic/GoOse"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) GetArticleTitleDescriptionImg(ctx context.Context, req *clippopb.ArticleURLRequest) (*clippopb.ArticleTitleDescriptionImgResponse, error) {
	fmt.Printf("GetArticleTitleDescriprion was invoked with %v\n", req)

	// リクエストURLのタイトルとディスクリプションをスクレイピング
	url := req.GetUrl()

	g := goose.New()
	article, _ := g.ExtractFromURL(url)

	// gRPCレスポンスの作成
	res := &clippopb.ArticleTitleDescriptionImgResponse{
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
	clippopb.RegisterArticleServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
