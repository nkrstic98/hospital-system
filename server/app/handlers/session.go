package handlers

import (
	"hospital-system/server/app/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

const (
	Session_Login_BadRequestErr            = "bad_request"
	Session_Login_UserPasswordPairWrongErr = "user_password_pair_wrong"
	Session_Login_UserNotFoundErr          = "user_not_found"
	Session_Login_CreateSessionFailedErr   = "create_session_failed"

	Session_Logout_BadRequestErr          = "bad_request"
	Session_Logout_DeleteSessionFailedErr = "delete_session_failed"

	Session_Validate_BadRequestErr           = "bad_request"
	Session_Validate_SessionGetFailedErr     = "session_get_failed"
	Session_Validate_SessionGetUserFailedErr = "session_get_user_failed"
	Session_Validate_SessionExpiredErr       = "session_expired"
	Session_Validate_UserNotFoundErr         = "user_not_found"
)

func (h *Handler) login(ctx *gin.Context) {
	var request dto.LoginRequest
	if err := ctx.BindJSON(&request); err != nil {
		h.log.Error("Failed to bind JSON login request", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": Session_Login_BadRequestErr})
		return
	}

	existingUser, err := h.userService.GetByUsername(ctx, request.Username)
	if err != nil {
		h.log.Error("Failed to get user by username", zap.String("username", request.Username), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": Session_Login_UserPasswordPairWrongErr})
		return
	}
	if existingUser == nil {
		h.log.Warn("User not found", zap.String("username", request.Username))
		ctx.JSON(http.StatusNotFound, gin.H{"error": Session_Login_UserNotFoundErr})
		return
	}

	isPasswordValid, err := h.userService.ValidateUserPassword(ctx, existingUser.ID, request.Password)
	if err != nil {
		h.log.Error("Failed to validate user password", zap.String("username", request.Username), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": Session_Login_UserNotFoundErr})
		return
	}
	if !isPasswordValid {
		h.log.Warn("User password pair is wrong", zap.String("username", request.Username))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": Session_Login_UserPasswordPairWrongErr})
		return
	}

	jwtToken, err := h.sessionService.CreateSession(ctx, dto.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: existingUser.ID.String(),
		},
		UserID:                       existingUser.ID,
		Username:                     existingUser.Username,
		NationalIdentificationNumber: existingUser.NationalIdentificationNumber,
		Email:                        existingUser.Email,
		Firstname:                    existingUser.Firstname,
		Lastname:                     existingUser.Lastname,
		Role:                         existingUser.Role,
		Team:                         existingUser.Team,
		Permissions:                  existingUser.Permissions,
	})
	if err != nil {
		h.log.Error("Failed to create session", zap.String("username", request.Username), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": Session_Login_CreateSessionFailedErr})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(AuthorizationCookieName, jwtToken, 900, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, dto.LoginResponse{
		Token: jwtToken,
		User:  *existingUser,
	})
}

func (h *Handler) logout(ctx *gin.Context) {
	var request dto.LogoutRequest
	if err := ctx.BindJSON(&request); err != nil {
		h.log.Error("Failed to bind JSON logout request", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": Session_Logout_BadRequestErr})
		return
	}

	if err := h.sessionService.DeleteSession(ctx, request.Token); err != nil {
		h.log.Error("Failed to delete session", zap.String("token", request.Token), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": Session_Logout_DeleteSessionFailedErr})
		return
	}

	ctx.SetCookie(AuthorizationCookieName, "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, nil)
}

func (h *Handler) validateSession(ctx *gin.Context) {
	var request dto.ValidateSessionRequest
	if err := ctx.BindJSON(&request); err != nil {
		h.log.Error("Failed to bind JSON validate session request", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": Session_Validate_BadRequestErr})
		return
	}

	claims, err := h.sessionService.GetSession(ctx, request.Token)
	if err != nil {
		h.log.Error("Failed to get session", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": Session_Validate_SessionGetFailedErr})
		return
	}
	if claims == nil {
		h.log.Warn("Session expired", zap.String("token", request.Token))
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": Session_Validate_SessionExpiredErr})
		return
	}

	existingUser, err := h.userService.GetByUsername(ctx, claims.Username)
	if err != nil {
		h.log.Error("Failed to get user by username", zap.String("username", claims.Username), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": Session_Validate_SessionGetUserFailedErr})
		return
	}
	if existingUser == nil {
		h.log.Warn("User not found", zap.String("username", claims.Username))
		ctx.JSON(http.StatusNotFound, gin.H{"error": Session_Validate_UserNotFoundErr})
		return
	}

	ctx.JSON(http.StatusOK, dto.ValidateSessionResponse{
		User: *existingUser,
	})
}
