package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/gin/controller"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	r.POST("/add", c.CreateUser)
	r.GET("/get", c.RetrieveUser)
	r.PUT("/update", c.UpdateUser)
	r.DELETE("/delete", c.DeleteUser)
	r.Run(":8080")
}
