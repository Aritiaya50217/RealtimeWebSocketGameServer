package http

import (
	"net/http"
	"realtime_web_socket_game_server/auth-service/internal/application/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	loginUsecase    *usecase.LoginUsecase
	registerUsecase *usecase.RegisterUsecase
	refreshUsecase  *usecase.RefreshTokenUsecase
	accessSecret    string
}

func NewAuthHandler(uc *usecase.LoginUsecase, registerUsecase *usecase.RegisterUsecase, refreshUsecase *usecase.RefreshTokenUsecase) *AuthHandler {
	return &AuthHandler{loginUsecase: uc, registerUsecase: registerUsecase, refreshUsecase: refreshUsecase}
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

	accessToken, refreshToken, err := h.loginUsecase.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	type req struct {
		RefreshToken string `json:"refresh_token"`
	}
	var r req
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAccess, err := h.refreshUsecase.Refresh(r.RefreshToken, h.accessSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": newAccess})
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
