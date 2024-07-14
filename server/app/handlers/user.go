package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hospital-system/server/app/dto"
	"net/http"
)

func (h *Handler) registerUser(c *gin.Context) {
	var request dto.RegisterUserRequest
	if err := c.BindJSON(&request); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.userService.CreateUser(dto.User{
		Firstname:                    request.Firstname,
		Lastname:                     request.Lastname,
		NationalIdentificationNumber: request.NationalIdentificationNumber,
		Email:                        request.Email,
		Role:                         request.Role,
		Team:                         request.Team,
	}); err != nil {
		h.log.Error("Failed to create user", zap.String("email", request.Email), zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusCreated, nil)
}

func (h *Handler) getUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		h.log.Error("Failed to get users", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}
