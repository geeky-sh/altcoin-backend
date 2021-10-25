package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", testWorking)
	router.POST("/signup", signupUser)
	router.POST("/login", loginUser)
	router.POST("/validate", validateToken)
	router.POST("/buy", buyBitcoin)
	router.POST("/sell", sellBitcoin)

	router.GET("/users", getUsers)
	router.GET("/transactions", getTransactions)

	router.Run(":8000")
}
