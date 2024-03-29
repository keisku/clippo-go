package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/kskumgk63/clippo-go/front/proto/cachepb"
	"github.com/kskumgk63/clippo-go/front/proto/postpb"
	"github.com/kskumgk63/clippo-go/front/proto/userpb"
	"github.com/kskumgk63/clippo-go/front/template"
	"golang.org/x/crypto/bcrypt"
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
		log.Fatalln(err)
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
		if res == nil {
			log.SetFlags(log.Lshortfile)
			log.Println("*** JWT Token is empty ***")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
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
		}
	}
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
	resUser, err := s.UserClient.GetUser(r.Context(), reqUser)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	user := resUser.User

	// パスワードの正誤判断
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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
		log.Fatalln(err)
	}

	template.Render(w, "top/top.tmpl", resPost.Posts)
}
