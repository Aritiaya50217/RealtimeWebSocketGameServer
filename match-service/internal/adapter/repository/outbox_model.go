package repository

import "time"

type OutboxEventModel struct {
	ID          int64 `json:"id" gorm:"primaryKey"`
	AggregateID int64
	EventType   string
	Payload     string // JSON string
	Processed   bool
	CreatedAt   time.Time
	ProcessedAt time.Time
}
