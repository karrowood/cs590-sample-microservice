package main

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/currenttime", current_time)

	router.Run(":8081")
}

func current_time(c *gin.Context) {
	currentTime := time.Now()
	c.JSON(http.StatusOK, gin.H{
		"time": currentTime.Format("2006-01-02 3:4:5 pm"),
		},
	)
}

