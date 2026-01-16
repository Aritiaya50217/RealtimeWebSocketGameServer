package repository

import "time"

type RefreshModel struct {
	ID        int64      `gorm:"primaryKey;autoIncrement"`
	UserID    int64      `gorm:"not null;index"`
	User      *UserModel `gorm:"foreignKey:UserID"`
	Token     string     `gorm:"not null" json:"-"`
	ExpiresAt time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
