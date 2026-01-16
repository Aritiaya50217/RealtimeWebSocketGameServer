package port

import "realtime_web_socket_game_server/match-service/internal/domain"

type MatchRepository interface {
	Save(match *domain.Match) error
	GetByID(id int64) (*domain.Match, error)
	List(status string, limit, offset int) ([]*domain.Match, int64, error)
}

type MatchKafkaProducer interface {
	ProduceMatchCreated(match *domain.Match) error
}
