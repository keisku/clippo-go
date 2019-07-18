package handler

import (
	"github.com/kskumgk63/clippo-go/front/template"
	"fmt"
	"log"
	"net/http"

	"github.com/kskumgk63/clippo-go/user/userpb"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"

	"github.com/kskumgk63/clippo-go/cache/cachepb"
	"github.com/kskumgk63/clippo-go/database"
	"github.com/kskumgk63/clippo-go/post/postpb"
)

// FrontServer クライアントスタブを作成
type FrontServer struct {
	PostClient  postpb.PostServiceClient
	CacheClient cachepb.CacheServiceClient
	UserClient  userpb.UserServiceClient
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

	SAMPLEURL         = "http://loc alhost:8080/"
	SAMPLETITLE       = "まだ投稿されていないようなので、記事をクリップしてみてください"
	SAMPLEDESCRIPTION = "250文字以内で記事の簡単なサマリーを書いてください。この記事は何を目的としているか、ジャンルは何かひと目でわかるようになっています。できるだけシンプルにサマリーを書くことをおすすめします。"
	SAMPLEIMAGE       = "http://designers-tips.com/wp-content/uploads/2015/03/paper-clip6.jpg"
	SAMPLEUSECASE     = "エラー解決"
	SAMPLEGENRE       = "プログラミング言語"
	SAMPLEID          = "0000"
)

// GenerateJWTToken JWT認証トークンを生成
func GenerateJWTToken(user *userpb.User) (string, error) {
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
		URL:         SAMPLETITLE,
		Title:       SAMPLETITLE,
		Description: SAMPLEDESCRIPTION,
		Image:       SAMPLEIMAGE,
		Usecase:     SAMPLEUSECASE,
		Genre:       SAMPLEGENRE,
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
	// フォームから取得する値
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	// gRPC通信
	reqUser := &userpb.GetUserRequest{
		Email: email,
	}
	resUser, _ := s.UserClient.GetUser(r.Context(), reqUser)
	user := resUser.User

	// パスワードの正誤判断
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// 有効なユーザーに対して認証トークン生成
	token, err := GenerateJWTToken(user)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		return
	}
	// 認証トークンをキャッシュに格納
	reqToken := &cachepb.SetTokenRequest{
		Token: token,
		Key:   TOKENCACHE,
	}
	res, _ := s.CacheClient.SetToken(r.Context(), reqToken)
	log.Println(res.Message)

	// ログインユーザーのIdをキャッシュに格納
	reqID := &cachepb.SetIDRequest{
		Id:  resUser.Id,
		Key: LOGINUSER,
	}
	resID, _ := s.CacheClient.SetID(r.Context(), reqID)
	log.Println(resID.Message)

	// 投稿一覧取得
	reqPost := &postpb.GetAllPostsByUserIDRequest{
		UserId: reqID.Id,
	}
	resPost, err := s.PostClient.GetAllPostsByUserID(r.Context(), reqPost)
	if err != nil {
		log.Println(err)
	}

	template.Render(w, "top/top.tmpl", resPost.Posts)
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
	// 投稿一覧取得
	reqPost := &postpb.GetAllPostsByUserIDRequest{
		UserId: res.Id,
	}
	resPost, err := s.PostClient.GetAllPostsByUserID(r.Context(), reqPost)
	if err != nil {
		log.Println(err)
	}

	template.Render(w, "top/top.tmpl", resPost.Posts)
}

// UserRegister returns "user/register/init"
func (s *FrontServer) UserRegister(w http.ResponseWriter, r *http.Request) {
	template.RenderBeforeLogin(w, "user/userRegisterForm.tmpl", nil)
}

// UserRegisterConfirm returns "user/register/confirm"
func (s *FrontServer) UserRegisterConfirm(w http.ResponseWriter, r *http.Request) {
	var user database.User

	// フォームから取得した値
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	// MySQLと接続
	db := database.GormConnect()
	defer db.Close()

	// エラーハンドリング
	err := db.Find(&user, "email=?", email).Error
	if err == nil {
		// DBにフォームから来たメールが存在していたらリダイレクト
		// 存在していたら err == nil になる
		log.Println("This email is already registered")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
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
	reqCache := &cachepb.GetIDRequest{
		Key: LOGINUSER,
	}
	resCache, _ := s.CacheClient.GetID(r.Context(), reqCache)

	// 投稿を作成するgRPCリクエスト
	reqPost := &postpb.CreatePostRequest{
		Post: &postpb.Post{
			Url:         url,
			Title:       title,
			Description: description,
			Image:       image,
			Usecase:     usecase,
			Genre:       genre,
			UserId:      resCache.Id,
		},
	}
	resPost, err := s.PostClient.CreatePost(r.Context(), reqPost)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		http.Redirect(w, r, "/post/register/init", http.StatusFound)
		return
	}
	log.Println(resPost.GetMessage())

	http.Redirect(w, r, "/top", http.StatusFound)
}

// PostDelete deletes a post
func (s *FrontServer) PostDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postID := r.FormValue("post_id")

	if postID == "xxxx" {
		template.Render(w, "top/top.tmpl", nil)
	}

	req := &postpb.DeletePostRequest{
		Id: postID,
	}
	res, _ := s.PostClient.DeletePost(r.Context(), req)
	log.Panicln(res.GetMessage())
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
	// 投稿一覧取得
	reqPost := &postpb.SearchPostsByTitleRequest{
		UserId: res.Id,
		Title:  title,
	}
	resPost, _ := s.PostClient.SearchPostsByTitle(r.Context(), reqPost)

	template.Render(w, "top/top.tmpl", resPost.Posts)
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
	// 投稿一覧取得
	reqPost := &postpb.SearchPostsByUsecaseRequest{
		UserId:  res.Id,
		Usecase: usecase,
	}
	resPost, _ := s.PostClient.SearchPostsByUsecase(r.Context(), reqPost)

	template.Render(w, "top/top.tmpl", resPost.Posts)
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
	// 投稿一覧取得
	reqPost := &postpb.SearchPostsByGenreRequest{
		UserId: res.Id,
		Genre:  genre,
	}
	resPost, _ := s.PostClient.SearchPostsByGenre(r.Context(), reqPost)

	template.Render(w, "top/top.tmpl", resPost.Posts)
}
