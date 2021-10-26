package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func testWorking(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"working": true})
}

func signupUser(c *gin.Context) {
	var userRequest UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser := findUserByUserName(userRequest.Username)
	if existingUser != nil {
		fmt.Println(existingUser)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "User already exists"})
		return
	}

	fmt.Println(userRequest)
	user := User{Username: userRequest.Username,
		Password: userRequest.Password, WalletAmount: 50000, BitcoinAmount: 0}
	users = append(users, user)

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func loginUser(c *gin.Context) {
	var userRequest UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser := findUserByUserName(userRequest.Username)
	if existingUser == nil || existingUser.Password != userRequest.Password {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "incorrect creds"})
		return
	}

	token, err := createToken(existingUser.Username)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "unable to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func validateToken(c *gin.Context) {
	var TokenRequest TokenRequest
	if err := c.ShouldBindJSON(&TokenRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, terr := verifyToken(TokenRequest.Token)
	if terr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func buyBitcoin(c *gin.Context) {
	userptr, err := doTokenAuthentication(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var transactionRequest TransactionRequest
	if err := c.ShouldBindJSON(&transactionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	canBuy := canBuyBitcoin(&transactionRequest, userptr)
	if !canBuy {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient funds"})
		return
	}

	if (*userptr).WalletAmount < transactionRequest.BitcoinValue {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient funds"})
		return
	}

	transactionptr := CreateTransaction(
		(*userptr).Username, "buy", transactionRequest.BitcoinAmount, transactionRequest.BitcoinValue)
	UpdateUser(userptr, transactionptr)

	c.JSON(http.StatusOK, userptr)
}

func sellBitcoin(c *gin.Context) {
	userptr, err := doTokenAuthentication(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var transactionRequest TransactionRequest
	if err := c.ShouldBindJSON(&transactionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	canSell := canSellBitcoin(&transactionRequest, userptr)
	if !canSell {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient funds"})
		return
	}

	transactionptr := CreateTransaction(
		(*userptr).Username, "sell", transactionRequest.BitcoinAmount, transactionRequest.BitcoinValue)
	UpdateUser(userptr, transactionptr)

	c.JSON(http.StatusOK, userptr)
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, &users)
}

func getTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, &transactions)
}
