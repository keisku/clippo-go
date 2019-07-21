package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kskumgk63/clippo-go/cache/cachepb"
	"github.com/kskumgk63/clippo-go/post/postpb"
	"github.com/kskumgk63/clippo-go/user/userpb"

	"github.com/gorilla/mux"
	// "github.com/kskumgk63/clippo-go/repository"
	"github.com/kskumgk63/clippo-go/front/service"
	"google.golang.org/grpc"
)

// getGRPCConnection gRPCと接続
func getGRPCConnection(port string) *grpc.ClientConn {
	connection, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	return connection
}

func main() {
	// テーブル作成
	// repository.CreateTable()

	fmt.Println("***** SERVER RUNNING *****")

	r := mux.NewRouter()

	cacheClient := cachepb.NewCacheServiceClient(getGRPCConnection(":50051"))
	postClient := postpb.NewPostServiceClient(getGRPCConnection(":50052"))
	userClient := userpb.NewUserServiceClient(getGRPCConnection(":50053"))

	frontSrv := &service.FrontServer{
		PostClient:  postClient,
		CacheClient: cacheClient,
		UserClient:  userClient,
	}

	r.Path("/").Methods(http.MethodGet).HandlerFunc(frontSrv.TopBeforeLogin)
	r.Path("/test").Methods(http.MethodPost).HandlerFunc(frontSrv.Test)
	r.Path("/test/do").Methods(http.MethodPost).HandlerFunc(frontSrv.TestDo)
	r.Path("/test/delete").Methods(http.MethodPost).HandlerFunc(frontSrv.TopBeforeLogin)
	r.Path("/login").Methods(http.MethodGet).HandlerFunc(frontSrv.Login)
	r.Path("/logout").Methods(http.MethodGet).HandlerFunc(frontSrv.Logout)
	r.Path("/login/success").Methods(http.MethodPost).HandlerFunc(frontSrv.LoginSuccess)
	r.Path("/top").Methods(http.MethodGet).HandlerFunc(frontSrv.AuthToken(frontSrv.Top))

	r.Path("/user/register/init").Methods(http.MethodGet).HandlerFunc(frontSrv.UserRegister)
	r.Path("/user/register/confirm").Methods(http.MethodPost).HandlerFunc(frontSrv.UserRegisterConfirm)
	r.Path("/user/register/do").Methods(http.MethodPost).HandlerFunc(frontSrv.UserRegisterDo)

	r.Path("/post/register/confirm").Methods(http.MethodPost).HandlerFunc(frontSrv.AuthToken(frontSrv.PostRegisterConfirm))
	r.Path("/post/register/do").Methods(http.MethodPost).HandlerFunc(frontSrv.AuthToken(frontSrv.PostDo))
	r.Path("/post/delete").Methods(http.MethodPost).HandlerFunc(frontSrv.AuthToken(frontSrv.PostDelete))
	r.Path("/post/search/title").Methods(http.MethodPost).HandlerFunc(frontSrv.AuthToken(frontSrv.PostSearchTitle))
	r.Path("/post/search/usecase").Methods(http.MethodPost).HandlerFunc(frontSrv.AuthToken(frontSrv.PostSearchUsecase))
	r.Path("/post/search/genre").Methods(http.MethodPost).HandlerFunc(frontSrv.AuthToken(frontSrv.PostSearchGenre))

	/*
		static フォルダの読み取り
		.clippo-go/front/
		上記パスで実行されることを前提とする。
	*/
	r.PathPrefix("/static/css/").Handler(http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css"))))
	r.PathPrefix("/static/js/").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("static/js"))))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("failed to exit serve: ", err)
	}
}
