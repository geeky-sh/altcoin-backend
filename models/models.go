package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            uint
	Username      string
	Password      string
	WalletAmount  float32 `gorm:"default:50000"`
	BitcoinAmount float32 `gorm:"default:0"`
}

type Transaction struct {
	gorm.Model
	ID            uint
	Ttype         string
	User          User
	UserID        uint
	BitcoinAmount float32
	BitcoinPrice  float32
	CreatedAt     int64 `gorm:"autoCreateTime"`
}
