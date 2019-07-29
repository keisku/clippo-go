package service

import (
	"context"
	"fmt"
	"log"

	"github.com/kskumgk63/clippo-go/post/repository"

	goose "github.com/advancedlogic/GoOse"
	"github.com/kskumgk63/clippo-go/post/postpb"
)

// PostServer post server
type PostServer struct{}

// CreatePost create a new post
func (*PostServer) CreatePost(ctx context.Context, req *postpb.CreatePostRequest) (*postpb.CreatePostResponse, error) {
	fmt.Println("CreatePost RUN")
	err := repository.Create(req)
	if err != nil {
		log.Println(err)
	}
	// create a respponse message
	res := &postpb.CreatePostResponse{
		Message: "**** SAVE POST ****",
	}
	return res, nil
}

// DeletePost delete a post
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

// GetPostDetail get title, description and image from URL
func (*PostServer) GetPostDetail(ctx context.Context, req *postpb.PostURLRequest) (*postpb.PostResponse, error) {
	fmt.Printf("GetPostDetail RUN %v\n", req)

	// scraping this URL
	url := req.GetUrl()

	g := goose.New()
	article, err := g.ExtractFromURL(url)
	// if it is not available to scrape, return empty
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Println(err)
		return &postpb.PostResponse{
			Url:         "",
			Title:       "",
			Description: "",
			Image:       "",
		}, err
	}

	// create gRPC response
	return &postpb.PostResponse{
		Url:         url,
		Title:       article.Title,
		Description: article.MetaDescription,
		Image:       article.TopImage,
	}, nil
}
