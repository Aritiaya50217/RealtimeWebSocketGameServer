package repository

import (
	"realtime_web_socket_game_server/auth-service/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUsername(username string) (*domain.User, error) {
	var userModel UserModel
	if err := r.db.Where("username = ?", username).First(&userModel).Error; err != nil {
		return nil, err
	}
	user := ToUserDomain(userModel)

	return user, nil
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	userModel := ToUserModel(user)
	return r.db.Create(userModel).Error
}
