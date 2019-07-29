package service

import (
	"github.com/jinzhu/gorm"
	"github.com/kskumgk63/clippo-go/front/proto/cachepb"
	"github.com/kskumgk63/clippo-go/front/proto/postpb"
	"github.com/kskumgk63/clippo-go/front/proto/userpb"
)

const (
	// LOGINUSER key for  the user try to login
	LOGINUSER = "login-user"
	// TOKENCACHE key for jwt token
	TOKENCACHE = "token-cache"
)

// TestPost post for test
type TestPost struct {
	URL, Title, Description, Image string
	TagNames                       []string
}

func (t TestPost) makeTestPost() *TestPost {
	var tagArray []string
	tagArray = append(tagArray, "お試し")

	t.URL = "http://localhost:8080/"
	t.Title = "機能を試してください！URLをタイプして「Clip」するだけです！"
	t.Description = "150文字以内で記事の簡単なサマリーを書いてください。この記事は何を目的としているか、ジャンルは何かひと目でわかるようになっています。できるだけシンプルにサマリーを書くことをおすすめします。"
	t.Image = "http://designers-tips.com/wp-content/uploads/2015/03/paper-clip6.jpg"
	t.TagNames = tagArray

	return &t
}

// User user
type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255);unique_index;not null"`
	Password string `gorm:"type:varchar(60);not null"`
	Posts    []Post
}

// Post post
type Post struct {
	gorm.Model
	URL, Title, Description, Image, TagID string
	UserID                                uint
}

// JWT token for authorization
type JWT struct {
	Token string `json:"token"`
}

// FrontServer create a client
type FrontServer struct {
	CacheClient cachepb.CacheServiceClient
	PostClient  postpb.PostServiceClient
	UserClient  userpb.UserServiceClient
}
