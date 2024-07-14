package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) getDepartments(c *gin.Context) {
	result, err := h.departmentService.GetDepartments()
	if err != nil {
		h.log.Error("Failed to get departments", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"departments": result,
	})
}
