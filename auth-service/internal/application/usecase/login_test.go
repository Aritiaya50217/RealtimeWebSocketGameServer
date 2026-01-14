package usecase

import (
	"realtime_web_socket_game_server/auth-service/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogin_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	jwtSecret := "secret123"
	uc := NewLoginUsecase(mockRepo, jwtSecret)

	hashedPassword, err := HashPassword("password123")
	require.NoError(t, err)

	require.True(t,CheckPassword(hashedPassword,"password123"))

	mockRepo.On("GetByUsername", "alice").Return(&domain.User{
		ID:       1,
		Username: "alice",
		Password: hashedPassword,
	}, nil).Once()

	token, err := uc.Login("alice", "password123")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}
