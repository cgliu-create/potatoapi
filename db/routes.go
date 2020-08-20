package db

import (
  "github.com/gorilla/mux"
  "github.com/cgliu-create/potatoapi/middleware"
  "encoding/json"
  "net/http"
)

func createProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
  json.NewDecoder(r.Body).Decode(&p)
  err := CreateProduct(&p)
  if err == nil{
    w.WriteHeader(201)
  } else {
    w.WriteHeader(500)
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(&p)
}
func readAllProduct(w http.ResponseWriter, r *http.Request) {
	var p []Product
  err := ReadAllProduct(&p)
  if err == nil{
    w.WriteHeader(200)
  } else {
    w.WriteHeader(500)
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(&p)
}
func readOneProduct(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
	var p Product
  err := ReadIDProduct(&p, params["id"])
  if err == nil{
    w.WriteHeader(200)
  } else {
    w.WriteHeader(500)
  }
  w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&p)
}
func updateProduct(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
	var p Product
  err := ReadIDProduct(&p, params["id"])
  if err != nil{
    w.WriteHeader(404)
  }
  json.NewDecoder(r.Body).Decode(&p)
  err = UpdateProduct(&p, params["id"])
  w.Header().Set("Content-Type", "text/plain; charset=utf-8")
  if err == nil{
    w.WriteHeader(200)
	  json.NewEncoder(w).Encode(&p)
  } else {
    w.WriteHeader(304)
    w.Write([]byte("no change"))
  }
}
func deleteProduct(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
	var p Product
  err := DeleteProduct(&p, params["id"])
  w.Header().Set("Content-Type", "text/plain; charset=utf-8")
  if err == nil{
    w.WriteHeader(200)
	  json.NewEncoder(w).Encode(&p)
  } else {
    w.WriteHeader(304)
    w.Write([]byte("no change"))
  }
}

func productFunc(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET" {
    readOneProduct(w, r)
  }
  if r.Method == "PUT"{
    updateProduct(w, r)
  }
  if r.Method == "DELETE"{
    deleteProduct(w, r)
  }
}

func productsFunc(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET"{
    readAllProduct(w, r)
  }
  if r.Method == "POST" {
    createProduct(w, r)
  }
}

func AddAPIRoutes(r *mux.Router){
  apirouter := r.PathPrefix("/api").Subrouter()
  apirouter.HandleFunc("/products/{id}", middleware.Process(productFunc)).Methods("GET", "PUT", "DELETE")
  apirouter.HandleFunc("/products", middleware.Process(productsFunc)).Methods("GET", "POST")
}
