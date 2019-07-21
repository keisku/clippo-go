package service

import (
	"github.com/kskumgk63/clippo-go/front/proto/cachepb"
	"github.com/kskumgk63/clippo-go/front/proto/postpb"
	"github.com/kskumgk63/clippo-go/front/proto/userpb"
)

// FrontServer クライアントスタブを作成
type FrontServer struct {
	CacheClient cachepb.CacheServiceClient
	PostClient  postpb.PostServiceClient
	UserClient  userpb.UserServiceClient
}
