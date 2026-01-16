package repository

import "time"

type OutboxEventModel struct {
	ID          int64     `gorm:"primaryKey"`
	AggregateID int64     `gorm:"column:aggregate_id"`
	EventType   string    `gorm:"column:event_type"`
	Payload     string    `gorm:"column:payload;type:json"`
	Processed   bool      `gorm:"column:processed;default:false"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	ProcessedAt time.Time `gorm:"column:processed_at"`
}
