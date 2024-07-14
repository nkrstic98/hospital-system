package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hospital-system/server/app/dto"
	"net/http"
)

func (h *Handler) registerUser(ctx *gin.Context) {
	var request dto.RegisterUserRequest
	if err := ctx.BindJSON(&request); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.userService.CreateUser(ctx, dto.User{
		Firstname:                    request.Firstname,
		Lastname:                     request.Lastname,
		NationalIdentificationNumber: request.NationalIdentificationNumber,
		Email:                        request.Email,
		Role:                         request.Role,
		Team:                         request.Team,
	}); err != nil {
		h.log.Error("Failed to create user", zap.String("email", request.Email), zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, nil)
}

func (h *Handler) getUsers(ctx *gin.Context) {
	users, err := h.userService.GetUsers(ctx)
	if err != nil {
		h.log.Error("Failed to get users", zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, users)
}

func (h *Handler) getDepartments(ctx *gin.Context) {
	var request dto.GetDepartmentsRequest
	if err := ctx.BindJSON(&request); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	response, err := h.userService.GetDepartments(ctx, request.Team, request.Role)
	if err != nil {
		h.log.Error("Failed to get departments response", zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, dto.GetDepartmentsResponse{Departments: response})
}
