package service

import (
	"github.com/kskumgk63/clippo-go/cache/cachepb"
	"github.com/kskumgk63/clippo-go/post/postpb"
	"github.com/kskumgk63/clippo-go/user/userpb"
)

// FrontServer クライアントスタブを作成
type FrontServer struct {
	CacheClient cachepb.CacheServiceClient
	PostClient  postpb.PostServiceClient
	UserClient  userpb.UserServiceClient
}
