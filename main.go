package main

import (
	"inspirit/assignment/bitcoin/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type isWorking struct {
	Working bool `json:"working"`
}

func testWorking(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, isWorking{Working: true})
}

func main() {
	// var err error
	// models.ConnectToDatabase(models.DB, err)
	// models.DB.AutoMigrate(&models.User{}, &models.Transaction{})

	router := gin.Default()
	router.GET("/", testWorking)

	v1 := router.Group("/v1")
	v1.POST("signup", controllers.SignupUser)
	// v1.POST("validate", controllers.ValidateToken)
	// v1.POST("login", controllers.LoginUser)
	// v1.POST("buy", controllers.BuyBitcoin)
	// v1.POST("sell", controllers.SellBitcoin)

	router.Run("localhost:8000")
}
