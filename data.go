package main

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TransactionRequest struct {
	BitcoinValue  float32 `json:"bitcoin_value"`
	BitcoinAmount float32 `json:"bitcoin_amount"`
}

type TokenRequest struct {
	Token string `json:"token"`
}

type User struct {
	Username      string  `json:"username"`
	Password      string  `json:"-"`
	WalletAmount  float32 `json:"wallet_amount"`
	BitcoinAmount float32 `json:"bitcoin_amount"`
}

type Transaction struct {
	Ttype         string  `json:"ttype"`
	Username      string  `json:"username"`
	BitcoinAmount float32 `json:"bitcoin_amount"`
	BitcoinValue  float32 `json:"bitcoin_value"`
	CreatedAt     int64   `json:"created_at"`
}

var users = []User{}
var transactions = []Transaction{}

func findUserByUserName(username string) *User {
	var t User
	for i := 0; i < len(users); i++ {
		t = users[i]
		if t.Username == username {
			return &t
		}
	}
	return nil
}

func CreateTransaction(username string, ttype string, bitcoinAmount float32, BitcoinValue float32) *Transaction {
	transaction := Transaction{
		Username: username, BitcoinValue: BitcoinValue,
		BitcoinAmount: bitcoinAmount, Ttype: ttype}
	transactions = append(transactions, transaction)
	return &transaction
}

func UpdateUser(user *User, transaction *Transaction) {
	if transaction.Ttype == "sell" {
		user.BitcoinAmount = user.BitcoinAmount - transaction.BitcoinAmount
		user.WalletAmount = user.WalletAmount + transaction.BitcoinValue
	} else {
		user.BitcoinAmount = user.BitcoinAmount + transaction.BitcoinAmount
		user.WalletAmount = user.WalletAmount - transaction.BitcoinValue
	}

	var t User
	for i := 0; i < len(users); i++ {
		t = users[i]
		if t.Username == user.Username {
			users[i] = *user
		}
	}
}
