package repository

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) Save(userID int64, token string, expiresAt time.Time) error {
	refresh := &RefreshModel{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
	}
	return r.db.Create(refresh).Error
}

func (r *RefreshTokenRepository) Find(token string) (string, time.Time, error) {
	var refresh RefreshModel
	if err := r.db.Where("token = ?", token).First(&refresh).Error; err != nil {
		return "", time.Time{}, err
	}
	return strconv.FormatInt(refresh.UserID, 10), refresh.ExpiresAt, nil
}

func (r *RefreshTokenRepository) Delete(token string) error {
	return r.db.Where("token = ?", token).Delete(&RefreshModel{}).Error
}
