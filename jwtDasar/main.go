package main

import (
	"fmt"
	"net/http"

	// "strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.RegisteredClaims
}

func main() {
	r := gin.Default()
	r.POST("/login", Login)
	r.GET("/welcome", Welcome)
	r.Run(":8080")
}

func Login(c *gin.Context) {
	var u User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if u.Username != "yanto" || u.Password != "password" {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "gk dak ada data nya"})
		return
	}

	expireTime := time.Now().Add(25 * time.Minute)
	claim := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "yanto",
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Subject:   u.Username,
		},
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Welcome(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Welcome %s", claims.RegisteredClaims.Subject)})
}
