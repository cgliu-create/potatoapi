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
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    json.NewEncoder(w).Encode(&p)
  } else {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(500)
    w.Write([]byte("Product not created"))
  }
}
func readAllProduct(w http.ResponseWriter, r *http.Request) {
	var p []Product
  err := ReadAllProduct(&p)
  if err == nil{
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    json.NewEncoder(w).Encode(&p)
  } else {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(404)
    w.Write([]byte("Product not found"))
  }
}
func readOneProduct(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
	var p Product
  err := ReadIDProduct(&p, params["id"])
  if err == nil{
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
	  json.NewEncoder(w).Encode(&p)
  } else {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(404)
    w.Write([]byte("Product not found"))
  }
}
func updateProduct(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
	var p Product
  err := ReadIDProduct(&p, params["id"])
  if err != nil{
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(404)
    w.Write([]byte("Product not found"))
  } else {
    json.NewDecoder(r.Body).Decode(&p)
    err = UpdateProduct(&p, params["id"])
    if err == nil{
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(200)
	    json.NewEncoder(w).Encode(&p)
    } else {
      w.Header().Set("Content-Type", "text/plain; charset=utf-8")
      w.WriteHeader(304)
      w.Write([]byte("Product not changed"))
    }
  }
}
func deleteProduct(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
	var p Product
  err := DeleteProduct(&p, params["id"])
  w.Header().Set("Content-Type", "text/plain; charset=utf-8")
  if err == nil{
    w.WriteHeader(200)
    w.Write([]byte("Product removed"))
  } else {
    w.WriteHeader(404)
    w.Write([]byte("Product not found"))
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
