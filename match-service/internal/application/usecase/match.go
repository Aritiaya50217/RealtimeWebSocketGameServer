package usecase

import (
	"encoding/json"
	"realtime_web_socket_game_server/match-service/internal/domain"
	"realtime_web_socket_game_server/match-service/internal/port"
	"time"
)

type MatchUsecase struct {
	matchRepo  port.MatchRepository
	outboxRepo port.OutboxRepository
}

func NewMatchUsecase(matchRepo port.MatchRepository, outboxRepo port.OutboxRepository) *MatchUsecase {
	return &MatchUsecase{matchRepo: matchRepo, outboxRepo: outboxRepo}
}

func (uc *MatchUsecase) Create(playerIDs []int64) (*domain.Match, error) {
	match := &domain.Match{
		PlayerIDs: playerIDs,
		Status:    "created",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.matchRepo.Save(match); err != nil {
		return nil, err
	}

	// save event to outbox
	payload, err := json.Marshal(match)
	if err != nil {
		return nil, err
	}

	outboxEvent := &domain.OutboxEvent{
		AggregateID: match.ID,
		EventType:   "MatchCreated",
		Payload:     string(payload),
		ProcessedAt: time.Now(),
		CreatedAt:   time.Now(),
	}

	if err := uc.outboxRepo.Save(outboxEvent); err != nil {
		return nil, err
	}

	return match, nil
}

func (uc *MatchUsecase) GetByID(id int64) (*domain.Match, error) {
	return uc.matchRepo.GetByID(id)
}
