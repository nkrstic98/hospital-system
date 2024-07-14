package handlers

import (
	"hospital-system/server/app/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	Patients_Get_MissingPatientID = "missing_patient_id"
	Patients_Get_PatientNotFound  = "patient_not_found"
)

func (h *Handler) registerPatient(ctx *gin.Context) {
	var request dto.RegisterPatientRequest
	if err := ctx.BindJSON(&request); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result, err := h.patientService.RegisterPatient(ctx, dto.Patient{
		Firstname:                    request.Firstname,
		Lastname:                     request.Lastname,
		NationalIdentificationNumber: request.NationalIdentificationNumber,
		MedicalRecordNumber:          request.MedicalRecordNumber,
		Email:                        request.Email,
		PhoneNumber:                  request.PhoneNumber,
	})
	if err != nil {
		h.log.Error("Failed to create patient", zap.String("nid", request.NationalIdentificationNumber), zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

func (h *Handler) getPatient(ctx *gin.Context) {
	patientId := ctx.Param("id")
	if patientId == "" {
		h.log.Warn("Get patient called without providing patient id")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": Patients_Get_MissingPatientID})
		return
	}

	result, err := h.patientService.GetPatient(ctx, patientId)
	if err != nil {
		h.log.Error("Failed to get patient", zap.String("id", patientId), zap.Error(err))
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	if result == nil {
		h.log.Warn("Patient not found", zap.String("id", patientId))
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": Patients_Get_PatientNotFound})
		return
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

func (h *Handler) admitPatient(ctx *gin.Context) {
	var request dto.AdmitPatientRequest
	if err := ctx.BindJSON(&request); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.patientService.RegisterPatientAdmission(ctx, request.PatientId, dto.AdmissionDetails{
		Anamnesis: dto.Anamnesis{
			ChiefComplaint:          request.ChiefComplaint,
			HistoryOfPresentIllness: request.HistoryOfPresentIllness,
			PastMedicalHistory:      request.PastMedicalHistory,
			Medications:             request.Medications,
			Allergies:               request.Allergies,
			FamilyHistory:           request.FamilyHistory,
			SocialHistory:           request.SocialHistory,
			PhysicalExamination:     request.PhysicalExamination,
		},
		Department: request.Department,
		Physician:  request.Physician,
	}); err != nil {
		h.log.Error("Failed to admit patient", zap.Any("patient_id", request.PatientId), zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) getActiveAdmissions(ctx *gin.Context) {
	admissions, err := h.patientService.GetActiveAdmissions(ctx)
	if err != nil {
		h.log.Error("Failed to get admissions", zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, admissions)
}
