package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRefreshToken_Generate(t *testing.T) {
	jwtSecret := "secret123"
	mockRepo := &MockRefreshRepo{}

	refreshUseacse := NewRefreshTokenUsecase(mockRepo, jwtSecret)

	token, err := refreshUseacse.Generate("1")
	require.NoError(t, err)
	assert.NotEmpty(t, token, "refresh token should be generated")
}

func TestRefreshToken_Refresh(t *testing.T) {
	jwtSecret := "secret123"
	mockRepo := &MockRefreshRepo{}

	refreshUsecase := NewRefreshTokenUsecase(mockRepo, jwtSecret)

	accessToken, err := refreshUsecase.Refresh("refresh_token", "xxx")
	require.NoError(t, err)
	assert.NotEmpty(t, accessToken, "new access token should be returned")
}

func TestRefreshToken_Expired(t *testing.T) {
	mockRepo := &ExpiredMockRefreshRepo{}
	jwtSecret := "secret123"
	refreshUC := NewRefreshTokenUsecase(mockRepo, jwtSecret)

	_, err := refreshUC.Refresh("dummy_token", jwtSecret)
	assert.Error(t, err)
	assert.Equal(t, "refresh token expired", err.Error())
}
