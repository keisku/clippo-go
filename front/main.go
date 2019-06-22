package main

import (
	"log"
	"net/http"

	"github.com/kskumgk63/Clippo-api/front/handler"
	"github.com/gorilla/mux"
	"github.com/kskumgk63/Clippo-api/proto/post"
	"google.golang.org/grpc"
)

func main() {
	r := mux.NewRouter()

	postClient := post.NewPostServiceClient(getGRPCConnection())

	frontSrv := &handler.FrontServer{
		PostClient: postClient,
	}

	r.Path("/").Methods(http.MethodGet).HandlerFunc(frontSrv.PostRegister)
	r.Path("/confirm").Methods(http.MethodPost).HandlerFunc(frontSrv.PostRegisterConfirm)
	r.Path("/post").Methods(http.MethodPost).HandlerFunc(frontSrv.PostResult)

	// static フォルダの読み取り
	static := http.StripPrefix("/static", http.FileServer(http.Dir("/static")))
	r.PathPrefix("/static/").Handler(static)
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
