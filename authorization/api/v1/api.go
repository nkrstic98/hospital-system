package v1

import (
	"github.com/gin-gonic/gin"
	"hospital-system/authorization/api/v1/actors"
)

type API struct {
	actorsHandler actors.Handler
}

func NewAPI(actorsHandler actors.Handler) *API {
	return &API{
		actorsHandler: actorsHandler,
	}
}

func (api *API) RegisterRoutes(router *gin.Engine) {
	api.actorsHandler.RegisterRoutes(router)
}
