package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"

	"github.com/kskumgk63/clippo-go/front/database"
	"github.com/kskumgk63/clippo-go/front/template"
	"github.com/kskumgk63/clippo-go/server_cache/cachepb"
	"github.com/kskumgk63/clippo-go/server_post/postpb"
)

// FrontServer クライアントスタブを作成
type FrontServer struct {
	PostClient  postpb.PostServiceClient
	CacheClient cachepb.CacheServiceClient
}

// Posts トップページへ構造体をマッピング
type Posts struct {
	Posts []database.Post
}

// JWT 認証用トークン
type JWT struct {
	Token string `json:"token"`
}

const (
	// TOKENCACHE 認証トークンのキー
	TOKENCACHE = "token-cache"
	// LOGINUSER ログインユーザーIdのキー
	LOGINUSER = "login-user"
)

// GenerateJWTToken JWT認証トークンを生成
func GenerateJWTToken(user database.User) (string, error) {
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
func (s *FrontServer) AuthToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// キャッシュサーバーへアクセス
		req := &cachepb.GetTokenRequest{
			Key: TOKENCACHE,
		}
		res, _ := s.CacheClient.GetToken(r.Context(), req)
		if res.Token == "" {
			log.SetFlags(log.Lshortfile)
			log.Printf("*** %v\n", "JWT Token is empty.")
			http.Redirect(w, r, "/login", http.StatusFound)
		}
		bearerToken := res.Token
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

// TopBeforeLogin returns "/"
func (s *FrontServer) TopBeforeLogin(w http.ResponseWriter, r *http.Request) {
	post := &database.Post{
		URL:         "http://localhost:8080/",
		Title:       "タイトルの見本",
		Description: "250文字以内で記事の簡単な要約を書いてください。",
		Image:       "http://designers-tips.com/wp-content/uploads/2015/03/paper-clip6.jpg",
		Usecase:     "エラー解決",
		Genre:       "プログラミング言語",
	}
	template.RenderBeforeLogin(w, "top/topBeforeLogin.tmpl", post)
}

// Test returns "/test"
func (s *FrontServer) Test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")

	req := &postpb.PostURLRequest{
		Url: url,
	}
	res, err := s.PostClient.GetPostDetail(r.Context(), req)
	if err != nil {
		log.Fatalln(err)
	}
	template.RenderBeforeLogin(w, "post/testConfirmForm.tmpl", res)
}

// TestDo returns "/"
func (s *FrontServer) TestDo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")
	title := r.FormValue("title")
	description := r.FormValue("description")
	image := r.FormValue("image")
	usecase := r.FormValue("usecase")
	genre := r.FormValue("genre")

	post := &database.Post{
		URL:         url,
		Title:       title,
		Description: description,
		Image:       image,
		Usecase:     usecase,
		Genre:       genre,
	}
	template.RenderBeforeLogin(w, "top/topBeforeLogin.tmpl", post)
}

// Login returns "/login"
func (s *FrontServer) Login(w http.ResponseWriter, r *http.Request) {
	template.RenderBeforeLogin(w, "login/loginForm.tmpl", nil)
}

// Logout returns "/"
func (s *FrontServer) Logout(w http.ResponseWriter, r *http.Request) {
	req := &cachepb.DeleteIDRequest{
		Key: LOGINUSER,
	}
	res, _ := s.CacheClient.DeleteID(r.Context(), req)
	log.Println(res)
	template.RenderBeforeLogin(w, "login/loginForm.tmpl", nil)
}

// LoginSuccess returns "/top"
func (s *FrontServer) LoginSuccess(w http.ResponseWriter, r *http.Request) {
	var user database.User

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
	reqToken := &cachepb.SetTokenRequest{
		Token: token,
		Key:   TOKENCACHE,
	}
	res, _ := s.CacheClient.SetToken(r.Context(), reqToken)
	log.Println(res.Message)
	// ログインユーザーのIdをキャッシュに格納
	reqID := &cachepb.SetIDRequest{
		Id:  fmt.Sprint(user.ID),
		Key: LOGINUSER,
	}
	resID, _ := s.CacheClient.SetID(r.Context(), reqID)
	log.Println(resID.Message)

	// 投稿一覧取得
	posts := []database.Post{}
	db.Find(&posts)
	db.Where("user_id = ?", user.ID).Find(&posts)

	template.Render(w, "top/top.tmpl", &Posts{
		Posts: posts,
	})
}

// Top returns "/top"
func (s *FrontServer) Top(w http.ResponseWriter, r *http.Request) {
	// MySQLと接続
	db := database.GormConnect()
	defer db.Close()

	// トークンをキャッシュに格納
	req := &cachepb.GetIDRequest{
		Key: LOGINUSER,
	}
	res, _ := s.CacheClient.GetID(r.Context(), req)
	log.Println(res.Id)
	if res.Id == "" {
		log.Println("id is empty")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	cached := res.Id
	posts := []database.Post{}
	err := db.Where("user_id = ?", cached).Find(&posts).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		return
	}

	template.Render(w, "top/top.tmpl", &Posts{
		Posts: posts,
	})
}

