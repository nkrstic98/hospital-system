package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	Authorization_MissingTokenErr         = "missing_token"
	Authorization_FailedToFetchSessionErr = "failed_to_fetch_session"
	Authorization_SessionExpiredErr       = "session_expired"
	Authorization_UserForbiddenError      = "user_not_allowed_to_execute_operation"
	Authorization_RefreshSessionErr       = "failed_to_refresh_session"
)

func (h *Handler) rbacAuthMiddleware(section, permission string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Fetch the Authorization header
		tokenString := ctx.GetHeader(AuthorizationCookieName)
		if tokenString == "" {
			h.log.Warn("Missing token in Authorization header")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": Authorization_MissingTokenErr})
			ctx.Abort()
			return
		}

		// Fetch the session from Redis store
		// If the session has expired, it will return nil for claims value
		claims, err := h.sessionService.GetSession(ctx, tokenString)
		if err != nil {
			h.log.Error("Failed to fetch session", zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": Authorization_FailedToFetchSessionErr})
			ctx.Abort()
			return
		}
		if claims == nil {
			h.log.Warn("Session expired")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": Authorization_SessionExpiredErr})
			ctx.Abort()
			return
		}

		if section != "%" {
			givenPermission, found := claims.Permissions[section]
			if !found || (permission == "WRITE" && givenPermission != "WRITE") {
				h.log.Warn("User is not allowed to access this resource",
					zap.Any("user", claims.UserID), zap.String("section", section))
				ctx.JSON(http.StatusForbidden, gin.H{"error": Authorization_UserForbiddenError})
				ctx.Abort()
				return
			}
		}

		// Refresh the session
		if err = h.sessionService.RefreshSession(ctx, tokenString); err != nil {
			h.log.Error("Failed to refresh session", zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": Authorization_RefreshSessionErr})
			ctx.Abort()
			return
		}

		// If everything goes well, allow user to execute the operation
		ctx.Next()
	}
}

func (h *Handler) basicAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Fetch the Authorization header
		tokenString := ctx.GetHeader(AuthorizationCookieName)
		if tokenString == "" {
			h.log.Warn("Missing token in Authorization header")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": Authorization_MissingTokenErr})
			ctx.Abort()
			return
		}

		// Fetch the session from Redis store
		// If the session has expired, it will return nil for claims value
		claims, err := h.sessionService.GetSession(ctx, tokenString)
		if err != nil {
			h.log.Error("Failed to fetch session", zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": Authorization_FailedToFetchSessionErr})
			ctx.Abort()
			return
		}
		if claims == nil {
			h.log.Warn("Session expired")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": Authorization_SessionExpiredErr})
			ctx.Abort()
			return
		}

		// Refresh the session
		if err = h.sessionService.RefreshSession(ctx, tokenString); err != nil {
			h.log.Error("Failed to refresh session", zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": Authorization_RefreshSessionErr})
			ctx.Abort()
			return
		}

		// If everything goes well, allow user to execute the operation
		ctx.Next()
	}
}
