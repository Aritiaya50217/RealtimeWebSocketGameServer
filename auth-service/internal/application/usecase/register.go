package usecase

import (
	"errors"
	"realtime_web_socket_game_server/auth-service/internal/domain"
	"realtime_web_socket_game_server/auth-service/internal/port"
)

type RegisterUsecase struct {
	repo port.UserRepository
}

func NewRegisterUsecase(repo port.UserRepository) *RegisterUsecase {
	return &RegisterUsecase{repo: repo}
}

func (uc *RegisterUsecase) Register(email, password string) error {
	if existing, _ := uc.repo.GetByUsername(email); existing != nil {
		return errors.New("email already exists")
	}

	hashed, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Username: email,
		Password: hashed,
	}

	return uc.repo.CreateUser(user)
}
