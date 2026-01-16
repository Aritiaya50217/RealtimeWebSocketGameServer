package port

import (
	"realtime_web_socket_game_server/match-service/internal/domain"

	"gorm.io/gorm"
)

type OutboxRepository interface {
	Save(event *domain.OutboxEvent) error
	FindUnprocessed(limit int) ([]*domain.OutboxEvent, error)
	MarkProcessed(eventID int64) error
	SaveTx(tx *gorm.DB, event *domain.OutboxEvent) error
}
