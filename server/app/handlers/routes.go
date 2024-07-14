package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.POST(path("sessions/login"), h.login)
	router.POST(path("sessions/logout"), h.logout)
	router.POST(path("sessions/validate"), h.validateSession)

	router.POST(path("users"), h.rbacAuthMiddleware("EMPLOYEES", "WRITE"), h.registerUser)
	router.GET(path("users"), h.rbacAuthMiddleware("EMPLOYEES", "%"), h.getUsers)

	router.POST(path("patients"), h.rbacAuthMiddleware("INTAKE", "WRITE"), h.registerPatient)
	router.GET(path("patients/:id"), h.rbacAuthMiddleware("PATIENTS", "READ"), h.getPatient)

	router.POST(path("patients/admissions/register"), h.rbacAuthMiddleware("INTAKE", "WRITE"), h.admitPatient)
	router.GET(path("patients/admissions"), h.rbacAuthMiddleware("INTAKE", "READ"), h.getAdmissions)

	router.GET(path("departments"), h.basicAuthMiddleware(), h.getDepartments)
}

func path(endpoint string) string {
	return fmt.Sprintf("api/v1/%s", endpoint)
}
