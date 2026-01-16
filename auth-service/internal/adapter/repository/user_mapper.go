package repository

import "realtime_web_socket_game_server/auth-service/internal/domain"

func ToUserDomain(m UserModel) *domain.User {
	return &domain.User{
		ID:        m.ID,
		Username:  m.Username,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToUserModel(d *domain.User) *UserModel {
	return &UserModel{
		ID:        d.ID,
		Username:  d.Username,
		Password:  d.Password,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
