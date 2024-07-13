package admin

import "github.com/gin-gonic/gin"

type Handler interface {
	RegisterRoutes(router *gin.Engine, middleware func(string) gin.HandlerFunc)
	AddUser(c *gin.Context)
	GetUsers(c *gin.Context)
	AddPatient(c *gin.Context)
	GetPatient(c *gin.Context)
	AdmitPatient(c *gin.Context)
	GetActiveAdmissions(c *gin.Context)
	GetAdmission(c *gin.Context)
	GetActiveAdmissionsByPhysician(c *gin.Context)
	GetDepartments(c *gin.Context)
}
