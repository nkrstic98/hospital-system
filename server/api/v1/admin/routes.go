package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const AdminRole = "ADMIN"

func (handler *HandlerImpl) RegisterRoutes(router *gin.Engine, middleware func(string) gin.HandlerFunc) {
	// Register user
	router.POST(path("users"), middleware(AdminRole), handler.AddUser)
	// Get users
	router.GET(path("users"), middleware(AdminRole), handler.GetUsers)

	// Get patient by id
	router.GET(path("patients/:id"), middleware(AdminRole), handler.GetPatient)
	// Register patient
	router.POST(path("patients"), middleware(AdminRole), handler.AddPatient)

	// Admit patient
	router.POST(path("patients/admissions/register"), middleware(AdminRole), handler.AdmitPatient)
	// Get all active admissions
	router.GET(path("patients/admissions"), middleware(AdminRole), handler.GetActiveAdmissions)
	// Get admission by id
	router.GET(path("patients/admissions/:id"), middleware(AdminRole), handler.GetAdmission)
	// Get active admissions by physician id
	router.POST(path("patients/admissions/physician/:id"), middleware(AdminRole), handler.GetActiveAdmissionsByPhysician)

	// Get departments
	router.GET(path("departments"), middleware(AdminRole), handler.GetDepartments)
}

func path(endpoint string) string {
	return fmt.Sprintf("api/v1/admin/%s", endpoint)
}
