package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (handler *HandlerImpl) RegisterRoutes(router *gin.Engine) {
	router.POST(path("login"), handler.Login)
	router.POST(path("logout"), handler.Logout)
	router.POST(path("validate"), handler.ValidateSession)
}

func path(endpoint string) string {
	return fmt.Sprintf("api/v1/session/%s", endpoint)
}
