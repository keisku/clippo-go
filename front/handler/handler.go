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

// TopContetnt トップページへ構造体をマッピング
type TopContetnt struct {
	Posts []Post
}

// Post DB格納用の構造体
type Post struct {
	URL, Title, Description, Image string
}

// Top returns HTML("/")
func (s *FrontServer) Top(w http.ResponseWriter, r *http.Request) {
	// MySQLと接続
	db := database.GormConnect()
	defer db.Close()

	posts := []Post{}
	db.Find(&posts)

	template.Render(w, "top/top.html", &TopContetnt{
		Posts: posts,
	})
}

// PostRegister returns HTML("/post/register/init")
func (s *FrontServer) PostRegister(w http.ResponseWriter, r *http.Request) {
	template.Render(w, "post/postRegisterForm.html", nil)
}

// PostRegisterConfirm returns HTML("/post/register/confirm")
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

// PostResult returns HTML("/")
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
		URL:         url,
		Title:       title,
		Description: description,
		Image:       image,
	})

	http.Redirect(w, r, "/", http.StatusFound)
}