// UserRegister returns "user/register/init"
func (s *FrontServer) UserRegister(w http.ResponseWriter, r *http.Request) {
	template.RenderBeforeLogin(w, "user/userRegisterForm.tmpl", nil)
}

// UserRegisterConfirm returns "user/register/confirm"
func (s *FrontServer) UserRegisterConfirm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	// エラーハンドリング
	if email == "" {
		http.Redirect(w, r, "/user/register/init", http.StatusFound)
		return
	}
	if password == "" {
		http.Redirect(w, r, "/user/register/init", http.StatusFound)
		return
	}
	if password != confirmPassword {
		template.Render(w, "user/userRegisterForm.tmpl", database.User{
			Email:    email,
			Password: "",
		})
		return
	}

	template.Render(w, "user/userRegisterConfirmForm.tmpl", &database.User{
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

	// MySQLと接続
	db := database.GormConnect()
	defer db.Close()
	user := database.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	db.Create(&user)
	db.Model(&user).Update("CreatedAt", time.Now().Add(9*time.Hour))
	db.Model(&user).Update("UpdatedAt", time.Now().Add(9*time.Hour))

	http.Redirect(w, r, "/login", http.StatusFound)
}

// PostRegister returns "/post/register/init"
func (s *FrontServer) PostRegister(w http.ResponseWriter, r *http.Request) {
	template.Render(w, "post/postRegisterForm.tmpl", nil)
}

// PostRegisterConfirm returns "/post/register/confirm"
func (s *FrontServer) PostRegisterConfirm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")

	req := &postpb.PostURLRequest{
		Url: url,
	}
	res, err := s.PostClient.GetPostDetail(r.Context(), req)
	if err != nil {
		log.Fatalln(err)
	}
	template.Render(w, "post/postRegisterConfirmForm.tmpl", res)
}

// PostDo returns "/"
func (s *FrontServer) PostDo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")
	title := r.FormValue("title")
	description := r.FormValue("description")
	image := r.FormValue("image")
	usecase := r.FormValue("usecase")
	genre := r.FormValue("genre")

	// キャッシュされているログインユーザーのIdを取得
	req := &cachepb.GetIDRequest{
		Key: LOGINUSER,
	}
	res, _ := s.CacheClient.GetID(r.Context(), req)
	log.Println(res.Id)
	if res.Id == "" {
		log.Println("token is empty")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	id, err := strconv.ParseUint(res.Id, 10, 64)
	if err != nil {
		log.Println(err)
	}

	// MySQLと接続
	db := database.GormConnect()
	defer db.Close()
	post := database.Post{
		URL:         url,
		Title:       title,
		Description: description,
		Image:       image,
		Usecase:     usecase,
		Genre:       genre,
		UserID:      uint(id),
	}
	db.Create(&post)
	db.Model(&post).Update("CreatedAt", time.Now().Add(9*time.Hour))
	db.Model(&post).Update("UpdatedAt", time.Now().Add(9*time.Hour))

	http.Redirect(w, r, "/top", http.StatusFound)
}

// PostSearchTitle return Posts which is match with input
func (s *FrontServer) PostSearchTitle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.FormValue("title")

	// MySQLと接続
	db := database.GormConnect()
	defer db.Close()

	// キャッシュされているログインユーザーのIdを取得
	req := &cachepb.GetIDRequest{
		Key: LOGINUSER,
	}
	res, _ := s.CacheClient.GetID(r.Context(), req)
	log.Println(res.Id)
	if res.Id == "" {
		log.Println("token is empty")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	posts := []database.Post{}

	err := db.Where("user_id = ? AND title LIKE ?", res.Id, "%"+title+"%").Find(&posts).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		return
	}

	template.Render(w, "top/top.tmpl", &Posts{
		Posts: posts,
	})
}

// PostSearchUsecase return Posts which is match with input
func (s *FrontServer) PostSearchUsecase(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usecase := r.FormValue("usecase")

	// MySQLと接続
	db := database.GormConnect()
	defer db.Close()

	// キャッシュされているログインユーザーのIdを取得
	req := &cachepb.GetIDRequest{
		Key: LOGINUSER,
	}
	res, _ := s.CacheClient.GetID(r.Context(), req)
	log.Println(res.Id)
	if res.Id == "" {
		log.Println("token is empty")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	posts := []database.Post{}

	err := db.Where("user_id = ? AND usecase LIKE ?", res.Id, "%"+usecase+"%").Find(&posts).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		return
	}

	template.Render(w, "top/top.tmpl", &Posts{
		Posts: posts,
	})
}

// PostSearchGenre return Posts which is match with input
func (s *FrontServer) PostSearchGenre(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	genre := r.FormValue("genre")

	// MySQLと接続
	db := database.GormConnect()
	defer db.Close()

	// キャッシュされているログインユーザーのIdを取得
	req := &cachepb.GetIDRequest{
		Key: LOGINUSER,
	}
	res, _ := s.CacheClient.GetID(r.Context(), req)
	log.Println(res.Id)
	if res.Id == "" {
		log.Println("token is empty")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	posts := []database.Post{}

	err := db.Where("user_id = ? AND genre LIKE ?", res.Id, "%"+genre+"%").Find(&posts).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		return
	}

	template.Render(w, "top/top.tmpl", &Posts{
		Posts: posts,
	})
}
