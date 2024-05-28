package actors

import (
	"github.com/gin-gonic/gin"
	"hospital-system/authorization/app/services/actor"
	"net/http"
)

type HandlerImpl struct {
	actorService actor.Service
}

func NewHandler(actorService actor.Service) *HandlerImpl {
	return &HandlerImpl{
		actorService: actorService,
	}
}

func (handler *HandlerImpl) AddActor(c *gin.Context) {
	var addActorRequest AddActorRequest
	if err := c.BindJSON(&addActorRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := handler.actorService.AddActor(actor.Actor{
		ActorID: addActorRequest.ActorID,
		Role:    addActorRequest.Role,
		Team:    addActorRequest.Team,
	}); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{})
}

func (handler *HandlerImpl) GetActor(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func (handler *HandlerImpl) GetAllActors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}
