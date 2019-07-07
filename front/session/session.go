package session

import (
	"fmt"
	"log"
	"net/http"
	"time"

	goCache "github.com/patrickmn/go-cache"
)

const (
	sessionExpires = 3 * time.Minute
	clearInterval  = 10 * time.Minute
	jwtToken       = "jwt-token"
)

type Store interface {
	Set(k string, x interface{})
	Delete(k string)
	Get(k string) (interface{}, bool)
}

func NewStoreOnMemory() *StoreOnMemory {
	return &StoreOnMemory{
		goCache.New(sessionExpires, clearInterval),
	}
}

type StoreOnMemory struct {
	cache *goCache.Cache
}

func (ss *StoreOnMemory) Set(k string, x interface{}) {
	ss.cache.Set(k, x, sessionExpires)
}

func (ss *StoreOnMemory) Delete(k string) {
	ss.cache.Delete(k)
}

func (ss *StoreOnMemory) Get(k string) (interface{}, bool) {
	return ss.cache.Get(k)
}

func GetTokenFromRequest(r *http.Request) string {
	c, err := r.Cookie(jwtToken)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Printf("*** %v\n", fmt.Sprint(err))
		return ""
	}
	return c.Value
}

func SetTokenToResponse(w http.ResponseWriter, token string) {
	fmt.Printf("session.go token =>> %v\n", token)
	http.SetCookie(w, &http.Cookie{
		Name:     jwtToken,
		Value:    token,
		HttpOnly: true,
	})
}

func DeleteTokenFromResponse(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     jwtToken,
		Value:    "",
		MaxAge:   0,
		HttpOnly: true,
	})
}
