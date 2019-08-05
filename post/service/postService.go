package service

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/kskumgk63/clippo-go/post/entity"
	"github.com/kskumgk63/clippo-go/post/repository"

	goose "github.com/advancedlogic/GoOse"
	"github.com/kskumgk63/clippo-go/post/postpb"
)

// PostServer post server
type PostServer struct{}

func convertStringToUint(s string) uint {
	u, _ := strconv.ParseUint(s, 10, 64)
	return uint(u)
}

func makeSamplePost() *postpb.Post {
	var tagArray []string
	tagArray = append(tagArray, "sample")
	return &postpb.Post{
		Id:          "xxxx",
		Url:         "http://localhost:8080/",
		Title:       "まだ投稿されていないようなので、記事をクリップしてみてください",
		Description: "150文字以内で記事の簡単なサマリーを書いてください。この記事は何を目的としているか、ジャンルは何かひと目でわかるようになっています。できるだけシンプルにサマリーを書くことをおすすめします。",
		Image:       "http://designers-tips.com/wp-content/uploads/2015/03/paper-clip6.jpg",
		Tag:         tagArray,
		UserId:      "0000",
	}
}

func convertPost(post *entity.Post) *postpb.Post {
	// convert uint to string
	post64 := uint64(post.ID)
	postID := strconv.FormatUint(post64, 10)

	// convert uint to string
	user64 := uint64(post.UserID)
	userID := strconv.FormatUint(user64, 10)

	var tagArray []string
	tags := post.Tags

	// convert struct to []string for the view
	for _, tag := range tags {
		tagArray = append(tagArray, tag.TagName)
	}

	return &postpb.Post{
		Id:          postID,
		Url:         post.URL,
		Title:       post.Title,
		Description: post.Description,
		Image:       post.Image,
		Tag:         tagArray,
		UserId:      userID,
	}
}

// CreatePost create a new post
func (*PostServer) CreatePost(ctx context.Context, req *postpb.CreatePostRequest) (*postpb.CreatePostResponse, error) {
	fmt.Println("CreatePost RUN")

	var post entity.Post

	id := convertStringToUint(req.GetPost().GetUserId())
	tagNames := req.GetPost().GetTag()

	// array to string
	var tags []entity.Tag
	for _, tagName := range tagNames {
		tag, err := repository.GetTag(tagName)
		if err != nil {
			tag = repository.CreateTag(tagName)
		}
		tags = append(tags, *tag)
	}
	post = entity.Post{
		URL:         req.GetPost().GetUrl(),
		Title:       req.GetPost().GetTitle(),
		Description: req.GetPost().GetDescription(),
		Image:       req.GetPost().GetImage(),
		Tags:        tags,
		UserID:      id,
	}

	err := repository.Create(&post)
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

	var post entity.Post

	post.ID = convertStringToUint(req.GetId())
	err := repository.Delete(&post)
	if err != nil {
		log.Println(err)
	}
	return &postpb.DeletePostResponse{Message: "**** DELETE POST ****"}, nil
}

// GetAllPostsByUserID get all posts by user_id
func (*PostServer) GetAllPostsByUserID(ctx context.Context, req *postpb.GetAllPostsByUserIDRequest) (*postpb.GetAllPostsByUserIDResponse, error) {
	fmt.Println("GetAllPostsByUserID RUN")

	var post entity.Post
	post.UserID = convertStringToUint(req.GetUserId())

	// search posts grom DB
	posts := repository.GetByUserID(&post)

	// if not found any posts in DB, return sample
	var pbs []*postpb.Post
	if posts == nil {
		pb := makeSamplePost()
		pbs = append(pbs, pb)
	} else {
		for _, post := range posts {
			pbs = append(pbs, convertPost(post))
		}
	}

	return &postpb.GetAllPostsByUserIDResponse{Posts: pbs}, nil
}

// SearchPosts search posts with keywords of title or tags
func (*PostServer) SearchPosts(ctx context.Context, req *postpb.SearchPostsRequest) (*postpb.SearchPostsResponse, error) {
	fmt.Println("SearchPosts RUN")

	var pbs []*postpb.Post

	// search posts grom DB
	posts := repository.Search(req.GetHow(), req.GetUserId(), req.GetKeywords())
	if posts == nil {
		pb := makeSamplePost()
		pbs = append(pbs, pb)
	} else {
		for _, post := range posts {
			pbs = append(pbs, convertPost(post))
		}
	}

	return &postpb.SearchPostsResponse{Posts: pbs}, nil
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
		return nil, err
	}

	// create gRPC response
	return &postpb.PostResponse{
		Url:         url,
		Title:       article.Title,
		Description: article.MetaDescription,
		Image:       article.TopImage,
	}, nil
}
