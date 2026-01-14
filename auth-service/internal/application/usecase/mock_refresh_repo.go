package usecase

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type MockRefreshRepo struct {
	mock.Mock
}

func (m *MockRefreshRepo) Save(userID int64, token string, expiresAt time.Time) error {
	return nil
}

func (m *MockRefreshRepo) Find(token string) (string, time.Time, error) {
	return "1", time.Now().Add(time.Hour), nil
}

func (m *MockRefreshRepo) Delete(token string) error {
	return nil
}

type ExpiredMockRefreshRepo struct {
	mock.Mock
}

func (m *ExpiredMockRefreshRepo) Save(userID int64, token string, expiresAt time.Time) error {
	return nil
}

func (m *ExpiredMockRefreshRepo) Find(token string) (string, time.Time, error) {
	// หมดอายุ
	return "1", time.Now().Add(-time.Hour), nil
}

func (m *ExpiredMockRefreshRepo) Delete(token string) error {
	return nil
}
