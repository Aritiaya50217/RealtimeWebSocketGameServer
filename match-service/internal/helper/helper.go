package helper

import (
	"fmt"
	"realtime_web_socket_game_server/match-service/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	UserID string `json:"user_id"`
	Scope  string `json:"scope"` // ต้องเป็น "access"
	jwt.RegisteredClaims
}

// serialize payload
func BuildMatchCreatedPayload(match *domain.Match) string {
	payload := fmt.Sprintf(`{"match_id":"%v","player_ids":[`, match.ID)
	for i, id := range match.PlayerIDs {
		payload += fmt.Sprintf(`%v`, id)
		if i < len(match.PlayerIDs)-1 {
			payload += ","
		}
	}
	payload += "]}"
	return payload
}
