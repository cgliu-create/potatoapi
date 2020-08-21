import requests


apiurl = "https://localhost:8000/api/products"

class potato:
    def __init__(self):
        self.key = "" 
    def authorize(self, token):
        self.key = token
    def checkRequest(self, response, code):
        if response.status_code == code:
            print("Success")
        else:
            print(f"Something went wrong:{response.status_code}")
    def createProduct(self, name, price):
        headers = {"Token": self.key}
        data = {"Name": name, "Price": price}
        r = requests.post(apiurl, headers=headers, data=data)
        self.checkRequest(r, 201)
        return r
    def readAllProduct(self):
        headers = {"Token": self.key}
        r = requests.get(apiurl, headers=headers)
        self.checkRequest(r, 200)
        return r 
    def readProduct(self, id):
        headers = {"Token": self.key}
        sapiurl = f'{apiurl}/{id}'
        r = requests.get(sapiurl, headers=headers)
        self.checkRequest(r, 200)
        return r
    def updateProduct(self, id, name, price):
        headers = {"Token": self.key}
        data = {"Name": name, "Price": price}
        sapiurl = f'{apiurl}/{id}'
        r = requests.put(sapiurl, headers=headers, data=data) 
        self.checkRequest(r, 200)
        return r
    def deleteProduct(self, id):
        headers = {"Token": self.key}
        sapiurl = f'{apiurl}/{id}'
        r = requests.deleteProduct(sapiurl, headers=headers) 
        self.checkRequest(r, 200)
        return r

