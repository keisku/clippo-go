package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/kskumgk63/clippo-go/database"

	goose "github.com/advancedlogic/GoOse"
	"github.com/kskumgk63/clippo-go/post/postpb"
	"google.golang.org/grpc"
)

type postServer struct{}

const (
	// SAMPLEPOSTID サンプル投稿ID
	SAMPLEPOSTID      = "xxxx"
	// SAMPLEURL サンプルURL
	SAMPLEURL         = "http://localhost:8080/"
	// SAMPLETITLE サンプルタイトル
	SAMPLETITLE       = "まだ投稿されていないようなので、記事をクリップしてみてください"
	// SAMPLEDESCRIPTION サンプルディスクリプション
	SAMPLEDESCRIPTION = "150文字以内で記事の簡単なサマリーを書いてください。この記事は何を目的としているか、ジャンルは何かひと目でわかるようになっています。できるだけシンプルにサマリーを書くことをおすすめします。"
	// SAMPLEIMAGE サンプルイメージ
	SAMPLEIMAGE       = "http://designers-tips.com/wp-content/uploads/2015/03/paper-clip6.jpg"
	// SAMPLEUSECASE サンプルユースケース
	SAMPLEUSECASE     = "エラー解決"
	// SAMPLEGENRE サンプルジャンル
	SAMPLEGENRE       = "プログラミング言語"
	// SAMPLEUSERID サンプルユーザーID
	SAMPLEUSERID      = "0000"
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
func (*postServer) CreatePost(ctx context.Context, req *postpb.CreatePostRequest) (*postpb.CreatePostResponse, error) {
	// 文字列で受け取るのでuintへ変換
	resID := req.GetPost().GetUserId()
	id64, _ := strconv.ParseUint(resID, 10, 64)
	id := uint(id64)

	// データベースと接続
	db := database.GormConnect()
	defer db.Close()
	post := database.Post{
		URL:         req.GetPost().GetUrl(),
		Title:       req.GetPost().GetTitle(),
		Description: req.GetPost().GetDescription(),
		Image:       req.GetPost().GetImage(),
		Usecase:     req.GetPost().GetUsecase(),
		Genre:       req.GetPost().GetGenre(),
		UserID:      id,
	}
	err := db.Create(&post).Error
	if err != nil {
		return nil, err
	}
	db.Model(&post).Update("CreatedAt", time.Now().Add(9*time.Hour))
	db.Model(&post).Update("UpdatedAt", time.Now().Add(9*time.Hour))
	// 投稿作成レスポンスメッセージ
	res := &postpb.CreatePostResponse{
		Message: fmt.Sprintf("|| ** SAVE ** || << Post >> Title = < %v > URL = < %v >\n", post.Title, post.URL),
	}
	return res, nil
}
func (*postServer) DeletePost(ctx context.Context, req *postpb.DeletePostRequest) (*postpb.DeletePostResponse, error) {
	var post database.Post
	id64, _ := strconv.ParseUint(req.GetId(), 10, 64)
	post.ID = uint(id64)
	// データベースと接続
	db := database.GormConnect()
	defer db.Close()
	db.Delete(&post)
	return &postpb.DeletePostResponse{Message: "**** DELETE POST ****"}, nil
}
func (*postServer) GetAllPostsByUserID(ctx context.Context, req *postpb.GetAllPostsByUserIDRequest) (*postpb.GetAllPostsByUserIDResponse, error) {
	fmt.Println("GetAllPostsByUserID RUN")
	var posts []*postpb.Post

	// このユーザーIDを基にDB検索
	id := req.GetUserId()

	// データベースと接続
	db := database.GormConnect()
	defer db.Close()
	// 投稿一覧取得
	dbPosts := []database.Post{}
	db.Find(&dbPosts)
	db.Where("user_id = ?", id).Find(&dbPosts)
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

func (*postServer) SearchPostsByTitle(ctx context.Context, req *postpb.SearchPostsByTitleRequest) (*postpb.SearchPostsByTitleResponse, error) {
	fmt.Println("SearchPostsByTitle RUN")
	var posts []*postpb.Post

	// このユーザーIDを基にDB検索
	id := req.GetUserId()
	title := req.GetTitle()

	// データベースと接続
	db := database.GormConnect()
	defer db.Close()
	// 投稿一覧取得
	dbPosts := []database.Post{}
	db.Where("user_id = ? AND title LIKE ?", id, "%"+title+"%").Find(&dbPosts)
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
func (*postServer) SearchPostsByUsecase(ctx context.Context, req *postpb.SearchPostsByUsecaseRequest) (*postpb.SearchPostsByUsecaseResponse, error) {
	fmt.Println("SearchPostsByUsecase RUN")
	var posts []*postpb.Post

	// このユーザーIDを基にDB検索
	id := req.GetUserId()
	usecase := req.GetUsecase()

	// データベースと接続
	db := database.GormConnect()
	defer db.Close()
	// 投稿一覧取得
	dbPosts := []database.Post{}
	db.Where("user_id = ? AND usecase LIKE ?", id, "%"+usecase+"%").Find(&dbPosts)
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
func (*postServer) SearchPostsByGenre(ctx context.Context, req *postpb.SearchPostsByGenreRequest) (*postpb.SearchPostsByGenreResponse, error) {
	fmt.Println("SearchPostsByGenre RUN")
	var posts []*postpb.Post

	// このユーザーIDを基にDB検索
	id := req.GetUserId()
	genre := req.GetGenre()

	// データベースと接続
	db := database.GormConnect()
	defer db.Close()
	// 投稿一覧取得
	dbPosts := []database.Post{}
	db.Where("user_id = ? AND genre LIKE ?", id, "%"+genre+"%").Find(&dbPosts)
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

func (*postServer) GetPostDetail(ctx context.Context, req *postpb.PostURLRequest) (*postpb.PostResponse, error) {
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

func main() {
	fmt.Println("***** POST SERVER RUNNING *****")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	postpb.RegisterPostServiceServer(s, &postServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
