package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", testWorking)
	v1 := router.Group("api/v1")

	v1.POST("/signup", signupUser)
	v1.POST("/login", loginUser)
	v1.POST("/validate", validateToken)
	v1.POST("/buy", buyBitcoin)
	v1.POST("/sell", sellBitcoin)

	// urls for debugging purposes
	v1.GET("/users", getUsers)
	v1.GET("/transactions", getTransactions)

	router.Run(":8000")
}
