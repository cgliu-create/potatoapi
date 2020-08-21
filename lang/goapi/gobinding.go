package goapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
  apiurl = "https://localhost:8000/api/products"
)

type Model struct {
	Name  string
	Price uint
}

type RequestManager struct {
  key string
  client *http.Client
}

// Authorize returns a api request manager with the given authorization token 
func Authorize(token string) RequestManager {
  return RequestManager{key: token, client: &http.Client{}}
}

func (r RequestManager) makeRequest(method, url string, body []byte) *http.Response {
  req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
  if err != nil {
    fmt.Printf("an error occured: %v", err)
  }
  req.Header.Set("Token", r.key)
  resp, err := r.client.Do(req)
  if err != nil {
    fmt.Printf("an error occured: %v", err)
  }
  return resp
}

// CreateProduct makes a request for a new Product instance of given values
func (r RequestManager) CreateProduct(name string, price uint) *http.Response {
  body, err := json.Marshal(Model{Name: name, Price: price})
  if err != nil {
    fmt.Printf("an error occured: %v", err)
  }
  return r.makeRequest("POST", apiurl, body)
}

// ReadAllProduct makes a request for all Product instances
func (r RequestManager) ReadAllProduct() *http.Response{
  return r.makeRequest("GET", apiurl, []byte(""))
}

// ReadProduct makes a request for the Product instance with the given pk
func (r RequestManager) ReadProduct(id uint) *http.Response{
  sapiurl := fmt.Sprintf("%v/%v", apiurl, id)
  return r.makeRequest("GET", sapiurl, []byte(""))
}

// UpdateProduct makes a request for replacing the Product instance with the given pk for a new Product instance of given values
func (r RequestManager) UpdateProduct(id uint, name string, price uint) *http.Response{
  body, err := json.Marshal(Model{Name: name, Price: price})
  if err != nil {
    fmt.Printf("an error occured: %v", err)
  }
  sapiurl := fmt.Sprintf("%v/%v", apiurl, id)
  return r.makeRequest("PUT", sapiurl, body)
}

// DeleteProduct makes a request for removing the Product instance with the given pk
func (r RequestManager) DeleteProduct(id uint) *http.Response{
  sapiurl := fmt.Sprintf("%v/%v", apiurl, id)
  return r.makeRequest("DELETE", sapiurl, []byte(""))
}
