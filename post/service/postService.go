package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kskumgk63/clippo-go/post/entity"

	"github.com/kskumgk63/clippo-go/post/repository"

	goose "github.com/advancedlogic/GoOse"
	"github.com/kskumgk63/clippo-go/post/postpb"
)

// PostServer 投稿サーバー
type PostServer struct{}

func convertPost(post *entity.Post) (pb *postpb.Post) {
	// convert uint to string
	post64 := uint64(post.ID)
	postID := strconv.FormatUint(post64, 10)

	// convert uint to string
	user64 := uint64(post.UserID)
	userID := strconv.FormatUint(user64, 10)

	// convert string to string array
	tagArray := strings.Split(post.Tag, "/")

	pb = &postpb.Post{
		Id:          postID,
		Url:         post.URL,
		Title:       post.Title,
		Description: post.Description,
		Image:       post.Image,
		Tag:         tagArray,
		UserId:      userID,
	}
	return pb
}

// CreatePost 投稿作成
func (*PostServer) CreatePost(ctx context.Context, req *postpb.CreatePostRequest) (*postpb.CreatePostResponse, error) {
	fmt.Println("CreatePost RUN")
	err := repository.Create(req)
	if err != nil {
		log.Println(err)
	}
	// 投稿作成レスポンスメッセージ
	res := &postpb.CreatePostResponse{
		Message: "**** SAVE POST ****",
	}
	return res, nil
}

// DeletePost 投稿削除
func (*PostServer) DeletePost(ctx context.Context, req *postpb.DeletePostRequest) (*postpb.DeletePostResponse, error) {
	fmt.Println("DeletePost RUN")
	err := repository.Delete(req)
	if err != nil {
		log.Println(err)
	}
	return &postpb.DeletePostResponse{Message: "**** DELETE POST ****"}, nil
}

// GetAllPostsByUserID get all posts by user_id
func (*PostServer) GetAllPostsByUserID(ctx context.Context, req *postpb.GetAllPostsByUserIDRequest) (*postpb.GetAllPostsByUserIDResponse, error) {
	fmt.Println("GetAllPostsByUserID RUN")

	// search posts grom DB
	posts := repository.GetByUserID(req)

	return &postpb.GetAllPostsByUserIDResponse{Posts: posts}, nil
}

// SearchPosts search posts with keywords of title or tags
func (*PostServer) SearchPosts(ctx context.Context, req *postpb.SearchPostsRequest) (*postpb.SearchPostsResponse, error) {
	fmt.Println("SearchPosts RUN")

	// search posts grom DB
	posts := repository.Search(req)

	return &postpb.SearchPostsResponse{Posts: posts}, nil
}

// GetPostDetail URLを基にWebスクレイピング
func (*PostServer) GetPostDetail(ctx context.Context, req *postpb.PostURLRequest) (*postpb.PostResponse, error) {
	fmt.Printf("GetPostDetail RUN %v\n", req)

	// リクエストURLのタイトルとディスクリプションをスクレイピング
	url := req.GetUrl()

	g := goose.New()
	article, err := g.ExtractFromURL(url)
	if err != nil {
		return &postpb.PostResponse{
			Url:         "",
			Title:       "",
			Description: "",
			Image:       "",
		}, err
	}

	// gRPCレスポンスの作成
	return &postpb.PostResponse{
		Url:         url,
		Title:       article.Title,
		Description: article.MetaDescription,
		Image:       article.TopImage,
	}, nil
}
