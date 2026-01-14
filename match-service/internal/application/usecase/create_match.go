package usecase

import (
	"realtime_web_socket_game_server/match-service/internal/domain"
	"realtime_web_socket_game_server/match-service/internal/helper"
	"realtime_web_socket_game_server/match-service/internal/port"
	"time"
)

type MatchUsecase struct {
	matchRepo port.MatchRepository
	kafkaProd port.MatchKafkaProducer
}

func NewMatchUsecase(matchRepo port.MatchRepository, kafkaProd port.MatchKafkaProducer) *MatchUsecase {
	return &MatchUsecase{matchRepo: matchRepo, kafkaProd: kafkaProd}
}

func (uc *MatchUsecase) Create(playerIDs []string) (*domain.Match, error) {
	match := &domain.Match{
		ID:        helper.GenerateID(),
		PlayerIDs: playerIDs,
		CreatedAt: time.Now(),
	}
	if err := uc.matchRepo.Save(match); err != nil {
		return nil, err
	}

	// ส่ง event ไป game-service
	if err := uc.kafkaProd.ProduceMatchCreated(match); err != nil {
		return nil, err
	}

	return match, nil
}
