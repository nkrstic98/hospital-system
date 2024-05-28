package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const AdminRole = "ADMIN"

func (handler *HandlerImpl) RegisterRoutes(router *gin.Engine, middleware func(string) gin.HandlerFunc) {
	router.POST(path("users"), middleware(AdminRole), handler.AddUser)
	router.GET(path("users"), middleware(AdminRole), handler.GetUsers)
	router.POST(path("patients"), middleware(AdminRole), handler.AddPatient)
	router.GET(path("patients/:id"), middleware(AdminRole), handler.GetPatient)
	router.POST(path("patients/admission"), middleware(AdminRole), handler.AdmitPatient)
	router.POST(path("patients/admissions"), middleware(AdminRole), handler.GetAdmissions)
	router.PATCH(path("patients/admissions/:id/discharge"), middleware(AdminRole), handler.Discharge)
	router.GET(path("departments"), middleware(AdminRole), handler.GetDepartments)
	router.POST(path("roles"), middleware(AdminRole), handler.AddRole)
	router.POST(path("teams"), middleware(AdminRole), handler.AddTeam)
}

func path(endpoint string) string {
	return fmt.Sprintf("api/v1/admin/%s", endpoint)
}
