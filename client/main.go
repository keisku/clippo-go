package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	clippopb "projects/Clippo-api/proto"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	r := mux.NewRouter()

	r.Path("/get").Methods(http.MethodPost).HandlerFunc(GetArticleTitleDescriptionImg)

	svc := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatalln(svc.ListenAndServe())
}

// getGRPCConnection gRPC通信の接続
func getGRPCConnection() clippopb.ArticleServiceClient {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	return clippopb.NewArticleServiceClient(connection)
}

// GetArticleTitleDescriptionImg 送信したURLでtitle, descriptionをスクレイピングする
func GetArticleTitleDescriptionImg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")

	req := &clippopb.ArticleURLRequest{
		Url: url,
	}

	c := getGRPCConnection()
	res, err := c.GetArticleTitleDescriptionImg(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	bytes, _ := json.Marshal(&res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
