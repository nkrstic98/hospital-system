package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) getLabs(ctx *gin.Context) {
	labs, err := h.labService.GetLabs(ctx)
	if err != nil {
		h.log.Error("failed to get labs", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get labs"})
		return
	}

	ctx.JSON(http.StatusOK, labs)
}

func (h *Handler) processLabTest(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		h.log.Error("User id not provided in context")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user id not provided in context"})
		return
	}

	labId := ctx.Param("id")
	if labId == "" {
		h.log.Error("Lab id not provided")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "lab id not provided"})
		return
	}

	if err := h.labService.ProcessLabTest(ctx, uuid.MustParse(labId), userId.(uuid.UUID)); err != nil {
		h.log.Error("failed to process lab test", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process lab test"})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
