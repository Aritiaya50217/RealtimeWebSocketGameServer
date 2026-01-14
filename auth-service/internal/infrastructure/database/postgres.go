package database

import (
	"fmt"
	"log"
	"os"
	"realtime_web_socket_game_server/auth-service/internal/domain"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("database environment variables are not set properly")
	}

	if sslmode == "" {
		sslmode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC",
		host, port, user, password, dbname, sslmode,
	)

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ { // retry 10 ครั้ง
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			break
		}
		log.Println("waiting for database to be ready...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// AutoMigrate
	if err := db.AutoMigrate(&domain.User{}, &domain.Refresh{}); err != nil {
		log.Fatalf("failed to auto-migrate: %v", err)
	}

	log.Println("database connected and migrated successfully")
	return db
}
