package main

import (
	"net/http"
	"io/ioutil"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/getdoge", getDoge)

	router.Run(":8082")
}

func getDoge(c *gin.Context) {
	currentTime := time.Now()
	// https://api.coingecko.com/api/v3/simple/price?ids=dogecoin&vs_currencies=usd
	response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=dogecoin&vs_currencies=usd")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error": err.Error(),
		},
		)
	} else {
		out, err := ioutil.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Error": err.Error(),
			},
			)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Doge/USD": string(out),
			},
			)
		}
	}
	
}
