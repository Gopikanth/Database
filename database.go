package main

import (
       "github.com/gin-gonic/gin"
       "github.com/jinzhu/gorm"
       _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
func init() {
 //open a db connection
 var err error
 db, err = gorm.Open("mysql", "save:12345@/demo?charset=utf8&parseTime=True&loc=Local")

//mysql is a database driver
//save is database username
//12345 is password
//demo is database name
 if err != nil {
  panic("failed to connect database") // handling the error
 }

 db.AutoMigrate(&todoModel{}) //Migrate the schema
}