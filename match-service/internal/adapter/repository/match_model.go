package repository

import "time"

type MatchModel struct {
	ID        int64     `gorm:"primaryKey"`
	PlayerIDs string    `gorm:"column:player_ids"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
