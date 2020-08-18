package db

import (
  "github.com/gorilla/mux"
  "github.com/cgliu-create/potatoapi/middleware"
  "encoding/json"
  "net/http"
)

func createFunc(w http.ResponseWriter, r *http.Request) {
	var p Product
  json.NewDecoder(r.Body).Decode(&p)
	CreateProduct(&p)
  json.NewEncoder(w).Encode(&p)
}
func readAllFunc(w http.ResponseWriter, r *http.Request) {
	var p []Product
	ReadAllProduct(&p)
  json.NewEncoder(w).Encode(&p)
}
func readOneFunc(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
	var p Product
  ReadIDProduct(&p, params["id"])
	json.NewEncoder(w).Encode(&p)
}
func updateFunc(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
	var p Product
  ReadIDProduct(&p, params["id"])
  json.NewDecoder(r.Body).Decode(&p)
  UpdateProduct(&p, params["id"])
  json.NewEncoder(w).Encode(&p)
}
func deleteFunc(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
	var p Product
  DeleteProduct(&p, params["id"])
}

func AddAPIRoutes(r *mux.Router){
  apirouter := r.PathPrefix("/api").Subrouter()
  apirouter.HandleFunc("/create", middleware.Process(createFunc)).Methods("POST")
  apirouter.HandleFunc("/readall", middleware.Process(readAllFunc)).Methods("GET")
  apirouter.HandleFunc("/read/{id}", middleware.Process(readOneFunc)).Methods("GET")
  apirouter.HandleFunc("/update/{id}", middleware.Process(updateFunc)).Methods("PUT")
  apirouter.HandleFunc("/delete/{id}", middleware.Process(deleteFunc)).Methods("DELETE")
}
