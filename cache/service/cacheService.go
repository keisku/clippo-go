package service

import (
	"context"
	"time"

	gocache "github.com/pmylund/go-cache"

	"github.com/kskumgk63/clippo-go/cache/cachepb"
)

// CacheServer キャッシュサーバー
type CacheServer struct{}

var cache = gocache.New(1*time.Hour, 2*time.Hour)

// SetToken トークン格納
func (*CacheServer) SetToken(ctx context.Context, req *cachepb.SetTokenRequest) (*cachepb.SetTokenResponse, error) {
	token := req.GetToken()
	key := req.GetKey()
	cache.Set(key, token, gocache.DefaultExpiration)
	res := &cachepb.SetTokenResponse{
		Message: "Set Token",
	}
	return res, nil
}

// GetToken トークンの取得
func (*CacheServer) GetToken(ctx context.Context, req *cachepb.GetTokenRequest) (*cachepb.GetTokenResponse, error) {
	key := req.GetKey()
	// キャッシュを取り出す
	cached, found := cache.Get(key)
	// 見つからなければリダイレクト
	if !found {
		return nil, nil
	}
	res := &cachepb.GetTokenResponse{
		Token: cached.(string),
	}
	return res, nil
}

// DeleteToken トークンの削除
func (*CacheServer) DeleteToken(ctx context.Context, req *cachepb.DeleteTokenRequest) (*cachepb.DeleteTokenResponse, error) {
	key := req.GetKey()
	cache.Delete(key)
	res := &cachepb.DeleteTokenResponse{
		Message: "Delete cache",
	}
	return res, nil
}

// SetID ユーザーIDの格納
func (*CacheServer) SetID(ctx context.Context, req *cachepb.SetIDRequest) (*cachepb.SetIDResponse, error) {
	id := req.GetId()
	key := req.GetKey()
	cache.Set(key, id, gocache.DefaultExpiration)
	res := &cachepb.SetIDResponse{
		Message: "SetID success",
	}
	return res, nil
}

// GetID ユーザーIDの取得
func (*CacheServer) GetID(ctx context.Context, req *cachepb.GetIDRequest) (*cachepb.GetIDResponse, error) {
	key := req.GetKey()
	// キャッシュを取り出す
	cached, found := cache.Get(key)
	// 見つからなければリダイレクト
	if !found {
		res := &cachepb.GetIDResponse{
			Id: "",
		}
		return res, nil
	}
	res := &cachepb.GetIDResponse{
		Id: cached.(string),
	}
	return res, nil
}

// DeleteID ユーザーIDの削除
func (*CacheServer) DeleteID(ctx context.Context, req *cachepb.DeleteIDRequest) (*cachepb.DeleteIDResponse, error) {
	key := req.GetKey()
	cache.Delete(key)
	res := &cachepb.DeleteIDResponse{
		Message: "Delete cache",
	}
	return res, nil
}
