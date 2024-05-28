package session

import "github.com/gin-gonic/gin"

type Handler interface {
	RegisterRoutes(router *gin.Engine)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	ValidateSession(c *gin.Context)
}
