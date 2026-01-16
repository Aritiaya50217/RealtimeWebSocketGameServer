package repository

import "realtime_web_socket_game_server/match-service/internal/domain"

func ToOutboxDomain(m OutboxEventModel) *domain.OutboxEvent {
	return &domain.OutboxEvent{
		ID:          m.ID,
		AggregateID: m.AggregateID,
		EventType:   m.EventType,
		Payload:     m.Payload,
		Processed:   m.Processed,
		CreatedAt:   m.CreatedAt,
	}
}

func ToOutboxModel(d *domain.OutboxEvent) *OutboxEventModel {
	return &OutboxEventModel{
		ID:          d.ID,
		AggregateID: d.AggregateID,
		EventType:   d.EventType,
		Payload:     d.Payload,
		Processed:   d.Processed,
		CreatedAt:   d.CreatedAt,
	}
}
