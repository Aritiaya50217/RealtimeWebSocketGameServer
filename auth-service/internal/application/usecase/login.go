package usecase

import (
	"errors"
	"realtime_web_socket_game_server/auth-service/internal/port"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginUsecase struct {
	repo      port.UserRepository
	JWTSecret string
}

func NewLoginUsecase(repo port.UserRepository, jwt string) *LoginUsecase {
	return &LoginUsecase{repo: repo, JWTSecret: jwt}
}

func (uc *LoginUsecase) Login(username, password string) (string, error) {
	user, err := uc.repo.GetByUsername(username)
	if err != nil {
		return "", err
	}

	if !CheckPassword(user.Password, password) {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.JWTSecret))
}
