package domain

import "time"

type Match struct {
	ID        int64     `json:"id"`
	PlayerIDs []int64   `json:"player_ids"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
