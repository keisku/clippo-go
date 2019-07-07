package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	gocache "github.com/pmylund/go-cache"

	"golang.org/x/crypto/bcrypt"

	"github.com/kskumgk63/Clippo-api/front/database"
	"github.com/kskumgk63/Clippo-api/front/template"
	"github.com/kskumgk63/Clippo-api/proto/post"
)

// FrontServer クライアントスタブを作成
type FrontServer struct {
	PostClient post.PostServiceClient
}

// Posts トップページへ構造体をマッピング
type Posts struct {
	Posts []Post
}

// User DB格納用の構造体
type User struct {
	Email, Password string
}

// Post DB格納用の構造体
type Post struct {
	URL, Title, Description, Image string
}

// JWT 認証用トークン
type JWT struct {
	Token string `json:"token"`
}

// TOKENCACHE 認証トークンのキャッシュ格納時のキー
const TOKENCACHE = "token-cache"

var cache = gocache.New(1*time.Hour, 2*time.Hour)

// GenerateJWTToken JWT認証トークンを生成
func GenerateJWTToken(user User) (string, error) {
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
	}

	return tokenString, nil
}

// AuthToken 認証トークンが含まれているかチェックするミドルウェア
func AuthToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cached, found := cache.Get(TOKENCACHE)
		if !found {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
		bearerToken := cached.(string)
		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			log.SetFlags(log.Lshortfile)
			log.Printf("*** %v\n", fmt.Sprint(err))
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			log.SetFlags(log.Lshortfile)
			log.Printf("*** %v\n", fmt.Sprint(err))
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}

// Login returns "/login"
func (s *FrontServer) Login(w http.ResponseWriter, r *http.Request) {
	template.Render(w, "login/loginForm.html", nil)
}

// LoginSuccess returns "/top"
func (s *FrontServer) LoginSuccess(w http.ResponseWriter, r *http.Request) {
	var user User

	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	// MySQLからユーザーの取得
	db := database.GormConnect()
	defer db.Close()
	err := db.Find(&user, "email=?", email).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// DBからのパスワード
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	token, err := GenerateJWTToken(user)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		return
	}
	// トークンをキャッシュに格納
	cache.Set(TOKENCACHE, token, gocache.DefaultExpiration)

	// 投稿一覧取得
	posts := []Post{}
	db.Find(&posts)

	template.Render(w, "top/top.html", &Posts{
		Posts: posts,
	})
}

// Top returns "/top"
func (s *FrontServer) Top(w http.ResponseWriter, r *http.Request) {
	// MySQLと接続
	db := database.GormConnect()
	defer db.Close()

	posts := []Post{}
	err := db.Find(&posts).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		return
	}

	template.Render(w, "top/top.html", &Posts{
		Posts: posts,
	})
}

// UserRegister returns "user/register/init"
func (s *FrontServer) UserRegister(w http.ResponseWriter, r *http.Request) {
	template.Render(w, "user/userRegisterForm.html", nil)
}

// UserRegisterConfirm returns "user/register/confirm"
func (s *FrontServer) UserRegisterConfirm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	if password != confirmPassword {
		template.Render(w, "user/userRegisterForm.html", User{
			Email:    email,
			Password: "",
		})
		return
	}

	template.Render(w, "user/userRegisterConfirmForm.html", User{
		Email:    email,
		Password: password,
	})
}

// UserRegisterDo returns "/login"
func (s *FrontServer) UserRegisterDo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	// ハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		return
	}

	// DBからのパスワード
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		http.Redirect(w, r, "/user/register/init", http.StatusFound)
		return
	}

	// MySQLと接続
	db := database.GormConnect()
	defer db.Close()
	db.Create(User{
		Email:    email,
		Password: string(hashedPassword),
	})

	http.Redirect(w, r, "/login", http.StatusFound)
}

// PostRegister returns "/post/register/init"
func (s *FrontServer) PostRegister(w http.ResponseWriter, r *http.Request) {
	template.Render(w, "post/postRegisterForm.html", nil)
}

// PostRegisterConfirm returns "/post/register/confirm"
func (s *FrontServer) PostRegisterConfirm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")

	req := &post.PostURLRequest{
		Url: url,
	}
	res, err := s.PostClient.GetPostDetail(r.Context(), req)
	if err != nil {
		log.Fatalln(err)
	}
	template.Render(w, "post/postRegisterConfirmForm.html", res)
}

// PostDo returns "/"
func (s *FrontServer) PostDo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")
	title := r.FormValue("title")
	description := r.FormValue("description")
	image := r.FormValue("image")

	// MySQLと接続
	db := database.GormConnect()
	defer db.Close()
	db.Create(Post{
		URL:         url,
		Title:       title,
		Description: description,
		Image:       image,
	})

	http.Redirect(w, r, "/top", http.StatusFound)
}
