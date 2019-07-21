package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kskumgk63/clippo-go/front/entity"
	"github.com/kskumgk63/clippo-go/front/template"
	"github.com/kskumgk63/clippo-go/front/proto/userpb"
	"golang.org/x/crypto/bcrypt"
)

// UserRegister returns "user/register/init"
func (s *FrontServer) UserRegister(w http.ResponseWriter, r *http.Request) {
	template.RenderBeforeLogin(w, "user/userRegisterForm.tmpl", nil)
}

// UserRegisterConfirm returns "user/register/confirm"
func (s *FrontServer) UserRegisterConfirm(w http.ResponseWriter, r *http.Request) {
	// フォームから取得した値
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	// 入力された値が空ならリダイレクト
	if email == "" {
		http.Redirect(w, r, "/user/register/init", http.StatusFound)
		return
	}
	if password == "" {
		http.Redirect(w, r, "/user/register/init", http.StatusFound)
		return
	}
	if password != confirmPassword {
		template.Render(w, "user/userRegisterForm.tmpl", entity.User{
			Email:    email,
			Password: "",
		})
		return
	}

	// 入力されたEメールがすでに存在しているかエラーハンドリング
	req := &userpb.IsUserByEmailExistedRequest{
		Email: email,
	}
	res, _ := s.UserClient.IsUserByEmailExisted(r.Context(), req)
	if res.Flag {
		http.Redirect(w, r, "/user/register/init", http.StatusFound)
		return
	}

	template.Render(w, "user/userRegisterConfirmForm.tmpl", &entity.User{
		Email:    email,
		Password: password,
	})
}

// UserRegisterDo returns "/login"
func (s *FrontServer) UserRegisterDo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	// エラーハンドリング
	if email == "" {
		http.Redirect(w, r, "/user/register/init", http.StatusFound)
		return
	}
	if password == "" {
		http.Redirect(w, r, "/user/register/init", http.StatusFound)
		return
	}

	// ハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		return
	}

	req := &userpb.CreateUserRequest{
		User: &userpb.User{
			Email:    email,
			Password: string(hashedPassword),
		},
	}
	res, _ := s.UserClient.CreateUser(r.Context(), req)
	log.Println(res.Message)

	http.Redirect(w, r, "/login", http.StatusFound)
}
