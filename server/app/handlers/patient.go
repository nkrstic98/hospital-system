package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"hospital-system/server/app/dto"
	"hospital-system/server/models"
	"net/http"
)

const (
	Patients_Get_MissingPatientID               = "missing_patient_id"
	Patients_Get_MissingUserID                  = "missing_user_id"
	Patients_Get_MissingAdmissionID             = "missing_admission_id"
	Patients_Get_PatientNotFound                = "patient_not_found"
	Authorization_MissingUserRole               = "missing_user_role"
	Admissions_Accept_Failed_MissingAdmissionID = "missing_admission_id"
	Admissions_Accept_Failed_MissingAcceptFlag  = "missing_accept_flag"
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

	userRole := ctx.GetString("userRole")
	if userRole == "" {
		h.log.Warn("userRole missing in context")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": Authorization_MissingUserRole})
		return
	}

	careTeam := dto.CareTeam{
		Department: request.AdmittedByTeam,
		TeamLead:   request.AdmittedBy,
	}
	status := models.StatusAdmitted
	if userRole != "ATTENDING" {
		careTeam.Department = request.Department
		careTeam.TeamLead = request.Physician
		status = models.StatusPending
	}

	var pendingTransfer *dto.JourneyStep
	if request.AdmittedByTeam != request.Department || request.AdmittedBy != request.Physician {
		pendingTransfer = &dto.JourneyStep{
			FromTeam:     request.AdmittedByTeam,
			ToTeam:       request.Department,
			FromTeamLead: request.AdmittedBy,
			ToTeamLead:   request.Physician,
		}
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
		CareTeam:        careTeam,
		PendingTransfer: pendingTransfer,
		Status:          status,
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

func (h *Handler) updateAdmission(ctx *gin.Context) {
	var request dto.AdmissionDetails
	if err := ctx.BindJSON(&request); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	response, err := h.patientService.UpdateAdmission(ctx, request)
	if err != nil {
		h.log.Error("Failed to update admission", zap.Any("admission_id", request.ID), zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}

func (h *Handler) getUserActiveAdmissions(ctx *gin.Context) {
	userId := ctx.Param("id")
	if userId == "" {
		h.log.Warn("Get active admissions by user id called without providing user id")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": Patients_Get_MissingUserID})
		return
	}

	admissions, err := h.patientService.GetActiveAdmissionsByUserId(ctx, userId)
	if err != nil {
		h.log.Error("Failed to get admissions by user id", zap.String("user_id", userId), zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, admissions)

}

func (h *Handler) getAdmissionDetails(ctx *gin.Context) {
	patientId := ctx.Param("id")
	if patientId == "" {
		h.log.Warn("Get admission details called without providing patient id")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": Patients_Get_MissingAdmissionID})
		return
	}

	result, err := h.patientService.GetAdmissionDetails(ctx, uuid.MustParse(patientId))
	if err != nil {
		h.log.Error("Failed to get admission details", zap.String("id", patientId), zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

func (h *Handler) acceptTransferRequest(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists || userId == "" {
		h.log.Warn("Get active admissions by user id called without providing user id")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": Patients_Get_MissingUserID})
		return
	}

	admissionId := ctx.Param("id")
	if admissionId == "" {
		h.log.Warn("Accept admission request called without providing admission id")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": Admissions_Accept_Failed_MissingAdmissionID})
		return
	}

	acceptResource := ctx.Query("accept")
	if acceptResource == "" {
		h.log.Warn("Accept admission request called without providing accept query")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": Admissions_Accept_Failed_MissingAcceptFlag})
		return
	}
	accept := lo.Ternary(acceptResource == "true", true, false)

	if err := h.patientService.AcceptTransferRequest(ctx, userId.(uuid.UUID), uuid.MustParse(admissionId), accept); err != nil {
		h.log.Error("Failed to accept transfer request", zap.String("id", admissionId), zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) acceptAdmissionRequest(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists || userId == "" {
		h.log.Warn("Get active admissions by user id called without providing user id")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": Patients_Get_MissingUserID})
		return
	}

	admissionId := ctx.Param("id")
	if admissionId == "" {
		h.log.Warn("Accept admission request called without providing admission id")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": Admissions_Accept_Failed_MissingAdmissionID})
		return
	}

	if err := h.patientService.AcceptAdmissionRequest(ctx, userId.(uuid.UUID), uuid.MustParse(admissionId)); err != nil {
		h.log.Error("Failed to accept admission request", zap.String("id", admissionId), zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) orderLabTest(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists || userId == "" {
		h.log.Warn("Get active admissions by user id called without providing user id")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": Patients_Get_MissingUserID})
		return
	}

	var orderLabTest dto.OrderLabTestRequest
	if err := ctx.BindJSON(&orderLabTest); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.patientService.OrderLabTest(ctx, userId.(uuid.UUID), orderLabTest); err != nil {
		h.log.Error("Failed to order lab test", zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) addTeamMember(ctx *gin.Context) {
	var req dto.AddTeamMemberRequest
	if err := ctx.BindJSON(&req); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.patientService.AddTeamMember(ctx, req.UserId, req.AdmissionId); err != nil {
		h.log.Error("Failed to add team member", zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) removeTeamMember(ctx *gin.Context) {
	var req dto.RemoveTeamMemberRequest
	if err := ctx.BindJSON(&req); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.patientService.RemoveTeamMember(ctx, req.UserId, req.AdmissionId); err != nil {
		h.log.Error("Failed to remove team member", zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) addTeamMemberPermissions(ctx *gin.Context) {
	var req dto.AddTeamMemberPermissionsRequest
	if err := ctx.BindJSON(&req); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.patientService.AddTeamMemberPermissions(ctx, req.UserId, req.AdmissionId, req.Section, req.Permission); err != nil {
		h.log.Error("Failed to add team member permissions", zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) removeTeamMemberPermissions(ctx *gin.Context) {
	var req dto.RemoveTeamMemberPermissionsRequest
	if err := ctx.BindJSON(&req); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.patientService.RemoveTeamMemberPermissions(ctx, req.UserId, req.AdmissionId, req.Section); err != nil {
		h.log.Error("Failed to remove team member permissions", zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) requestPatientTransfer(ctx *gin.Context) {
	var req dto.TransferPatientRequest
	if err := ctx.BindJSON(&req); err != nil {
		h.log.Error("Failed to bind json", zap.Error(err))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.patientService.RequestPatientTransfer(ctx, req.AdmissionId, req.ToTeam, req.ToTeamLead); err != nil {
		h.log.Error("Failed to request patient transfer", zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) dischargePatient(ctx *gin.Context) {
	admissionID := ctx.Param("id")
	if admissionID == "" {
		h.log.Warn("Discharge patient called without providing admission id")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": Patients_Get_MissingAdmissionID})
		return
	}

	if err := h.patientService.DischargePatient(ctx, uuid.MustParse(admissionID)); err != nil {
		h.log.Error("Failed to discharge patient", zap.Error(err))
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, nil)
}
