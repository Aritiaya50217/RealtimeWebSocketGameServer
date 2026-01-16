package port

import (
	"realtime_web_socket_game_server/match-service/internal/domain"

	"gorm.io/gorm"
)

type MatchRepository interface {
	Save(match *domain.Match) error
	GetByID(id int64) (*domain.Match, error)
	List(status string, limit, offset int) ([]*domain.Match, int64, error)
	UpdateStatus(id int64, status string) (*domain.Match, error)

	GetByIDTx(tx *gorm.DB, id int64) (*domain.Match, error)
	UpdateStatusTx(tx *gorm.DB, id int64, status string) (*domain.Match, error)
}

type MatchKafkaProducer interface {
	ProduceMatchCreated(match *domain.Match) error
}
