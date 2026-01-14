package kafka

import (
	"encoding/json"
	"realtime_web_socket_game_server/match-service/internal/domain"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(writer *kafka.Writer) *KafkaProducer {
	return &KafkaProducer{writer: writer}
}

func (p *KafkaProducer) ProduceMatchCreated(match *domain.Match) error {
	data, err := json.Marshal(match)

	if err != nil {
		return err
	}
	return p.writer.WriteMessages(nil, kafka.Message{
		Key:   []byte(match.ID),
		Value: data,
	})
}
