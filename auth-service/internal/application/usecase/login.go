package usecase

import (
	"context"
	"errors"
	"realtime_web_socket_game_server/auth-service/internal/adapter/repository/postgres"
	"realtime_web_socket_game_server/auth-service/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginUsecase struct {
	repo      *postgres.UserRepository
	JWTSecret string
}

func NewLoginUsecase(repo *postgres.UserRepository, jwt string) *LoginUsecase {
	return &LoginUsecase{repo: repo, JWTSecret: jwt}
}

func (uc *LoginUsecase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := uc.repo.GetByUsername(ctx, username)
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

func (uc *LoginUsecase) Register(ctx context.Context, username, password string) error {
	hashed, err := HashPassword(password)
	if err != nil {
		return err
	}
	user := &domain.User{
		Username: username,
		Password: hashed,
	}
	return uc.repo.CreateUser(ctx, user)
}
