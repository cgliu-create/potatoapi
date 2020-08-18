package middleware

import (
  "os"
  "log"
  "net/http"
  "time"
)

var (
  logger = log.New(os.Stdout, " potato ", log.LstdFlags|log.Lshortfile)
)

func Process(f http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    defer logger.Printf("%v %v processed in %v\n", r.Method, r.URL.Path,time.Since(start))
    f(w, r)
  }
}

