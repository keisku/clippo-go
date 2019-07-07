package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kskumgk63/clippo-go/front/handler"
	"github.com/kskumgk63/clippo-go/proto/post"
	"google.golang.org/grpc"
)

// getGRPCConnection gRPCと接続
func getGRPCConnection() *grpc.ClientConn {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
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

	postClient := post.NewPostServiceClient(getGRPCConnection())

	frontSrv := &handler.FrontServer{
		PostClient: postClient,
	}

	r.Path("/login").Methods(http.MethodGet).HandlerFunc(frontSrv.Login)
	r.Path("/login/success").Methods(http.MethodPost).HandlerFunc(frontSrv.LoginSuccess)
	r.Path("/top").Methods(http.MethodGet).HandlerFunc(handler.AuthToken(frontSrv.Top))

	r.Path("/user/register/init").Methods(http.MethodGet).HandlerFunc(frontSrv.UserRegister)
	r.Path("/user/register/confirm").Methods(http.MethodPost).HandlerFunc(frontSrv.UserRegisterConfirm)
	r.Path("/user/register/do").Methods(http.MethodPost).HandlerFunc(frontSrv.UserRegisterDo)

	r.Path("/post/register/init").Methods(http.MethodGet).HandlerFunc(handler.AuthToken(frontSrv.PostRegister))
	r.Path("/post/register/confirm").Methods(http.MethodPost).HandlerFunc(handler.AuthToken(frontSrv.PostRegisterConfirm))
	r.Path("/post/register/do").Methods(http.MethodPost).HandlerFunc(handler.AuthToken(frontSrv.PostDo))
	r.Path("/post/search").Methods(http.MethodPost).HandlerFunc(handler.AuthToken(frontSrv.PostSearch))

	/*
		static フォルダの読み取り
		.clippo-go/front/
		上記パスで実行されることを前提とする。
	*/
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	svc := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
	}
	log.Fatalln(svc.ListenAndServe())
}
