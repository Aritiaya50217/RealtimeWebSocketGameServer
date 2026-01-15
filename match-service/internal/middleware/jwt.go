package middleware

import (
	"net/http"
	"strings"
	"time"

	"realtime_web_socket_game_server/match-service/internal/helper"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(accessSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing token"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(
			tokenStr,
			&helper.AccessTokenClaims{},
			func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(accessSecret), nil
			},
			jwt.WithLeeway(5*time.Second), // buffer เล็ก ๆ ป้องกัน clock drift
		)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		claims := token.Claims.(*helper.AccessTokenClaims)

		if claims.Scope != "access" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "refresh token not allowed"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
