package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
 _"github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB 
var err error

type Person struct {
	ID uint `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastnme"`
	Email  string `json:"email"`

}
fucn main() {
  db, err = gorm.Open("sqlite3", "./gorm.db")
  if err != nil {
	  db, err = gorm.Open("sqlite3", "./gorm.db")
	  if err = nil {
		  fmt.Println(err)

	  }
	  defer db.Close()
	  db.AutoMigrate(&Person{})
	  r := gin.Default()
	  r.GET("/", GetProjects)
	  r.Run(":8080")
  }
  func GetPeojects (c *gin.Context) {
	  var people []Person 

	  if err := db.Find(&people) .Error; err != nil {
		  c.AbortWithStatus(404)
		  fmt.Println(err)
	  } else {
		  c.JSON( 200, people)
	  }
	  }
  
  