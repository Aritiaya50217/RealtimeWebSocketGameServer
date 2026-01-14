package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"

	httpAdapter "realtime_web_socket_game_server/match-service/internal/adapter/http"
	kafkaAdapter "realtime_web_socket_game_server/match-service/internal/adapter/kafka"
	"realtime_web_socket_game_server/match-service/internal/adapter/repository"
	"realtime_web_socket_game_server/match-service/internal/application/usecase"
	"realtime_web_socket_game_server/match-service/internal/infrastructure/database"
)

func main() {

	port := os.Getenv("MATCH_SERVICE_PORT")
	if port == "" {
		port = "8081"
	}

	// connect to database
	db := database.NewPostgresDB()
	// repository
	matchRepo := repository.NewMatchRepository(db)

	// kafka
	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "match_created",
	}
	producer := kafkaAdapter.NewKafkaProducer(writer)

	// usecase
	uc := usecase.NewMatchUsecase(matchRepo, producer)

	// handler
	handler := httpAdapter.NewMatchHandler(uc)

	r := gin.Default()
	r.POST("/match/create", handler.CreateMatch)

	log.Println("match-service running on :" + port)
	r.Run(":" + port)
}
