package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hospital-system/server/app/dto"
	"net/http"
)

const (
	Patients_Get_MissingPatientID = "missing_patient_id"
	Patients_Get_PatientNotFound  = "patient_not_found"
)

func (h *Handler) registerPatient(c *gin.Context) {
	var request dto.RegisterPatientRequest
	if err := c.BindJSON(&request); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result, err := h.patientService.RegisterPatient(dto.Patient{
		Firstname:                    request.Firstname,
		Lastname:                     request.Lastname,
		NationalIdentificationNumber: request.NationalIdentificationNumber,
		MedicalRecordNumber:          request.MedicalRecordNumber,
		Email:                        request.Email,
		PhoneNumber:                  request.PhoneNumber,
	})
	if err != nil {
		h.log.Error("Failed to create patient", zap.String("nid", request.NationalIdentificationNumber), zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func (h *Handler) getPatient(c *gin.Context) {
	patientId := c.Param("id")
	if patientId == "" {
		h.log.Warn("Get patient called without providing patient id")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": Patients_Get_MissingPatientID})
		return
	}

	result, err := h.patientService.GetPatient(patientId)
	if err != nil {
		h.log.Error("Failed to get patient", zap.String("id", patientId), zap.Error(err))
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	if result == nil {
		h.log.Warn("Patient not found", zap.String("id", patientId))
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": Patients_Get_PatientNotFound})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func (h *Handler) admitPatient(c *gin.Context) {
	var request dto.AdmitPatientRequest
	if err := c.BindJSON(&request); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.patientService.RegisterPatientAdmission(request.PatientId, dto.AdmissionDetails{
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
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) getAdmissions(c *gin.Context) {
	admissions, err := h.patientService.GetAdmissions()
	if err != nil {
		h.log.Error("Failed to get admissions", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, admissions)
}
