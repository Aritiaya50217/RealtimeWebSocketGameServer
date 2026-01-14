package http

import (
	"net/http"
	"realtime_web_socket_game_server/match-service/internal/application/usecase"

	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	usecase *usecase.MatchUsecase
}

func NewMatchHandler(usecase *usecase.MatchUsecase) *MatchHandler {
	return &MatchHandler{usecase: usecase}
}

func (h *MatchHandler) CreateMatch(c *gin.Context) {
	var req struct {
		PlayerIDs []string `json:"player_ids" binding:"required,min=2"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	match, err := h.usecase.Create(req.PlayerIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, match)
}
