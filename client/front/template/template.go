package template

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// Render テンプレートエンジン
func Render(w http.ResponseWriter, name string, content interface{}) {

	// ワーキングディレクトリをGET
	// 実行するディレクトリのパスを取得
	// Clippo-apiで実行すると、/Users/umegakikeisuke/go/src/projects/Clippo-api/ となる
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// client以下を
	wd += "/client/front/"
	t, err := template.ParseFiles(
		wd+"template/layout.html", wd+"template/header.html", wd+"template/"+name)
	if err != nil {
		log.Fatalln(err)
	}
	if err := t.ExecuteTemplate(w, "layout", content); err != nil {
		panic(err)
	}
}
