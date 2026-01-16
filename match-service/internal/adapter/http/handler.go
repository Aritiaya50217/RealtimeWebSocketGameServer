package http

import (
	"fmt"
	"net/http"
	"realtime_web_socket_game_server/match-service/internal/application/usecase"
	"realtime_web_socket_game_server/match-service/internal/middleware"
	"strconv"

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
	match.GET("/:id", handler.GetMatchByID)
	match.GET("/list", handler.List)
	match.POST("/:id", handler.UpdateStatus)

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

func (h *MatchHandler) GetMatchByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid match id"})
		return
	}

	match, err := h.usecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if match == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "match not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"match": match})

}

func (h *MatchHandler) List(c *gin.Context) {
	status := c.Query("status")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	matches, total, err := h.usecase.List(status, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": total, "matches": matches, "limit": limit, "offset": offset})
}

func (h *MatchHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	status := c.Query("status")

	match, err := h.usecase.UpdateStatus(id, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	message := fmt.Sprintf("id %d update status success", match.ID)
	c.JSON(http.StatusOK, gin.H{"message": message})
}
