package port

import (
	"realtime_web_socket_game_server/match-service/internal/domain"
)

type OutboxRepository interface {
	Save(event *domain.OutboxEvent) error
	FindUnprocessed(limit int) ([]*domain.OutboxEvent, error)
	MarkProcessed(eventID int64) error
}
