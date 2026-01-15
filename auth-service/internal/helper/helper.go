package helper

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaims struct {
	UserID string `json:"user_id"`
	Scope  string `json:"scope"` // ต้องเป็น "access"
	jwt.RegisteredClaims
}
