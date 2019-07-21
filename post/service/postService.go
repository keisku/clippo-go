package service

import (
	"context"
	"fmt"
	"log"
	"strconv"

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
	// SAMPLEUSECASE サンプルユースケース
	SAMPLEUSECASE = "エラー解決"
	// SAMPLEGENRE サンプルジャンル
	SAMPLEGENRE = "プログラミング言語"
	// SAMPLEUSERID サンプルユーザーID
	SAMPLEUSERID = "0000"
)

func mappingPost(id, url, title, description, image, usecase, genre, userID string) (p *postpb.Post) {
	p = &postpb.Post{
		Id:          id,
		Url:         url,
		Title:       title,
		Description: description,
		Image:       image,
		Usecase:     usecase,
		Genre:       genre,
		UserId:      userID,
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

// GetAllPostsByUserID ユーザーIDに紐づく投稿を全取得
func (*PostServer) GetAllPostsByUserID(ctx context.Context, req *postpb.GetAllPostsByUserIDRequest) (*postpb.GetAllPostsByUserIDResponse, error) {
	fmt.Println("GetAllPostsByUserID RUN")
	var posts []*postpb.Post

	dbPosts := repository.GetByUserID(req)

	// DBから何も見つからなければサンプルを返す
	if len(dbPosts) == 0 {
		p := mappingPost(
			SAMPLEPOSTID,
			SAMPLEURL,
			SAMPLETITLE,
			SAMPLEDESCRIPTION,
			SAMPLEIMAGE,
			SAMPLEUSECASE,
			SAMPLEGENRE,
			SAMPLEUSERID)
		posts = append(posts, p)
		return &postpb.GetAllPostsByUserIDResponse{Posts: posts}, nil
	}
	// DBから見つかれば、レスポンス用にマッピング
	for i := 0; i < len(dbPosts); i++ {
		u64 := uint64(dbPosts[i].UserID)
		uid := strconv.FormatUint(u64, 10)

		p64 := uint64(dbPosts[i].ID)
		pid := strconv.FormatUint(p64, 10)

		posts = append(posts, mappingPost(pid, dbPosts[i].URL, dbPosts[i].Title, dbPosts[i].Description, dbPosts[i].Image, dbPosts[i].Usecase, dbPosts[i].Genre, uid))
	}
	return &postpb.GetAllPostsByUserIDResponse{Posts: posts}, nil
}

// SearchPostsByTitle 投稿のタイトル検索
func (*PostServer) SearchPostsByTitle(ctx context.Context, req *postpb.SearchPostsByTitleRequest) (*postpb.SearchPostsByTitleResponse, error) {
	fmt.Println("SearchPostsByTitle RUN")
	var posts []*postpb.Post

	dbPosts := repository.SearchByTitle(req)
	// DBから何も見つからなければサンプルを返す
	if len(dbPosts) == 0 {
		p := mappingPost(
			SAMPLEPOSTID,
			SAMPLEURL,
			SAMPLETITLE,
			SAMPLEDESCRIPTION,
			SAMPLEIMAGE,
			SAMPLEUSECASE,
			SAMPLEGENRE,
			SAMPLEUSERID)
		posts = append(posts, p)
		return &postpb.SearchPostsByTitleResponse{Posts: posts}, nil
	}
	// DBから見つかれば、レスポンス用にマッピング
	for i := 0; i < len(dbPosts); i++ {
		u64 := uint64(dbPosts[i].UserID)
		uid := strconv.FormatUint(u64, 10)

		p64 := uint64(dbPosts[i].ID)
		pid := strconv.FormatUint(p64, 10)

		posts = append(posts, mappingPost(pid, dbPosts[i].URL, dbPosts[i].Title, dbPosts[i].Description, dbPosts[i].Image, dbPosts[i].Usecase, dbPosts[i].Genre, uid))
	}
	return &postpb.SearchPostsByTitleResponse{Posts: posts}, nil
}

// SearchPostsByUsecase 投稿のユースケース検索
func (*PostServer) SearchPostsByUsecase(ctx context.Context, req *postpb.SearchPostsByUsecaseRequest) (*postpb.SearchPostsByUsecaseResponse, error) {
	fmt.Println("SearchPostsByUsecase RUN")
	var posts []*postpb.Post

	dbPosts := repository.SearchByUsecase(req)
	// DBから何も見つからなければサンプルを返す
	if len(dbPosts) == 0 {
		p := mappingPost(
			SAMPLEPOSTID,
			SAMPLEURL,
			SAMPLETITLE,
			SAMPLEDESCRIPTION,
			SAMPLEIMAGE,
			SAMPLEUSECASE,
			SAMPLEGENRE,
			SAMPLEUSERID)
		posts = append(posts, p)
		return &postpb.SearchPostsByUsecaseResponse{Posts: posts}, nil
	}
	// DBから見つかれば、レスポンス用にマッピング
	for i := 0; i < len(dbPosts); i++ {
		u64 := uint64(dbPosts[i].UserID)
		uid := strconv.FormatUint(u64, 10)

		p64 := uint64(dbPosts[i].ID)
		pid := strconv.FormatUint(p64, 10)

		posts = append(posts, mappingPost(pid, dbPosts[i].URL, dbPosts[i].Title, dbPosts[i].Description, dbPosts[i].Image, dbPosts[i].Usecase, dbPosts[i].Genre, uid))
	}
	return &postpb.SearchPostsByUsecaseResponse{Posts: posts}, nil
}

// SearchPostsByGenre 投稿のジャンル検索
func (*PostServer) SearchPostsByGenre(ctx context.Context, req *postpb.SearchPostsByGenreRequest) (*postpb.SearchPostsByGenreResponse, error) {
	fmt.Println("SearchPostsByGenre RUN")
	var posts []*postpb.Post

	dbPosts := repository.SearchByGenre(req)
	// DBから何も見つからなければサンプルを返す
	if len(dbPosts) == 0 {
		p := mappingPost(
			SAMPLEPOSTID,
			SAMPLEURL,
			SAMPLETITLE,
			SAMPLEDESCRIPTION,
			SAMPLEIMAGE,
			SAMPLEUSECASE,
			SAMPLEGENRE,
			SAMPLEUSERID)
		posts = append(posts, p)
		return &postpb.SearchPostsByGenreResponse{Posts: posts}, nil
	}
	// DBから見つかれば、レスポンス用にマッピング
	for i := 0; i < len(dbPosts); i++ {
		u64 := uint64(dbPosts[i].UserID)
		uid := strconv.FormatUint(u64, 10)

		p64 := uint64(dbPosts[i].ID)
		pid := strconv.FormatUint(p64, 10)

		posts = append(posts, mappingPost(pid, dbPosts[i].URL, dbPosts[i].Title, dbPosts[i].Description, dbPosts[i].Image, dbPosts[i].Usecase, dbPosts[i].Genre, uid))
	}
	return &postpb.SearchPostsByGenreResponse{Posts: posts}, nil
}

// GetPostDetail URLを基にWebスクレイピング
func (*PostServer) GetPostDetail(ctx context.Context, req *postpb.PostURLRequest) (*postpb.PostResponse, error) {
	fmt.Printf("GetPostDetail RUN %v\n", req)

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
