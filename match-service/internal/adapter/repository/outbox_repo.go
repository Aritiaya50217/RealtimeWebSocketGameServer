package repository

import (
	"realtime_web_socket_game_server/match-service/internal/domain"
	"time"

	"gorm.io/gorm"
)

type OutboxRepository struct {
	db *gorm.DB
}

func NewOutboxRepository(db *gorm.DB) *OutboxRepository {
	return &OutboxRepository{db: db}
}

func (r *OutboxRepository) Save(event *domain.OutboxEvent) error {
	return r.db.Create(event).Error
}

func (r *OutboxRepository) FindUnprocessed(limit int) ([]*domain.OutboxEvent, error) {
	var events []*domain.OutboxEvent
	if err := r.db.Where("procressed = false").Limit(limit).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *OutboxRepository) MarkProcessed(eventID int64) error {
	return r.db.Model(&domain.OutboxEvent{}).Where("id = ?", eventID).Updates(map[string]interface{}{
		"processed":    true,
		"processed_at": time.Now(),
	}).Error
}
