package usecase

import (
	"realtime_web_socket_game_server/auth-service/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogin_Success(t *testing.T) {
	// Mock user repository
	mockRepo := new(MockUserRepository)
	jwtSecret := "secret123"

	mockRefreshRepo := &MockRefreshRepo{}
	refreshUC := NewRefreshTokenUsecase(mockRefreshRepo, jwtSecret)

	// สร้าง LoginUsecase
	uc := NewLoginUsecase(mockRepo, refreshUC, jwtSecret)

	// Hash password
	hashedPassword, err := HashPassword("password123")
	require.NoError(t, err)

	require.True(t, CheckPassword(hashedPassword, "password123"))

	// Mock GetByUsername
	mockRepo.On("GetByUsername", "alice").Return(&domain.User{
		ID:       1,
		Username: "alice",
		Password: hashedPassword,
	}, nil).Once()

	// Act
	access_token, refresh_token, err := uc.Login("alice", "password123")

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, access_token)
	assert.NotEmpty(t, refresh_token)
	mockRepo.AssertExpectations(t)
}
