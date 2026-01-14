package domain

import "time"

type Match struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	PlayerIDs []string  `json:"player_ids" gorm:"-"`
	CreatedAt time.Time `json:"created_at"`
}
