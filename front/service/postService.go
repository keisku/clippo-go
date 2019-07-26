package service

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/kskumgk63/clippo-go/front/proto/cachepb"
	"github.com/kskumgk63/clippo-go/front/proto/postpb"
	"github.com/kskumgk63/clippo-go/front/template"
)

// PostRegisterConfirm returns "/post/register/confirm"
func (s *FrontServer) PostRegisterConfirm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")

	req := &postpb.PostURLRequest{
		Url: url,
	}
	res, err := s.PostClient.GetPostDetail(r.Context(), req)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/top", http.StatusFound)
		return
	}
	template.Render(w, "post/postRegisterConfirmForm.tmpl", res)
}

// PostDo returns "/"
func (s *FrontServer) PostDo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")
	title := r.FormValue("title")
	description := r.FormValue("description")
	image := r.FormValue("image")
	usecase := r.FormValue("usecase")
	genre := r.FormValue("genre")

	// ディスクリプションが150文字より多かったらリダイレクト
	if utf8.RuneCountInString(description) > 150 {
		log.Printf("descriotion is too long | n = %v\n", utf8.RuneCountInString(description))
		http.Redirect(w, r, "/top", http.StatusFound)
		return
	}

	// キャッシュされているログインユーザーのIdを取得
	reqCache := &cachepb.GetIDRequest{
		Key: LOGINUSER,
	}
	resCache, _ := s.CacheClient.GetID(r.Context(), reqCache)

	// 投稿を作成するgRPCリクエスト
	reqPost := &postpb.CreatePostRequest{
		Post: &postpb.Post{
			Url:         url,
			Title:       title,
			Description: description,
			Image:       image,
			Usecase:     usecase,
			Genre:       genre,
			UserId:      resCache.Id,
		},
	}
	resPost, err := s.PostClient.CreatePost(r.Context(), reqPost)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		http.Redirect(w, r, "/post/register/init", http.StatusFound)
		return
	}
	log.Println(resPost.GetMessage())

	http.Redirect(w, r, "/top", http.StatusFound)
}

// PostDelete deletes a post
func (s *FrontServer) PostDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postID := r.FormValue("post_id")

	if postID == "xxxx" {
		template.Render(w, "top/top.tmpl", nil)
	}

	req := &postpb.DeletePostRequest{
		Id: postID,
	}
	res, _ := s.PostClient.DeletePost(r.Context(), req)
	log.Panicln(res.GetMessage())
	http.Redirect(w, r, "/top", http.StatusFound)
}

// PostSearchTitle return Posts which is match with input
func (s *FrontServer) PostSearchTitle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.FormValue("title")
	words := strings.Fields(title)
	fmt.Println(words[:])

	// キャッシュされているログインユーザーのIdを取得
	req := &cachepb.GetIDRequest{
		Key: LOGINUSER,
	}
	res, _ := s.CacheClient.GetID(r.Context(), req)
	log.Println(res.Id)
	if res.Id == "" {
		log.Println("token is empty")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	// 投稿一覧取得
	reqPost := &postpb.SearchPostsByMultiTitlesRequest{
		UserId: res.Id,
		Titles: words,
	}
	resPost, _ := s.PostClient.SearchPostsByMultiTitles(r.Context(), reqPost)

	template.Render(w, "top/top.tmpl", resPost.Posts)
}

// PostSearchUsecase return Posts which is match with input
func (s *FrontServer) PostSearchUsecase(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usecase := r.FormValue("usecase")

	// キャッシュされているログインユーザーのIdを取得
	req := &cachepb.GetIDRequest{
		Key: LOGINUSER,
	}
	res, _ := s.CacheClient.GetID(r.Context(), req)
	log.Println(res.Id)
	if res.Id == "" {
		log.Println("token is empty")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	// 投稿一覧取得
	reqPost := &postpb.SearchPostsByUsecaseRequest{
		UserId:  res.Id,
		Usecase: usecase,
	}
	resPost, _ := s.PostClient.SearchPostsByUsecase(r.Context(), reqPost)

	template.Render(w, "top/top.tmpl", resPost.Posts)
}

// PostSearchGenre return Posts which is match with input
func (s *FrontServer) PostSearchGenre(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	genre := r.FormValue("genre")

	// キャッシュされているログインユーザーのIdを取得
	req := &cachepb.GetIDRequest{
		Key: LOGINUSER,
	}
	res, _ := s.CacheClient.GetID(r.Context(), req)
	log.Println(res.Id)
	if res.Id == "" {
		log.Println("token is empty")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	// 投稿一覧取得
	reqPost := &postpb.SearchPostsByGenreRequest{
		UserId: res.Id,
		Genre:  genre,
	}
	resPost, _ := s.PostClient.SearchPostsByGenre(r.Context(), reqPost)

	template.Render(w, "top/top.tmpl", resPost.Posts)
}
