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
	tag := r.FormValue("tag_name")

	// chech if description is more than 150
	if utf8.RuneCountInString(description) > 150 {
		log.Printf("descriotion is too long | n = %v\n", utf8.RuneCountInString(description))
		http.Redirect(w, r, "/top", http.StatusFound)
		return
	}

	// get userID from cache
	reqCache := &cachepb.GetIDRequest{
		Key: LOGINUSER,
	}
	resCache, _ := s.CacheClient.GetID(r.Context(), reqCache)

	// create gRPC request
	reqPost := &postpb.CreatePostRequest{
		Post: &postpb.Post{
			Url:         url,
			Title:       title,
			Description: description,
			Image:       image,
			Tag:         strings.Fields(tag),
			UserId:      resCache.Id,
		},
	}
	// send the request
	_, err := s.PostClient.CreatePost(r.Context(), reqPost)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		http.Redirect(w, r, "/post/register/init", http.StatusFound)
		return
	}

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

// PostSearch return Posts which is match with input
func (s *FrontServer) PostSearch(w http.ResponseWriter, r *http.Request) {
	var keywords []string
	r.ParseForm()
	// check how to search posts by title or tags
	how := r.FormValue("HowSearch")
	if how == "title" {
		title := r.FormValue("title")
		keywords = strings.Fields(title)
	}
	if how == "tag" {
		tags := r.FormValue("tag_name")
		keywords = strings.Fields(tags)
	}
	// get user_id from cache
	req := &cachepb.GetIDRequest{
		Key: LOGINUSER,
	}
	res, _ := s.CacheClient.GetID(r.Context(), req)
	if res.Id == "" {
		log.Println("token is empty")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	// make a request
	reqPost := &postpb.SearchPostsRequest{
		UserId:   res.Id,
		How:      how,
		Keywords: keywords,
	}
	// send a request
	resPost, _ := s.PostClient.SearchPosts(r.Context(), reqPost)

	template.Render(w, "top/top.tmpl", resPost.Posts)
}
