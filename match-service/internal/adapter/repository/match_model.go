package repository

import "time"

type MatchModel struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	PlayerIDs []int64   `json:"player_ids" gorm:"-"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
