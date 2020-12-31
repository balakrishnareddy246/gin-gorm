package main 

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON (200, gin.H {
			"message": "Hello world balakrishna",
		})
		})
		r.Run() //listen and server on 8080
	}
