package actors

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (handler *HandlerImpl) RegisterRoutes(router *gin.Engine) {
	router.POST(path("/actors"), handler.AddActor)
	router.GET(path("/actors/:id"), handler.GetActor)
	router.GET(path("/actors"), handler.GetAllActors)
}

func path(endpoint string) string {
	return fmt.Sprintf("api/v1/authorization/%s", endpoint)
}
