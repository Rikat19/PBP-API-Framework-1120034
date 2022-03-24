package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/gin/controller"
)

func main() {
	r := gin.Default()

	r.POST("/add", c.CreateUser)
	r.GET("/get", c.RetrieveUser)
	r.PUT("/update", c.UpdateUser)
	r.DELETE("/delete", c.DeleteUser)
	r.Run(":8080")
}
