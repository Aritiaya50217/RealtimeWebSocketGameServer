package port

import "time"

type RefreshTokenRepository interface {
	Save(userID int64, token string, expiresAt time.Time) error
	Find(token string) (userID string, expiresAt time.Time, err error)
	Delete(token string) error
}
