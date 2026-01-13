package http

import (
	"net/http"
	"realtime_web_socket_game_server/auth-service/internal/application/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	loginUsecase    *usecase.LoginUsecase
	registerUsecase *usecase.RegisterUsecase
}

func NewAuthHandler(uc *usecase.LoginUsecase, registerUsecase *usecase.RegisterUsecase) *AuthHandler {
	return &AuthHandler{loginUsecase: uc, registerUsecase: registerUsecase}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	token, err := h.loginUsecase.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.registerUsecase.Register(req.Username, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}
