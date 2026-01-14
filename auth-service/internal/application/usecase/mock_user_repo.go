package usecase

import (
	"realtime_web_socket_game_server/auth-service/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	user, _ := args.Get(0).(*domain.User)
	return user, args.Error(1)
}

func (m *MockUserRepository) CreateUser(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}
