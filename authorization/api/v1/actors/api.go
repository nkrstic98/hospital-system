package actors

import "github.com/gin-gonic/gin"

type Handler interface {
	RegisterRoutes(router *gin.Engine)
	AddActor(c *gin.Context)
	GetActor(c *gin.Context)
	GetAllActors(c *gin.Context)
}
