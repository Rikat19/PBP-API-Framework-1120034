package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/gin/model"
)

func CreateUser(c *gin.Context) {
	db := Connect()
	defer db.Close()

	query := "INSERT INTO users (uid, username, age) VALUES (?,?,?)"
	var user m.User
	err := c.Bind(&user)
	if err != nil {
		fmt.Print(err)
		return
	}
	db.Exec(query, user.UID, user.Username, user.Age)
	c.IndentedJSON(http.StatusOK, user)
}

func RetrieveUser(c *gin.Context) {
	db := Connect()
	defer db.Close()

	query := "SELECT * from users"
	result, errQ := db.Query(query)
	if errQ != nil {
		fmt.Print("Error", errQ.Error())
	}

	var user m.User
	var users []m.User
	for result.Next() {
		errQ = result.Scan(&user.UID, &user.Username, &user.Age)
		if errQ != nil {
			panic(errQ.Error())
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, users)
	}
}

func UpdateUser(c *gin.Context) {
	db := Connect()
	defer db.Close()

	var user m.User
	err := c.Bind(&user)
	if err != nil {
		fmt.Print(err)
		return
	}
	query := "UPDATE users SET username=?, age=? WHERE UID=?"
	result, errQ := db.Exec(query, user.Username, user.Age, user.UID)
	num, _ := result.RowsAffected()
	if errQ == nil {
		if num == 0 {
			c.AbortWithStatusJSON(400, "Bad Request")
			return
		} else {
			c.IndentedJSON(http.StatusOK, user)
		}
	}
}

func DeleteUser(c *gin.Context) {
	db := Connect()
	defer db.Close()

	uid := c.Query("UID")

	query := "DELETE FROM users WHERE uid=?"
	result, errQ := db.Exec(query, uid)
	num, _ := result.RowsAffected()
	if errQ == nil {
		if num == 0 {
			c.AbortWithStatusJSON(400, "Bad Request")
			return
		} else {
			c.IndentedJSON(http.StatusOK, uid)
		}
	}
}
