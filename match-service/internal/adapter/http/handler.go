package http

import (
	"net/http"
	"realtime_web_socket_game_server/match-service/internal/application/usecase"
	"realtime_web_socket_game_server/match-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	usecase *usecase.MatchUsecase
}

func NewMatchHandler(r *gin.Engine, usecase *usecase.MatchUsecase, jwtSecret string) *MatchHandler {
	handler := &MatchHandler{
		usecase: usecase,
	}

	match := r.Group("/match")
	match.Use(middleware.JWTMiddleware(jwtSecret))
	match.POST("/create", handler.CreateMatch)

	return handler
}

func (h *MatchHandler) CreateMatch(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var req struct {
		PlayerIDs []int64 `json:"player_ids" binding:"required,min=2"`
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
