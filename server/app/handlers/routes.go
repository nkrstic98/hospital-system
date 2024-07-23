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
	router.POST(path("users/departments"), h.basicAuthMiddleware(), h.getDepartments)

	router.POST(path("patients"), h.rbacAuthMiddleware("INTAKE", "WRITE"), h.registerPatient)
	router.GET(path("patients/:id"), h.rbacAuthMiddleware("PATIENTS", "READ"), h.getPatient)

	router.POST(path("patients/admissions"), h.rbacAuthMiddleware("INTAKE", "WRITE"), h.admitPatient)
	router.GET(path("patients/admissions"), h.rbacAuthMiddleware("INTAKE", "READ"), h.getActiveAdmissions)
	router.PUT(path("patients/admissions"), h.rbacAuthMiddleware("ADMISSIONS", "%"), h.updateAdmission)
	router.GET(path("patients/users/:id/admissions"), h.rbacAuthMiddleware("ADMISSIONS", "READ"), h.getUserActiveAdmissions)
	router.GET(path("patients/admissions/:id"), h.rbacAuthMiddleware("PATIENTS", "READ"), h.getAdmissionDetails)
	router.GET(path("patients/admissions/:id/transfer"), h.rbacAuthMiddleware("ADMISSIONS", "WRITE"), h.acceptTransferRequest)
	router.GET(path("patients/admissions/:id/admit"), h.rbacAuthMiddleware("ADMISSIONS", "WRITE"), h.acceptAdmissionRequest)
	router.POST(path("patients/admissions/add-team-member"), h.rbacAuthMiddleware("PATIENTS:TEAM", "WRITE"), h.addTeamMember)
	router.POST(path("patients/admissions/remove-team-member"), h.rbacAuthMiddleware("PATIENTS:TEAM", "WRITE"), h.removeTeamMember)
	router.POST(path("patients/admissions/add-team-member-permissions"), h.rbacAuthMiddleware("PATIENTS:TEAM", "WRITE"), h.addTeamMemberPermissions)
	router.POST(path("patients/admissions/remove-team-member-permissions"), h.rbacAuthMiddleware("PATIENTS:TEAM", "WRITE"), h.removeTeamMemberPermissions)
	router.POST(path("patients/admissions/request-transfer"), h.rbacAuthMiddleware("PATIENTS:TRANSFER", "WRITE"), h.requestPatientTransfer)
	router.PATCH(path("patients/admissions/:id/discharge"), h.rbacAuthMiddleware("PATIENTS:DISCHARGE", "WRITE"), h.dischargePatient)
	router.POST(path("patients/labs"), h.rbacAuthMiddleware("PATIENTS:LABS:ORDER", "WRITE"), h.orderLabTest)

	router.GET(path("labs"), h.rbacAuthMiddleware("LABS", "WRITE"), h.getLabs)
	router.POST(path("labs/:id/process"), h.rbacAuthMiddleware("LABS", "WRITE"), h.processLabTest)
}

func path(endpoint string) string {
	return fmt.Sprintf("api/v1/%s", endpoint)
}
