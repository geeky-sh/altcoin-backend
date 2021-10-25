package services

import (
	"fmt"
	"inspirit/assignment/bitcoin/models"
	"inspirit/assignment/bitcoin/serializers"
)

func CreateUserService(userRequest *serializers.UserRequest) {
	user := *models.GetUserByUsername((*userRequest).Username)
	fmt.Println(user)
	if (user == models.User{}) {
		panic("user already exists")
	}
	models.CreateUser(user.Username, user.Password)
}

func CanBuyBitcoin(user models.User, transactionRequest serializers.TransactionRequest) bool {
	transactionAmount := transactionRequest.BitcoinAmount * transactionRequest.BitcoinPrice
	return user.WalletAmount >= transactionAmount
}

func CanSellBitcoin(user models.User, transactionRequest serializers.TransactionRequest) bool {
	bitcoinAmount := transactionRequest.BitcoinAmount

	return user.BitcoinAmount <= bitcoinAmount
}

func BuyBitcoin(user *models.User, transactionRequest serializers.TransactionRequest) {
	transaction := models.CreateTransaction(
		user, "buy", transactionRequest.BitcoinAmount, transactionRequest.BitcoinPrice)
	models.UpdateUser(user, transaction)
}

func SellBitcoin(user *models.User, transactionRequest serializers.TransactionRequest) {
	transaction := models.CreateTransaction(
		user, "sell", transactionRequest.BitcoinAmount, transactionRequest.BitcoinPrice)
	models.UpdateUser(user, transaction)
}
