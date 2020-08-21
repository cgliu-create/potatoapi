const axios = require('axios');

var key = ""
var apiurl = "https://localhost:8000/api/products"

function authorize(token){
  key = token
}

function createProduct(name, price){
  let data = axios({
    method: 'post',
    url: apiurl,
    data: {'Name':name, 'Price':price},
    headers: {
      'Content-Type': 'application/json',
      'Token': key
    }
  }).then((response) => {
    return response.data;
  });
  return data;
}

function readAllProduct(){
  let data = axios({
    method: 'get',
    url: apiurl,
    data: {},
    headers: {
      'Token': key
    }
  }).then((response) => {
    return response.data;
  });
  return data;
}

function readProduct(id){
  sapiurl = `${apiurl}/${id}`
  let data = axios({
    method: 'get',
    url: sapiurl,
    data: {},
    headers: {
      'Token': key
    }
  }).then((response) => {
    return response.data;
  });
  return data;   
}

function updateProduct(id, name, price){
  sapiurl = `${apiurl}/${id}`
  let data = axios({
    method: 'put',
    url: sapiurl,
    data: {'Name':name, 'Price':price},
    headers: {
      'Content-Type': 'application/json',
      'Token': key
    }
  }).then((response) => {
    return response.data;
  });
  return data;
}

function deleteProduct(id){
  sapiurl = `${apiurl}/${id}`
  let data = axios({
    method: 'delete',
    url: sapiurl,
    data: {},
    headers: {
      'Token': key
    }
  }).then((response) => {
    return response.data;
  });
  return data;
}
