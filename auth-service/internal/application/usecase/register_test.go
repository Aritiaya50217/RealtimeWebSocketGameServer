package usecase

import (
	"realtime_web_socket_game_server/auth-service/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
