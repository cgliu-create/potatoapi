package middleware

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"
  "fmt"
  jwt "github.com/dgrijalva/jwt-go/v4"
)

var (
  logger = log.New(os.Stdout, " potato ", log.LstdFlags|log.Lshortfile)
  secret = []byte(os.Getenv("SECRET_TOKEN"))
)
func GenerateJWT() (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    tokenString, err := token.SignedString(secret)
    if err != nil {
        problem := fmt.Sprintf("there was an error: %s", err.Error())
        return problem, err
    }
    return tokenString, nil
}
func isAuthorized(f http.HandlerFunc, w http.ResponseWriter, r *http.Request) {
    if r.Header["Token"] != nil{
      token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
          return nil, errors.New("there was an error")
        }
        return secret, nil
      })
      if err != nil {
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.WriteHeader(500)
        w.Write([]byte(err.Error()))
      }
      if token.Valid {
        f(w, r)
      }
    } else {
      w.Header().Set("Content-Type", "text/plain; charset=utf-8")
      w.WriteHeader(401)
      w.Write([]byte("not authorized"))
    }
}

func Process(f http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    defer logger.Printf("%v %v processed in %v\n", r.Method, r.URL.Path, time.Since(start))
    isAuthorized(f, w, r)
  }
}

