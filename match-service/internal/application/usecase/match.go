package usecase

import (
	"encoding/json"
	"errors"
	"realtime_web_socket_game_server/match-service/internal/domain"
	"realtime_web_socket_game_server/match-service/internal/infrastructure/database"
	"realtime_web_socket_game_server/match-service/internal/port"
	"time"

	"gorm.io/gorm"
)

type MatchUsecase struct {
	matchRepo   port.MatchRepository
	outboxRepo  port.OutboxRepository
	transaction database.Transaction
}

func NewMatchUsecase(matchRepo port.MatchRepository, outboxRepo port.OutboxRepository, transaction database.Transaction) *MatchUsecase {
	return &MatchUsecase{matchRepo: matchRepo, outboxRepo: outboxRepo, transaction: transaction}
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

func (uc *MatchUsecase) List(status string, limit, offset int) ([]*domain.Match, int64, error) {
	return uc.matchRepo.List(status, limit, offset)
}

func (uc *MatchUsecase) UpdateStatus(id int64, status string) (*domain.Match, error) {
	var updateMatch *domain.Match

	err := uc.transaction.WithTransaction(func(tx *gorm.DB) error {
		match, err := uc.matchRepo.GetByIDTx(tx, id)
		if err != nil {
			return err
		}

		if match == nil {
			return errors.New("match NotFound")
		}

		// check status
		if match.Status != domain.StatusCreated && status == "" {
			return errors.New("status cannot be started")
		}

		if match.Status == domain.StatusStarted {
			return errors.New("the current status is started")
		}

		_, err = uc.matchRepo.UpdateStatus(match.ID, status)
		if err != nil {
			return err
		}

		// save event to outbox
		match.Status = status
		payload, err := json.Marshal(match)
		if err != nil {
			return err
		}

		outboxEvent := &domain.OutboxEvent{
			AggregateID: match.ID,
			EventType:   "MatchStarted",
			Payload:     string(payload),
			ProcessedAt: time.Now(),
		}

		if err := uc.outboxRepo.Save(outboxEvent); err != nil {
			return err
		}

		updateMatch = match
		return nil
	})

	if err != nil {
		return nil, err
	}
	return updateMatch, nil
}
