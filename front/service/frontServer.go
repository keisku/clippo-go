package service

import (
	"github.com/kskumgk63/clippo-go/front/proto/cachepb"
	"github.com/kskumgk63/clippo-go/front/proto/postpb"
	"github.com/kskumgk63/clippo-go/front/proto/userpb"
)

const (
	// LOGINUSER ログインユーザーIdのキー
	LOGINUSER = "login-user"
	// SAMPLEURL サンプルのURL
	SAMPLEURL = "http://localhost:8080/"
	// SAMPLETITLE サンプルのタイトル
	SAMPLETITLE = "機能を試してください！URLをタイプして「Clip」するだけです！"
	// SAMPLEDESCRIPTION サンプルの詳細
	SAMPLEDESCRIPTION = "150文字以内で記事の簡単なサマリーを書いてください。この記事は何を目的としているか、ジャンルは何かひと目でわかるようになっています。できるだけシンプルにサマリーを書くことをおすすめします。"
	// SAMPLEIMAGE サンプルの画像
	SAMPLEIMAGE = "http://designers-tips.com/wp-content/uploads/2015/03/paper-clip6.jpg"
	// SAMPLETAG サンプルのタグ
	SAMPLETAG = "お試し"
	// SAMPLEID サンプルのID
	SAMPLEID = 0000

	// TOKENCACHE 認証トークンのキー
	TOKENCACHE = "token-cache"
)

// JWT 認証用トークン
type JWT struct {
	Token string `json:"token"`
}

// FrontServer クライアントスタブを作成
type FrontServer struct {
	CacheClient cachepb.CacheServiceClient
	PostClient  postpb.PostServiceClient
	UserClient  userpb.UserServiceClient
}
