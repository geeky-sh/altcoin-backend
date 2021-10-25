package models

import _ "gorm.io/driver/sqlite"

func CreateUser(username string, password string) *User {
	user := User{Username: username, Password: password}
	DB.Create(&user)
	return &user
}

func UpdateUser(user *User, transaction Transaction) {
	if transaction.Ttype == "sell" {
		user.BitcoinAmount = user.BitcoinAmount - transaction.BitcoinAmount
		user.WalletAmount = user.WalletAmount + (transaction.BitcoinAmount * transaction.BitcoinPrice)
	} else {
		user.BitcoinAmount = user.BitcoinAmount + transaction.BitcoinAmount
		user.WalletAmount = user.WalletAmount - (transaction.BitcoinAmount * transaction.BitcoinPrice)
	}
	DB.Save(&user)
}

func CreateTransaction(user *User, ttype string, bitcoinAmount float32, bitcoinPrice float32) Transaction {
	transaction := Transaction{
		User: *user, BitcoinPrice: bitcoinPrice,
		BitcoinAmount: bitcoinAmount, Ttype: ttype}
	DB.Save(&transaction)
	return transaction
}

func GetUserByUsername(username string) *User {
	var user User
	DB.Where(&User{Username: username}).First(&user)
	return &user
}

func GetUserById(id uint) User {
	user := User{}
	DB.First(&user, id)
	return user
}

func GetUserByNamePassword(username string, password string) (User, bool) {
	user := User{}
	result := DB.Where(&User{Username: username, Password: password}).First(user)
	return user, (result.RowsAffected > 0)
}
