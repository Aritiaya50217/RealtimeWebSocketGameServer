package port

import (
	"realtime_web_socket_game_server/auth-service/internal/domain"
)

type UserRepository interface {
	GetByUsername(username string) (*domain.User, error)
	CreateUser(user *domain.User) error
}
