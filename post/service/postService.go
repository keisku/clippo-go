package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kskumgk63/clippo-go/post/repository"

	goose "github.com/advancedlogic/GoOse"
	"github.com/kskumgk63/clippo-go/post/postpb"
)

// PostServer 投稿サーバー
type PostServer struct{}

const (
	// SAMPLEPOSTID サンプル投稿ID
	SAMPLEPOSTID = "xxxx"
	// SAMPLEURL サンプルURL
	SAMPLEURL = "http://localhost:8080/"
	// SAMPLETITLE サンプルタイトル
	SAMPLETITLE = "まだ投稿されていないようなので、記事をクリップしてみてください"
	// SAMPLEDESCRIPTION サンプルディスクリプション
	SAMPLEDESCRIPTION = "150文字以内で記事の簡単なサマリーを書いてください。この記事は何を目的としているか、ジャンルは何かひと目でわかるようになっています。できるだけシンプルにサマリーを書くことをおすすめします。"
	// SAMPLEIMAGE サンプルイメージ
	SAMPLEIMAGE = "http://designers-tips.com/wp-content/uploads/2015/03/paper-clip6.jpg"
	// SAMPLEUSERID サンプルユーザーID
	SAMPLEUSERID = "0000"
	// SAMPLETAG サンプルタグ
	SAMPLETAG = "サンプル"
)

func mappingPost(id, url, title, description, image, userID string, tag []string) (p *postpb.Post) {
	p = &postpb.Post{
		Id:          id,
		Url:         url,
		Title:       title,
		Description: description,
		Image:       image,
		UserId:      userID,
		Tag:         tag,
	}
	return p
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
	var posts []*postpb.Post

	// search posts grom DB
	dbPosts := repository.GetByUserID(req)

	// if posts from DB are not found, return SAMPLE
	if len(dbPosts) == 0 {
		var tagArray []string
		tagArray = append(tagArray, SAMPLETAG)
		p := mappingPost(
			SAMPLEPOSTID,
			SAMPLEURL,
			SAMPLETITLE,
			SAMPLEDESCRIPTION,
			SAMPLEIMAGE,
			SAMPLEUSERID,
			tagArray,
		)
		posts = append(posts, p)
		return &postpb.GetAllPostsByUserIDResponse{Posts: posts}, nil
	}
	// if posts from DB are EXISTED, return posts after convert
	for i := 0; i < len(dbPosts); i++ {
		// convert uint to string
		post64 := uint64(dbPosts[i].ID)
		postID := strconv.FormatUint(post64, 10)

		// convert uint to string
		user64 := uint64(dbPosts[i].UserID)
		userID := strconv.FormatUint(user64, 10)

		// convert string to string array
		tagArray := strings.Split(dbPosts[i].Tag, "/")

		// make posts array
		posts = append(posts, mappingPost(postID, dbPosts[i].URL, dbPosts[i].Title, dbPosts[i].Description, dbPosts[i].Image, userID, tagArray))
	}
	return &postpb.GetAllPostsByUserIDResponse{Posts: posts}, nil
}

// SearchPostsByTitle 投稿のタイトル検索
func (*PostServer) SearchPostsByTitle(ctx context.Context, req *postpb.SearchPostsByTitleRequest) (*postpb.SearchPostsByTitleResponse, error) {
	fmt.Println("SearchPostsByTitle RUN")
	var posts []*postpb.Post

	// search posts grom DB
	dbPosts := repository.SearchByTitle(req)

	// if posts from DB are not found, return SAMPLE
	if len(dbPosts) == 0 {
		var tagArray []string
		tagArray = append(tagArray, SAMPLETAG)
		p := mappingPost(
			SAMPLEPOSTID,
			SAMPLEURL,
			SAMPLETITLE,
			SAMPLEDESCRIPTION,
			SAMPLEIMAGE,
			SAMPLEUSERID,
			tagArray,
		)
		posts = append(posts, p)
		return &postpb.SearchPostsByTitleResponse{Posts: posts}, nil
	}
	// if posts from DB are EXISTED, return posts after convert
	for i := 0; i < len(dbPosts); i++ {
		// convert uint to string
		post64 := uint64(dbPosts[i].ID)
		postID := strconv.FormatUint(post64, 10)

		// convert uint to string
		user64 := uint64(dbPosts[i].UserID)
		userID := strconv.FormatUint(user64, 10)

		// convert string to string array
		tagArray := strings.Split(dbPosts[i].Tag, "/")

		// make posts array
		posts = append(posts, mappingPost(postID, dbPosts[i].URL, dbPosts[i].Title, dbPosts[i].Description, dbPosts[i].Image, userID, tagArray))
	}
	return &postpb.SearchPostsByTitleResponse{Posts: posts}, nil
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
