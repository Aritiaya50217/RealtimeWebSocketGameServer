package usecase

// import (
// 	"context"
// 	"realtime_web_socket_game_server/auth-service/internal/domain"
// 	"testing"

// 	"github.com/stretchr/testify/mock"
// )

// type MockUserRepo struct {
// 	mock.Mock
// }

// func (m *MockUserRepo) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
// 	if user, ok := m.Users[username]; ok {
// 		return user, nil
// 	}
// 	return nil, nil
// }

// func (m *MockUserRepo) CreateUser(user *domain.User) error {
// 	args := m.Called(user)
// 	return args.Error(0)
// }

// func (m *MockUserRepo) FindEmail(email string) (*domain.User, error) {
// 	args := m.Called(email)
// 	if args.Get(0) != nil {
// 		return args.Get(0).(*domain.User), args.Error(1)
// 	}
// 	return nil, args.Error(1)
// }

// func TestRegisterUsecase(t *testing.T) {
// 	mockRepo := new(MockUserRepo)
// 	uc := NewRegisterUsecase(mockRepo)

// 	_, _ = mockRepo, uc
// }
