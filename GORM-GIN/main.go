package main 

import "github.com/gin-gonic/gin"
func main() {

	r := gin.Default()
	r.GET("/hai", func(c*gin.Context) {
		c.String(200, "Hello world Balakrishna this is simple gin program")
	})
	r.Run()

	}
