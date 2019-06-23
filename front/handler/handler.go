package handler

import (
	"log"
	"net/http"

	"github.com/kskumgk63/Clippo-api/front/database"
	"github.com/kskumgk63/Clippo-api/front/template"
	"github.com/kskumgk63/Clippo-api/proto/post"
)

// FrontServer クライアントスタブを作成
type FrontServer struct {
	PostClient post.PostServiceClient
}

// PostRegister returns HTML
func (s *FrontServer) PostRegister(w http.ResponseWriter, r *http.Request) {
	template.Render(w, "post/postRegisterForm.html", nil)
}

// PostRegisterConfirm returns HTML
func (s *FrontServer) PostRegisterConfirm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")

	req := &post.PostURLRequest{
		Url: url,
	}
	res, err := s.PostClient.GetPostDetail(r.Context(), req)
	if err != nil {
		log.Fatalln(err)
	}
	template.Render(w, "post/postRegisterConfirmForm.html", res)
}

type Post struct{
	Url, Title, Description, Image string
}

// PostResult returns HTML
func (s *FrontServer) PostResult(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")
	title := r.FormValue("title")
	description := r.FormValue("description")
	image := r.FormValue("image")

	// MySQLと接続
	db := database.GormConnect()
	defer db.Close()
	db.Create(Post{
		Url: url,
		Title: title,
		Description: description,
		Image: image,
	})

	template.Render(w, "post/result.html", &post.PostResponse{
		Url:         url,
		Title:       title,
		Description: description,
		Image:       image,
	})
}
