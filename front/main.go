package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kskumgk63/Clippo-api/front/handler"
	"github.com/kskumgk63/Clippo-api/proto/post"
	"google.golang.org/grpc"
)

func main() {

	// gorilla
	r := mux.NewRouter()

	postClient := post.NewPostServiceClient(getGRPCConnection())

	frontSrv := &handler.FrontServer{
		PostClient: postClient,
	}

	r.Path("/").Methods(http.MethodGet).HandlerFunc(frontSrv.Top)
	r.Path("/post/register/init").Methods(http.MethodGet).HandlerFunc(frontSrv.PostRegister)
	r.Path("/post/register/confirm").Methods(http.MethodPost).HandlerFunc(frontSrv.PostRegisterConfirm)
	r.Path("/post/register/do").Methods(http.MethodPost).HandlerFunc(frontSrv.PostResult)

	/*
		static フォルダの読み取り
		.Clippo-api/front/
		上記パスで実行されることを前提とする。
	*/
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	svc := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
	}
	log.Fatalln(svc.ListenAndServe())
}

// getGRPCConnection gRPCとの通信の接続
func getGRPCConnection() *grpc.ClientConn {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	return connection
}