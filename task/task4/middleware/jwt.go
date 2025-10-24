package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("CHANGE_ME_256bit_random_string")

type Claims struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

// 生成 Token（双 Token 方案可同理扩展）
func GenerateToken(userID uint, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
}

// JWT 中间件：验证 access_token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "缺失 Authorization"})
			c.Abort()
			return
		}
		// 去掉 Bearer
		if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
			tokenStr = tokenStr[7:]
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Token 无效或已过期"})
			c.Abort()
			return
		}

		// 将 user_id 写入上下文，后续接口直接取
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
