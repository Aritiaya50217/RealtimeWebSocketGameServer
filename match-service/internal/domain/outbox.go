package domain

import "time"

type OutboxEvent struct {
	ID          int64 `gorm:"primaryKey"`
	AggregateID int64
	EventType   string
	Payload     string // JSON string
	Processed   bool
	CreatedAt   time.Time
	ProcessedAt time.Time
}
