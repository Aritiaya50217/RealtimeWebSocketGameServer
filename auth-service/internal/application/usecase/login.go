package usecase

import (
	"errors"
	"realtime_web_socket_game_server/auth-service/internal/helper"
	"realtime_web_socket_game_server/auth-service/internal/port"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginUsecase struct {
	repo          port.UserRepository
	refresUsecase *RefreshTokenUsecase
	JWTSecret     string
}

func NewLoginUsecase(repo port.UserRepository, refresUsecase *RefreshTokenUsecase, jwt string) *LoginUsecase {
	return &LoginUsecase{repo: repo, refresUsecase: refresUsecase, JWTSecret: jwt}
}

func (uc *LoginUsecase) Login(username, password string) (string, string, error) {
	user, err := uc.repo.GetByUsername(username)
	if err != nil {
		return "", "", err
	}

	if !CheckPassword(user.Password, password) {
		return "", "", errors.New("invalid credentials")
	}

	// generate access token ด้วย struct
	claims := helper.AccessTokenClaims{
		UserID: strconv.Itoa(int(user.ID)),
		Scope:  "access",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // access token 15 นาที
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(uc.JWTSecret))
	if err != nil {
		return "", "", err
	}

	// generate refresh token
	userID := strconv.Itoa(int(user.ID))

	refreshToken, err := uc.refresUsecase.Generate(userID)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
