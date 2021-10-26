package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func doTokenAuthentication(c *gin.Context) (*User, error) {
	tokenString, err := extractTokenString(c)
	if err != nil {
		return nil, err
	}
	token, terr := verifyToken(tokenString)
	if terr != nil {
		return nil, terr
	}
	username, uerr := getUsernamefromToken(token)
	if uerr != nil {
		return nil, uerr
	}
	userptr := findUserByUserName(username)
	if userptr == nil {
		return nil, fmt.Errorf("no user found")
	}
	return userptr, nil
}

func canBuyBitcoin(transactionRequest *TransactionRequest, user *User) bool {
	return user.WalletAmount >= transactionRequest.BitcoinValue
}

func canSellBitcoin(transactionRequest *TransactionRequest, user *User) bool {
	bitcoinAmount := transactionRequest.BitcoinAmount
	return user.BitcoinAmount >= bitcoinAmount
}
