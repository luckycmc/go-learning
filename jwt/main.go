package main

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	APP_KEY = "www.baidu.com"
)

func main() {
	http.HandleFunc("/token", TokenHandler)
	http.Handle("/", AuthMiddleware(http.HandlerFunc(ExampleHandler2)))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if username != "myusername" || password != "mypassword" {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"error":"invalid_credentials"}`)
		return
	}

	// 颁发有效期为一小时的证书
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"ext":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(APP_KEY))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"internal_server_error"}`)
		return
	}
	io.WriteString(w, `{"token":"`+tokenString+`"}`)
	return
}

func AuthMiddleware(next http.Handler) http.Handler {
	if len(APP_KEY) == 0 {
		log.Fatal("http server unable to start,expected an app_key for jwt auth")
	}
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(APP_KEY), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return jwtMiddleware.Handler(next)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}

func ExampleHandler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok111"}`)
}
