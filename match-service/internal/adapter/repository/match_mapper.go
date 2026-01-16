package repository

import (
	"encoding/json"
	"realtime_web_socket_game_server/match-service/internal/domain"
)

func ToMatchDomain(m MatchModel) *domain.Match {
	var playerIDs []int64
	_ = json.Unmarshal([]byte(m.PlayerIDs), &playerIDs)

	return &domain.Match{
		ID:        m.ID,
		PlayerIDs: playerIDs,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToMatchModel(d *domain.Match) *MatchModel {
	playerIDs, _ := json.Marshal(d.PlayerIDs)
	return &MatchModel{
		ID:        d.ID,
		PlayerIDs: string(playerIDs),
		Status:    d.Status,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
