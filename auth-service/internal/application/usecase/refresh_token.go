package usecase

import (
	"errors"
	"strconv"
	"time"

	"realtime_web_socket_game_server/auth-service/internal/port"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenUsecase struct {
	tokenRepo port.RefreshTokenRepository
	jwtSecret string
}

func NewRefreshTokenUsecase(tokenRepo port.RefreshTokenRepository, jwtSecret string) *RefreshTokenUsecase {
	return &RefreshTokenUsecase{tokenRepo: tokenRepo, jwtSecret: jwtSecret}
}

// Generate refresh token
func (uc *RefreshTokenUsecase) Generate(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(uc.jwtSecret))
	if err != nil {
		return "", err
	}

	userid, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return "", err
	}

	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	if err := uc.tokenRepo.Save(userid, tokenStr, expiresAt); err != nil {
		return "", err
	}

	return tokenStr, nil
}

// Refresh returns new access token
func (uc *RefreshTokenUsecase) Refresh(refreshToken, accessSecret string) (string, error) {
	userID, expiresAt, err := uc.tokenRepo.Find(refreshToken)
	if err != nil {
		return "", err
	}

	if time.Now().After(expiresAt) {
		return "", errors.New("refresh token expired")
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessSecret))
}

// Delete refresh token (logout)
func (uc *RefreshTokenUsecase) Delete(refreshToken string) error {
	return uc.tokenRepo.Delete(refreshToken)
}
