package postgres

import (
	"time"

	"realtime_web_socket_game_server/auth-service/internal/domain"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) Save(userID int64, token string, expiresAt time.Time) error {
	refresh := &domain.Refresh{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
	}
	return r.db.Create(refresh).Error
}

func (r *RefreshTokenRepository) Find(token string) (string, time.Time, error) {
	var refresh domain.Refresh
	if err := r.db.Where("token = ?", token).First(&refresh).Error; err != nil {
		return "", time.Time{}, err
	}
	return string(refresh.UserID), refresh.ExpiresAt, nil
}

func (r *RefreshTokenRepository) Delete(token string) error {
	return r.db.Where("token = ?", token).Delete(&domain.Refresh{}).Error
}
