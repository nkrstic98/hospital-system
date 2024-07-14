package handlers

import (
	"hospital-system/server/app/services"

	"go.uber.org/zap"
)

const AuthorizationCookieName = "Authorization"

type Handler struct {
	log            *zap.Logger
	userService    *services.UserService
	sessionService *services.SessionService
	patientService *services.PatientService
}

func NewHandler(
	log *zap.Logger,
	userService *services.UserService,
	sessionService *services.SessionService,
	patientService *services.PatientService,
) *Handler {
	return &Handler{
		log:            log,
		userService:    userService,
		sessionService: sessionService,
		patientService: patientService,
	}
}
