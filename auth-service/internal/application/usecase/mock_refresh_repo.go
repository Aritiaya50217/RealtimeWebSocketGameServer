package usecase

import "time"

type MockRefreshRepo struct{}

func (m *MockRefreshRepo) Save(userID int64, token string, expiresAt time.Time) error { return nil }
func (m *MockRefreshRepo) Find(token string) (string, time.Time, error) {
	return "1", time.Now().Add(time.Hour), nil
}
func (m *MockRefreshRepo) Delete(token string) error { return nil }
