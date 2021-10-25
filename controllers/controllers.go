package controllers

import (
	"fmt"
	"inspirit/assignment/bitcoin/models"
	"inspirit/assignment/bitcoin/serializers"
	"inspirit/assignment/bitcoin/services"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func retError(msg string) gin.H {
	return gin.H{"error": msg}
}

func retEmpty() gin.H {
	return gin.H{}
}

func SignupUser(c *gin.Context) {
	var userRequest serializers.UserRequest = serializers.UserRequest{Username: "ash", Password: "123"}
	fmt.Println(userRequest)
	c.BindJSON(&userRequest)

	services.CreateUserService(&userRequest)

	c.JSON(http.StatusOK, retEmpty())
}

func LoginUser(c *gin.Context) {
	var userRequest serializers.UserRequest
	c.BindJSON(userRequest)

	user, is_err := models.GetUserByNamePassword(userRequest.Username, userRequest.Password)
	if is_err {
		c.JSON(http.StatusNotFound, retError("User Not found"))
		return
	}

	token, is_terr := createToken(user.ID)
	if is_terr != nil {
		c.JSON(http.StatusUnprocessableEntity, retError("Unable to create token"))
		return
	}

	c.JSON(http.StatusOK, token)
}

func BuyBitcoin(c *gin.Context) {
	user_id, err := getUserIdfromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Invalid token")
		return
	}
	// validate input
	var transactionRequest serializers.TransactionRequest
	c.BindJSON(transactionRequest)

	user := models.GetUserById(user_id)
	if !services.CanBuyBitcoin(user, transactionRequest) {
		c.JSON(http.StatusUnprocessableEntity, retError("Insufficient funds"))
		return
	}
	services.BuyBitcoin(&user, transactionRequest)

	c.JSON(http.StatusOK, serializers.ConvertUserModelToResponse(user))
}

func SellBitcoin(c *gin.Context) {
	user_id, err := getUserIdfromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Invalid token")
		return
	}
	// validate input
	var transactionRequest serializers.TransactionRequest
	c.BindJSON(transactionRequest)

	user := models.GetUserById(user_id)
	if !services.CanSellBitcoin(user, transactionRequest) {
		c.JSON(http.StatusUnprocessableEntity, retError("Insufficient funds"))
		return
	}
	services.SellBitcoin(&user, transactionRequest)
	c.JSON(http.StatusOK, serializers.ConvertUserModelToResponse(user))
}

func ValidateToken(c *gin.Context) {
	_, err := getUserIdfromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Invalid token")
		return
	}
	c.JSON(http.StatusOK, retEmpty())
}

func getAccessToken() string {
	accessToken := os.Getenv("ACCESS_TOKEN")
	if accessToken == "" {
		return "SECRET"
	} else {
		return accessToken
	}
}

func createToken(userId uint) (string, error) {
	var err error
	accessToken := getAccessToken()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(accessToken))
	if err != nil {
		return "", err
	}
	return token, nil
}

func extractToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func verifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := extractToken(c)
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

func getUserIdfromToken(c *gin.Context) (uint, error) {
	token, err := verifyToken(c)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if !ok {
			return 0, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return uint(userId), nil
	}
	return 0, err
}
