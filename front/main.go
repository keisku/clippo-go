package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/kskumgk63/clippo-go/front/database"
	"github.com/kskumgk63/clippo-go/front/handler"
	"github.com/kskumgk63/clippo-go/server_cache/cachepb"
	"github.com/kskumgk63/clippo-go/server_post/postpb"
	"google.golang.org/grpc"
)

// getGRPCConnection gRPCと接続
func getGRPCConnection(port string) *grpc.ClientConn {
	connection, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	return connection
}

func main() {
	// テーブル作成
	// db := database.GormConnect()
	// database.CreateTable(db)
	// defer db.Close()

	fmt.Println("***** SERVER RUNNING *****")

	r := mux.NewRouter()

	cacheClient := cachepb.NewCacheServiceClient(getGRPCConnection("50051"))
	postClient := postpb.NewPostServiceClient(getGRPCConnection("50052"))

	frontSrv := &handler.FrontServer{
		PostClient:  postClient,
		CacheClient: cacheClient,
	}

	r.Path("/").Methods(http.MethodGet).HandlerFunc(frontSrv.TopBeforeLogin)
	r.Path("/test").Methods(http.MethodPost).HandlerFunc(frontSrv.Test)
	r.Path("/test/do").Methods(http.MethodPost).HandlerFunc(frontSrv.TestDo)
	r.Path("/login").Methods(http.MethodGet).HandlerFunc(frontSrv.Login)
	r.Path("/logout").Methods(http.MethodGet).HandlerFunc(frontSrv.Logout)
	r.Path("/login/success").Methods(http.MethodPost).HandlerFunc(frontSrv.LoginSuccess)
	r.Path("/top").Methods(http.MethodGet).HandlerFunc(frontSrv.AuthToken(frontSrv.Top))

	r.Path("/user/register/init").Methods(http.MethodGet).HandlerFunc(frontSrv.UserRegister)
	r.Path("/user/register/confirm").Methods(http.MethodPost).HandlerFunc(frontSrv.UserRegisterConfirm)
	r.Path("/user/register/do").Methods(http.MethodPost).HandlerFunc(frontSrv.UserRegisterDo)

	r.Path("/post/register/init").Methods(http.MethodGet).HandlerFunc(frontSrv.AuthToken(frontSrv.PostRegister))
	r.Path("/post/register/confirm").Methods(http.MethodPost).HandlerFunc(frontSrv.AuthToken(frontSrv.PostRegisterConfirm))
	r.Path("/post/register/do").Methods(http.MethodPost).HandlerFunc(frontSrv.AuthToken(frontSrv.PostDo))
	r.Path("/post/search/title").Methods(http.MethodPost).HandlerFunc(frontSrv.AuthToken(frontSrv.PostSearchTitle))
	r.Path("/post/search/usecase").Methods(http.MethodPost).HandlerFunc(frontSrv.AuthToken(frontSrv.PostSearchUsecase))
	r.Path("/post/search/genre").Methods(http.MethodPost).HandlerFunc(frontSrv.AuthToken(frontSrv.PostSearchGenre))

	/*
		static フォルダの読み取り
		.clippo-go/front/
		上記パスで実行されることを前提とする。
	*/
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("failed to exit serve: ", err)
	}
}
