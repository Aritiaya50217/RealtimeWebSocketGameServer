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
	return r.db.Create(ToOutboxModel(event)).Error
}

func (r *OutboxRepository) FindUnprocessed(limit int) ([]*domain.OutboxEvent, error) {
	var eventsModel []OutboxEventModel
	if err := r.db.Where("processed = false").Limit(limit).Find(&eventsModel).Error; err != nil {
		return nil, err
	}

	events := make([]*domain.OutboxEvent, 0, len(eventsModel))
	for _, m := range eventsModel {
		events = append(events, ToOutboxDomain(m))
	}

	return events, nil
}

func (r *OutboxRepository) MarkProcessed(eventID int64) error {
	var outboxModel OutboxEventModel
	return r.db.Model(&outboxModel).Where("id = ?", eventID).Updates(map[string]interface{}{
		"processed":    true,
		"processed_at": time.Now(),
	}).Error
}

func (r *OutboxRepository) SaveTx(tx *gorm.DB, event *domain.OutboxEvent) error {
	return tx.Create(&OutboxEventModel{
		AggregateID: event.AggregateID,
		EventType:   event.EventType,
		Payload:     event.Payload,
		Processed:   false,
	}).Error
}
