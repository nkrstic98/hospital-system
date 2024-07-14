package handlers

import (
	"go.uber.org/zap"
	"hospital-system/server/app/services/department"
	"hospital-system/server/app/services/patient"
	"hospital-system/server/app/services/session"
	"hospital-system/server/app/services/user"
)

const AuthorizationCookieName = "Authorization"

type Handler struct {
	log               *zap.Logger
	userService       user.Service
	sessionService    session.Service
	patientService    patient.Service
	departmentService department.Service
}

func NewHandler(
	log *zap.Logger,
	userService user.Service,
	sessionService session.Service,
	patientService patient.Service,
	departmentService department.Service,
) *Handler {
	return &Handler{
		log:               log,
		userService:       userService,
		sessionService:    sessionService,
		patientService:    patientService,
		departmentService: departmentService,
	}
}
