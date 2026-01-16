package repository

import "realtime_web_socket_game_server/auth-service/internal/domain"

func ToRefreshDomain(m RefreshModel) *domain.Refresh {
	return &domain.Refresh{
		ID:        m.ID,
		UserID:    m.User.ID,
		Token:     m.Token,
		ExpiresAt: m.ExpiresAt,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToRefreshModel(d *domain.Refresh) *RefreshModel {
	return &RefreshModel{
		ID:        d.ID,
		UserID:    d.UserID,
		Token:     d.Token,
		ExpiresAt: d.ExpiresAt,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
