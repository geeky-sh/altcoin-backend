package serializers

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TransactionRequest struct {
	BitcoinAmount float32 `json:"bitcoin_amount"`
	BitcoinPrice  float32 `json:"bitcoin_price"`
}

type UserResponse struct {
	Username      string  `json:"username"`
	WalletAmount  float32 `json:"wallet_amount"`
	BitcoinAmount float32 `json:"bitcoin_amount"`
}

// func ConvertUserModelToResponse(user models.User) UserResponse {
// 	up := UserResponse{}
// 	up.Username = user.Username
// 	up.WalletAmount = user.WalletAmount
// 	up.BitcoinAmount = user.BitcoinAmount
// 	return up
// }
