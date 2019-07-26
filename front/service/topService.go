package service

import (
	"log"
	"net/http"
	"unicode/utf8"

	"github.com/kskumgk63/clippo-go/front/entity"
	"github.com/kskumgk63/clippo-go/front/proto/cachepb"
	"github.com/kskumgk63/clippo-go/front/proto/postpb"
	"github.com/kskumgk63/clippo-go/front/template"
)

// Top returns "/top"
func (s *FrontServer) Top(w http.ResponseWriter, r *http.Request) {
	// トークンをキャッシュに格納
	req := &cachepb.GetIDRequest{
		Key: LOGINUSER,
	}
	res, _ := s.CacheClient.GetID(r.Context(), req)
	log.Println(res.Id)
	if res.Id == "" {
		log.Println("id is empty")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	// 投稿一覧取得
	reqPost := &postpb.GetAllPostsByUserIDRequest{
		UserId: res.Id,
	}
	resPost, err := s.PostClient.GetAllPostsByUserID(r.Context(), reqPost)
	if err != nil {
		log.Println(err)
	}

	template.Render(w, "top/top.tmpl", resPost.Posts)
}

// TopBeforeLogin returns "/"
func (s *FrontServer) TopBeforeLogin(w http.ResponseWriter, r *http.Request) {
	post := &entity.Post{
		URL:         SAMPLEURL,
		Title:       SAMPLETITLE,
		Description: SAMPLEDESCRIPTION,
		Image:       SAMPLEIMAGE,
		TagID:       SAMPLETAG,
		UserID:      SAMPLEID,
	}
	template.RenderBeforeLogin(w, "top/topBeforeLogin.tmpl", post)
}

// Test returns "/test"
func (s *FrontServer) Test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")

	req := &postpb.PostURLRequest{
		Url: url,
	}
	res, err := s.PostClient.GetPostDetail(r.Context(), req)
	if err != nil {
		log.Fatalln(err)
	}
	template.RenderBeforeLogin(w, "post/testConfirmForm.tmpl", res)
}

// TestDo returns "/"
func (s *FrontServer) TestDo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")
	title := r.FormValue("title")
	description := r.FormValue("description")
	image := r.FormValue("image")
	tagName := r.FormValue("tag_name")

	// ディスクリプションが150文字より多かったらリダイレクト
	if utf8.RuneCountInString(description) > 150 {
		http.Redirect(w, r, "/test", http.StatusFound)
	}

	post := &entity.Post{
		URL:         url,
		Title:       title,
		Description: description,
		Image:       image,
		TagID:       tagName,
	}
	template.RenderBeforeLogin(w, "top/topBeforeLogin.tmpl", post)
}
