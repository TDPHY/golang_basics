package utils

import (
	"blog/controllers"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwtKey 用于签名和验证JWT token
var jwtKey = []byte("your_secret_key")

// Claims 定义JWT声明结构
type Claims = controllers.Claims

// GenerateJWT 生成JWT token
func GenerateJWT(userID uint, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &controllers.Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

// ParseJWT 解析并验证JWT token
func ParseJWT(tokenString string) (*controllers.Claims, error) {
	claims := &controllers.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// ExtractToken 从Authorization头中提取JWT token
func ExtractToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return "", errors.New("bearer token is required")
	}

	return tokenString, nil
}