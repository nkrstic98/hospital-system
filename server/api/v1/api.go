package v1

import (
	"github.com/gin-gonic/gin"
	"hospital-system/server/api/v1/admin"
	session_handler "hospital-system/server/api/v1/session"
	session_service "hospital-system/server/app/services/session"
	"net/http"
)

const (
	AuthorizationCookieName = "Authorization"

	Authorization_MissingHeaderErr        = "missing_authorization_header"
	Authorization_MissingTokenErr         = "missing_token"
	Authorization_FailedToFetchSessionErr = "failed_to_fetch_session"
	Authorization_SessionExpiredErr       = "session_expired"
	Authorization_UserForbiddenError      = "user_not_allowed_to_execute_operation"
	Authorization_RefreshSessionErr       = "failed_to_refresh_session"
)

type API struct {
	adminHandler   admin.Handler
	sessionHandler session_handler.Handler
	sessionService session_service.Service
}

func NewAPI(adminHandler admin.Handler, sessionHandler session_handler.Handler, sessionService session_service.Service) *API {
	return &API{
		adminHandler:   adminHandler,
		sessionHandler: sessionHandler,
		sessionService: sessionService,
	}
}

func (api *API) RegisterRoutes(router *gin.Engine) {
	api.sessionHandler.RegisterRoutes(router)
	api.adminHandler.RegisterRoutes(router, api.authMiddleware)
}

func (api *API) authMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetch the Authorization header
		tokenString := c.GetHeader(AuthorizationCookieName)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": Authorization_MissingTokenErr})
			c.Abort()
			return
		}

		// Fetch the session from Redis store
		// If the session has expired, it will return nil for claims value
		claims, err := api.sessionService.GetSession(tokenString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": Authorization_FailedToFetchSessionErr})
			c.Abort()
			return
		}
		if claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": Authorization_SessionExpiredErr})
			c.Abort()
			return
		}

		// If user is not authorized to access certain resource, return status forbidden
		// TODO: Handle role-based access control
		//if claims.Role != role {
		//	c.JSON(http.StatusForbidden, gin.H{"error": Authorization_UserForbiddenError})
		//	c.Abort()
		//	return
		//}

		// Refresh the session
		if err = api.sessionService.RefreshSession(tokenString); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": Authorization_RefreshSessionErr})
			c.Abort()
			return
		}

		// If everything goes well, allow user to execute the operation
		c.Next()
	}
}
