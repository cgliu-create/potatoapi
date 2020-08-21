
# Authorization
```shell
# With shell, pass the correct header with each request
curl "api_endpoint_here"
  -H "Token: potato"
```

# CreateProduct
```shell
 curl -d '{"Name":"abcd", "Price":1234}' 
  -H "Content-Type: application/json" -H "Token: potato" 
  -X POST "https://localhost:8000/api/products"
```

# ReadAllProduct
```shell
 curl 
  -H "Token: potato" 
  -X GET "https://localhost:8000/api/products"
```

# ReadProduct
```shell
 curl 
  -H "Token: potato" 
  -X GET "https://localhost:8000/api/products/1"
```

# UpdateProduct
```shell
 curl -d '{"Name":"abcd", "Price":4321}' 
  -H "Content-Type: application/json" -H "Token: potato" 
  -X PUT "https://localhost:8000/api/products/1"
```

# DeleteProduct
```shell
 curl 
  -H "Token: potato" 
  -X DELETE "https://localhost:8000/api/products/1"
```




