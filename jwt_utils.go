package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func getAccessToken() string {
	accessToken := os.Getenv("ACCESS_TOKEN")
	if accessToken == "" {
		return "SECRET"
	} else {
		return accessToken
	}
}

func createToken(username string) (string, error) {
	var err error
	accessToken := getAccessToken()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(accessToken))
	if err != nil {
		return "", err
	}
	return token, nil
}

func extractTokenString(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	strArr := strings.Split(authHeader, " ")
	if len(strArr) != 2 {
		return "", fmt.Errorf("not a valid token")
	}
	return strArr[1], nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(getAccessToken()), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func getUsernamefromToken(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return fmt.Sprintf("%s", claims["username"]), nil
	}
	return "", fmt.Errorf("unknown error")
}
