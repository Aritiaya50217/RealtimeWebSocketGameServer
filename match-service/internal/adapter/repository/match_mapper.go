package repository

import "realtime_web_socket_game_server/match-service/internal/domain"

func ToMatchDomain(m MatchModel) *domain.Match {
	return &domain.Match{
		ID:        m.ID,
		PlayerIDs: m.PlayerIDs,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
	}
}

func ToMatchModel(d *domain.Match) *MatchModel {
	return &MatchModel{
		ID:        d.ID,
		PlayerIDs: d.PlayerIDs,
		Status:    d.Status,
		CreatedAt: d.CreatedAt,
	}
}
