package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupDB(host, port, user, dbname, password string)(*gorm.DB, error){
  db, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
  return db, err
}

