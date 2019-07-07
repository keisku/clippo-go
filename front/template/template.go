package template

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// Render テンプレートエンジン
func Render(w http.ResponseWriter, name string, content interface{}) {

	/*
		ワーキングディレクトリをGET
		.clippo-go/front/
		上記パスで実行されることを前提とする。
	*/
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	t, err := template.ParseFiles(
		wd+"/template/layout.tmpl", wd+"/template/header.tmpl", wd+"/template/"+name)
	if err != nil {
		log.Fatalln(err)
	}
	if err := t.ExecuteTemplate(w, "layout", content); err != nil {
		panic(err)
	}
}
