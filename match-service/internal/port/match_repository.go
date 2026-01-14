package port

import "realtime_web_socket_game_server/match-service/internal/domain"

type MatchRepository interface {
	Save(match *domain.Match) error
}

type MatchKafkaProducer interface {
	ProduceMatchCreated(match *domain.Match) error
}
