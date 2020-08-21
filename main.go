package main

import (
	"log"
	"github.com/cgliu-create/potatoapi/middleware"
	"github.com/cgliu-create/potatoapi/server"
	"github.com/cgliu-create/potatoapi/db"
	"github.com/gorilla/mux"
  "net/http"
  "fmt"
)

var (
	certFile       = "./cert/server.crt"
	keyFile        = "./cert/server.key"
	serviceAddress = ":8000"
)
var (
  host = "localhost"
  port = "5432"
  user = "cgliu"
  dbname = "potato"
  password = "1234"
)
func main() {
  var err error
	router := mux.NewRouter()
  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello world")
  })
  db.Database, err = db.SetupDB(host, port, user, dbname, password)
  if err != nil {
		log.Fatalf("database failed to start: %v", err)
	}
  db.MigrateProduct()
  db.AddAPIRoutes(router)
	srv := server.New(router, serviceAddress)
	log.Println("server starting")

  token, err := middleware.GenerateJWT()
  if err != nil {
    log.Println("jwt error")
  }
  log.Println(token)

	err = srv.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
