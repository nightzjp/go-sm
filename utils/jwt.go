/*
   @TIME: 2023/6/13 12:10
   @Author: Nightz
   @Site:
   @File: jwt.go
   @SoftWare: GoLand
*/

package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"yishi-ai.com/go-sm/constant"
)

// GenerateJWT 生成JWT令牌
func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := []byte(constant.JwtSecret)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

/*

 */
