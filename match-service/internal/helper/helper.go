package helper

import (
	"fmt"
	"realtime_web_socket_game_server/match-service/internal/domain"
)

// serialize payload
func BuildMatchCreatedPayload(match *domain.Match) string {
	payload := fmt.Sprintf(`{"match_id":"%s","player_ids":[`, match.ID)
	for i, id := range match.PlayerIDs {
		payload += fmt.Sprintf(`%d`, id)
		if i < len(match.PlayerIDs)-1 {
			payload += ","
		}
	}
	payload += "]}"
	return payload
}
