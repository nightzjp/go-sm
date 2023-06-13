/*
   @TIME: 2023/6/13 12:06
   @Author: Nightz
   @Site:
   @File: jwt.go
   @SoftWare: GoLand
*/

package middleware

import (
	"net/http"
	"strings"
	"yishi-ai.com/go-sm/constant"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTMiddleware 是用于验证和解析JWT的中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取JWT令牌
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未登录"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "JWT ", "", 1)

		// 解析和验证JWT令牌
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(constant.JwtSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token错误"})
			c.Abort()
			return
		}

		// 将解析后的令牌存储在上下文中，以便后续处理程序使用
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["userID"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

/*

 */
