package usecase_test

import (
	"errors"
	"realtime_web_socket_game_server/match-service/internal/application/usecase"
	"realtime_web_socket_game_server/match-service/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ------------------------
// Mock Repositories
// ------------------------

type MockMatchRepo struct {
	mock.Mock
}

func (m *MockMatchRepo) Save(match *domain.Match) error {
	args := m.Called(match)
	if match != nil {
		match.ID = 100 // simulate auto increment ID
	}
	return args.Error(0)
}

func (m *MockMatchRepo) GetByID(id int64) (*domain.Match, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Match), args.Error(1)
}

type MockOutboxRepo struct {
	mock.Mock
}

func (m *MockOutboxRepo) Save(event *domain.OutboxEvent) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockOutboxRepo) FindUnprocessed(limit int) ([]*domain.OutboxEvent, error) {
	args := m.Called(limit)
	return args.Get(0).([]*domain.OutboxEvent), args.Error(1)
}

func (m *MockOutboxRepo) MarkProcessed(eventID int64) error {
	args := m.Called(eventID)
	return args.Error(0)
}

// ------------------------
// Test MatchUsecase.Create
// ------------------------

func TestMatchUsecase_Create_Success(t *testing.T) {
	playerIDs := []int64{1, 2, 3}

	mockMatchRepo := new(MockMatchRepo)
	mockOutboxRepo := new(MockOutboxRepo)

	uc := usecase.NewMatchUsecase(mockMatchRepo, mockOutboxRepo)

	mockMatchRepo.On("Save", mock.AnythingOfType("*domain.Match")).Return(nil)
	mockOutboxRepo.On("Save", mock.AnythingOfType("*domain.OutboxEvent")).Return(nil)

	match, err := uc.Create(playerIDs)

	assert.NoError(t, err)
	assert.NotNil(t, match)
	assert.Equal(t, int64(100), match.ID)
	assert.Equal(t, playerIDs, match.PlayerIDs)
	assert.Equal(t, "created", match.Status)

	mockMatchRepo.AssertCalled(t, "Save", mock.AnythingOfType("*domain.Match"))
	mockOutboxRepo.AssertCalled(t, "Save", mock.AnythingOfType("*domain.OutboxEvent"))
}

func TestMatchUsecase_Create_MatchRepoError(t *testing.T) {
	playerIDs := []int64{1, 2}

	mockMatchRepo := new(MockMatchRepo)
	mockOutboxRepo := new(MockOutboxRepo)

	uc := usecase.NewMatchUsecase(mockMatchRepo, mockOutboxRepo)

	mockMatchRepo.On("Save", mock.AnythingOfType("*domain.Match")).Return(errors.New("db error"))

	match, err := uc.Create(playerIDs)

	assert.Nil(t, match)
	assert.EqualError(t, err, "db error")
	mockMatchRepo.AssertCalled(t, "Save", mock.AnythingOfType("*domain.Match"))
	mockOutboxRepo.AssertNotCalled(t, "Save", mock.Anything)
}

func TestMatchUsecase_Create_OutboxRepoError(t *testing.T) {
	playerIDs := []int64{1, 2}

	mockMatchRepo := new(MockMatchRepo)
	mockOutboxRepo := new(MockOutboxRepo)

	uc := usecase.NewMatchUsecase(mockMatchRepo, mockOutboxRepo)

	mockMatchRepo.On("Save", mock.AnythingOfType("*domain.Match")).Return(nil)
	mockOutboxRepo.On("Save", mock.AnythingOfType("*domain.OutboxEvent")).Return(errors.New("outbox error"))

	match, err := uc.Create(playerIDs)

	assert.Nil(t, match)
	assert.EqualError(t, err, "outbox error")
	mockMatchRepo.AssertCalled(t, "Save", mock.AnythingOfType("*domain.Match"))
	mockOutboxRepo.AssertCalled(t, "Save", mock.AnythingOfType("*domain.OutboxEvent"))
}

// ------------------------
// Test OutboxRepository
// ------------------------

func TestOutboxRepository_FindUnprocessed_Success(t *testing.T) {
	mockRepo := new(MockOutboxRepo)

	events := []*domain.OutboxEvent{
		{AggregateID: 1, EventType: "MatchCreated", Payload: "{}"},
		{AggregateID: 2, EventType: "MatchCreated", Payload: "{}"},
	}

	mockRepo.On("FindUnprocessed", 10).Return(events, nil)

	result, err := mockRepo.FindUnprocessed(10)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, int64(1), result[0].AggregateID)
	assert.Equal(t, int64(2), result[1].AggregateID)
	mockRepo.AssertCalled(t, "FindUnprocessed", 10)
}

// func TestOutboxRepository_FindUnprocessed_Error(t *testing.T) {
// 	mockRepo := new(MockOutboxRepo)

// 	mockRepo.On("FindUnprocessed", 5).Return(nil, errors.New("db error"))

// 	result, err := mockRepo.FindUnprocessed(5)

// 	assert.Nil(t, result)
// 	assert.EqualError(t, err, "db error")
// 	mockRepo.AssertCalled(t, "FindUnprocessed", 5)
// }

func TestOutboxRepository_MarkProcessed_Success(t *testing.T) {
	mockRepo := new(MockOutboxRepo)

	mockRepo.On("MarkProcessed", int64(100)).Return(nil)

	err := mockRepo.MarkProcessed(100)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "MarkProcessed", int64(100))
}

func TestOutboxRepository_MarkProcessed_Error(t *testing.T) {
	mockRepo := new(MockOutboxRepo)

	mockRepo.On("MarkProcessed", int64(200)).Return(errors.New("db error"))

	err := mockRepo.MarkProcessed(200)

	assert.EqualError(t, err, "db error")
	mockRepo.AssertCalled(t, "MarkProcessed", int64(200))
}

func TestMatchUsecase_GetByID_Success(t *testing.T) {
	mockMatchRepo := new(MockMatchRepo)
	mockOutboxRepo := new(MockOutboxRepo)

	uc := usecase.NewMatchUsecase(mockMatchRepo, mockOutboxRepo)

	mockMatch := &domain.Match{
		ID:        1,
		PlayerIDs: []int64{1, 2},
		Status:    "created",
		CreatedAt: time.Now(),
	}

	mockMatchRepo.On("GetByID", int64(1)).Return(mockMatch, nil)

	match, err := uc.GetByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, match)
	mockMatchRepo.AssertExpectations(t)
}

func TestMatchUsecase_GetByID_NotFound(t *testing.T) {
	mockMatchRepo := new(MockMatchRepo)
	mockOutboxRepo := new(MockOutboxRepo)

	uc := usecase.NewMatchUsecase(mockMatchRepo, mockOutboxRepo)

	mockMatchRepo.On("GetByID", int64(99)).Return(nil, errors.New("customer NotFound error"))

	match, err := uc.GetByID(99)
	assert.Nil(t, match)
	assert.EqualError(t, err, "customer NotFound error")

	mockMatchRepo.AssertExpectations(t)
}

func TestMatchUsecase_GetByID_Error(t *testing.T) {
	mockMatchRepo := new(MockMatchRepo)
	mockOutboxRepo := new(MockOutboxRepo)

	uc := usecase.NewMatchUsecase(mockMatchRepo, mockOutboxRepo)

	mockMatchRepo.On("GetByID", int64(1)).Return(nil, errors.New("DB error"))

	match, err := uc.GetByID(1)
	assert.Error(t, err)
	assert.Nil(t, match)
	assert.Equal(t, "DB error", err.Error())

	mockMatchRepo.AssertExpectations(t)
}
