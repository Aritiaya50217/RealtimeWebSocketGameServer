package usecase

import (
	"realtime_web_socket_game_server/auth-service/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestRegister_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	uc := NewRegisterUsecase(mockRepo)

	mockRepo.
		On("GetByUsername", "alice").
		Return(nil, nil).
		Once()

	mockRepo.
		On(
			"CreateUser",
			mock.AnythingOfType("*domain.User"),
		).
		Return(nil).
		Once()

	// Act
	err := uc.Register("alice", "password123")

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegister_UserAlreadyExists(t *testing.T) {
	mockRepo := new(MockUserRepository)
	uc := NewRegisterUsecase(mockRepo)

	mockRepo.On("GetByUsername", "alice").
		Return(&domain.User{Username: "alice"}, nil).Once()

	err := uc.Register("alice", "password123")

	assert.Error(t, err)
	mockRepo.AssertNotCalled(t, "CreateUser")
}
