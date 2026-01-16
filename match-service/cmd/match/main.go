package main

import (
	"log"
	"os"
	"time"

	httpAdapter "realtime_web_socket_game_server/match-service/internal/adapter/http"
	kafkaAdapter "realtime_web_socket_game_server/match-service/internal/adapter/kafka"
	repoAdapter "realtime_web_socket_game_server/match-service/internal/adapter/repository"
	"realtime_web_socket_game_server/match-service/internal/application/usecase"
	"realtime_web_socket_game_server/match-service/internal/infrastructure/database"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

func main() {
	port := os.Getenv("MATCH_SERVICE_PORT")
	if port == "" {
		port = "8081"
	}

	// ----------------------------
	// Connect to database
	db := database.NewPostgresDB()
	log.Println("Database connected successfully")

	transaction := database.NewTransaction(db)

	// ----------------------------
	// Repository
	matchRepo := repoAdapter.NewMatchRepository(db)
	outboxRepo := repoAdapter.NewOutboxRepository(db)

	// ----------------------------
	// Kafka admin: ensure topic exists
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	topicMatchCreated := os.Getenv("KAFKA_TOPIC_MATCH_CREATED")
	kafkaAdmin := kafkaAdapter.NewKafkaAdmin(kafkaBroker)

	if err := kafkaAdmin.EnsureTopic(topicMatchCreated, 1); err != nil {
		log.Fatalf("Failed to ensure topic %s: %v", topicMatchCreated, err)
	}
	log.Printf("Kafka topic %s ready", topicMatchCreated)

	// ----------------------------
	// Kafka writer / producer
	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    topicMatchCreated,
		Balancer: &kafka.LeastBytes{},
	}
	// ----------------------------
	// Usecase
	uc := usecase.NewMatchUsecase(matchRepo, outboxRepo, *transaction)

	r := gin.Default()
	// ----------------------------
	// HTTP Handler
	httpAdapter.NewMatchHandler(r, uc, os.Getenv("JWT_SECRET"))

	// Start outbox worker
	outboxWorker := kafkaAdapter.NewOutboxWorker(outboxRepo, writer, 5*time.Second)
	outboxWorker.Start()

	log.Println("match-service running on port :" + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
