package repository

import (
	"errors"
	"realtime_web_socket_game_server/match-service/internal/domain"

	"gorm.io/gorm"
)

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (r *MatchRepository) Save(match *domain.Match) error {
	matchModel := ToMatchModel(match)
	err := r.db.Create(matchModel).Error
	if err != nil {
		return err
	}

	// Set auto-generated ID back to domain
	match.ID = matchModel.ID
	return nil
}

func (r *MatchRepository) GetByID(id int64) (*domain.Match, error) {
	var matchModel MatchModel
	if err := r.db.Where("id = ? ", id).First(&matchModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer NotFound error")
		}
		return nil, err
	}

	return ToMatchDomain(matchModel), nil
}
