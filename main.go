package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type User struct {
	Id int
	Username string
	Email string
	Address string
}

var users []User

func main() {
	router := gin.Default()

	router.GET("/randomuser", random_user)
	router.GET("/getuser/:id", get_user)
	router.POST("/adduser", add_user)

	router.Run(":8080")
}

func random_user(c *gin.Context) {
	if len(users) > 0 {
		user := users[rand.Intn(len(users))]
		c.JSON(http.StatusOK, gin.H{
			"random user": gin.H{
				"Id": user.Id,
				"Username": user.Username,
				"Email": user.Email,
				"Address": user.Address,
			},
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"random user": gin.H{
			"Id": -1,
		},
		"message": "No users exist",
	})
}

func add_user(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	address := c.PostForm("address")

	if len(username) > 0 && len(email) > 0 && len(address) > 0 {
		newuser := User{
			Id: len(users) + 1,
			Username: username,
			Email: email,
			Address: address,
		}
		users = append(users, newuser)

		msg := "Created user "+username+" ("+strconv.Itoa(newuser.Id)+") with email "+email+" and address "+address
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": msg,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": "POST request must include username, email and address fields",
		})
	}
}

func get_user(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	
	if len(users) > id && err == nil {
		user := users[id]
		c.JSON(http.StatusOK, gin.H{
			"random user": gin.H{
				"Id": user.Id,
				"Username": user.Username,
				"Email": user.Email,
				"Address": user.Address,
			},
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"random user": gin.H{
			"Id": -1,
		},
		"message": "No users exist",
	})
}
