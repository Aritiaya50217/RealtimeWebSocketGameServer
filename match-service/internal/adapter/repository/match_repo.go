package repository

import (
	"realtime_web_socket_game_server/match-service/internal/domain"

	"gorm.io/gorm"
)

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (r *MatchRepository) Save(match *domain.Match)error{
	return r.db.Create(match).Error
}
