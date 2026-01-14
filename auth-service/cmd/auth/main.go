package main

import (
	"log"
	"os"
	"realtime_web_socket_game_server/auth-service/internal/adapter/http"
	"realtime_web_socket_game_server/auth-service/internal/adapter/repository/postgres"
	"realtime_web_socket_game_server/auth-service/internal/application/usecase"
	"realtime_web_socket_game_server/auth-service/internal/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("AUTH_SERVICE_PORT")
	if port == "" {
		port = "8080"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	// connect to database
	db := database.NewPostgresDB()

	authRepo := postgres.NewUserRepository(db)
	refreshRepo := postgres.NewRefreshTokenRepository(db)

	refreshUsecase := usecase.NewRefreshTokenUsecase(refreshRepo, jwtSecret)
	loginUsecase := usecase.NewLoginUsecase(authRepo, refreshUsecase, jwtSecret)
	registerUsecase := usecase.NewRegisterUsecase(authRepo)

	r := gin.Default()
	handler := http.NewAuthHandler(loginUsecase, registerUsecase, refreshUsecase)

	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)
	r.POST("/refresh", handler.Refresh)

	r.Run(":" + port)

}
