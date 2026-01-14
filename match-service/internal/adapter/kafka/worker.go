package kafka

import (
	"context"
	"log"
	"realtime_web_socket_game_server/match-service/internal/port"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

type OutboxWorker struct {
	outboxRepo port.OutboxRepository
	writer     *kafka.Writer
	interval   time.Duration
}

func NewOutboxWorker(outboxRepo port.OutboxRepository, writer *kafka.Writer, interval time.Duration) *OutboxWorker {
	return &OutboxWorker{outboxRepo: outboxRepo, writer: writer, interval: interval}
}

// Start the worker
func (w *OutboxWorker) Start() {
	go func() {
		for {
			events, err := w.outboxRepo.FindUnprocessed(100)
			if err != nil {
				log.Println("Outbox find error:", err)
				time.Sleep(w.interval)
				continue
			}

			for _, event := range events {
				// produce to kafka
				msg := kafka.Message{
					Key:   []byte(strconv.FormatInt(event.AggregateID, 10)),
					Value: []byte(event.Payload),
				}

				if err := w.writer.WriteMessages(context.Background(), msg); err != nil {
					log.Println("Kafka produce error:", err)
					continue
				}

				// mark processed
				if err := w.outboxRepo.MarkProcessed(event.ID); err != nil {
					log.Println("Failed to mark processed:", err)
				}
			}

			time.Sleep(w.interval)
		}
	}()
}
