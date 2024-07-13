package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"hospital-system/server/app/services/department"
	"hospital-system/server/app/services/patient"
	"hospital-system/server/app/services/user"
	"net/http"
)

type HandlerImpl struct {
	userService       user.Service
	patientService    patient.Service
	departmentService department.Service
}

func NewHandler(userService user.Service, patientService patient.Service, department department.Service) *HandlerImpl {
	return &HandlerImpl{
		userService:       userService,
		patientService:    patientService,
		departmentService: department,
	}
}

func (handler *HandlerImpl) AddUser(c *gin.Context) {
	var addUserRequest AddUserRequest
	if err := c.BindJSON(&addUserRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := handler.userService.CreateUser(user.User{
		Firstname:                    addUserRequest.Firstname,
		Lastname:                     addUserRequest.Lastname,
		NationalIdentificationNumber: addUserRequest.NationalIdentificationNumber,
		Email:                        addUserRequest.Email,
		Role:                         addUserRequest.Role,
		Team:                         addUserRequest.Team,
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if userId == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "user id is nil"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": AddUserResponse{Id: *userId},
	})
}

func (handler *HandlerImpl) GetUsers(c *gin.Context) {
	users, err := handler.userService.GetUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	users = lo.Filter(users, func(item user.User, _ int) bool {
		return item.Role != "ADMIN"
	})

	c.IndentedJSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (handler *HandlerImpl) AddPatient(c *gin.Context) {
	var addPatientRequest AddPatientRequest
	if err := c.BindJSON(&addPatientRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPatient, err := handler.patientService.CreatePatient(patient.Patient{
		Firstname:                    addPatientRequest.Firstname,
		Lastname:                     addPatientRequest.Lastname,
		NationalIdentificationNumber: addPatientRequest.NationalIdentificationNumber,
		MedicalRecordNumber:          addPatientRequest.MedicalRecordNumber,
		Email:                        addPatientRequest.Email,
		PhoneNumber:                  addPatientRequest.PhoneNumber,
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if newPatient == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "patient id is nil"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"patient": newPatient,
	})
}

func (handler *HandlerImpl) GetPatient(c *gin.Context) {
	patientId := c.Param("id")
	if patientId == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "patient id is required"})
		return
	}

	patient, err := handler.patientService.GetPatient(patientId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"patient": patient,
	})
}

func (handler *HandlerImpl) AdmitPatient(c *gin.Context) {
	var admitPatientRequest AdmitPatientRequest
	if err := c.BindJSON(&admitPatientRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := handler.patientService.RegisterPatientAdmission(admitPatientRequest.PatientId, patient.AdmissionDetails{
		Anamnesis: patient.Anamnesis{
			ChiefComplaint:          admitPatientRequest.ChiefComplaint,
			HistoryOfPresentIllness: admitPatientRequest.HistoryOfPresentIllness,
			PastMedicalHistory:      admitPatientRequest.PastMedicalHistory,
			Medications:             admitPatientRequest.Medications,
			Allergies:               admitPatientRequest.Allergies,
			FamilyHistory:           admitPatientRequest.FamilyHistory,
			SocialHistory:           admitPatientRequest.SocialHistory,
			PhysicalExamination:     admitPatientRequest.PhysicalExamination,
		},
		Department: admitPatientRequest.Department,
		Physician:  admitPatientRequest.Physician,
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{})
}

func (handler *HandlerImpl) GetActiveAdmissions(c *gin.Context) {
	admissions, err := handler.patientService.GetAdmissionsByStatuses([]string{"pending", "admitted"})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := GetAdmissionsResponse{
		Admissions: make([]AdmissionResponse, 0),
	}

	for _, admission := range admissions {
		patientName, err := handler.patientService.GetPatientName(admission.Patient.ID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		physician, err := handler.userService.GetUser(admission.Physician)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response.Admissions = append(response.Admissions, AdmissionResponse{
			Id:            admission.ID,
			Patient:       patientName,
			Department:    admission.Department,
			Physician:     fmt.Sprintf("Doctor %s, %s, MD", physician.Lastname, physician.Firstname),
			AdmissionTime: admission.StartTime,
			Status:        admission.Status,
		})
	}

	c.IndentedJSON(http.StatusOK, response)
}

// TODO: Fix this flow
func (handler *HandlerImpl) GetAdmission(c *gin.Context) {
	admissionId := c.Param("id")
	if admissionId == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "admission id is required"})
		return
	}

	admission, err := handler.patientService.GetAdmission(uuid.MustParse(admissionId))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	patientName, err := handler.patientService.GetPatientName(admission.Patient.ID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	physician, err := handler.userService.GetUser(admission.Physician)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"admission": AdmissionResponse{
			Id:            admission.ID,
			Patient:       patientName,
			Department:    admission.Department,
			Physician:     fmt.Sprintf("Doctor %s, %s, MD", physician.Lastname, physician.Firstname),
			AdmissionTime: admission.StartTime,
			Status:        admission.Status,
		},
	})

}

func (handler *HandlerImpl) GetActiveAdmissionsByPhysician(c *gin.Context) {
	physicianId := c.Param("id")
	if physicianId == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "physician id is required"})
		return
	}

	admissions, err := handler.patientService.GetActiveAdmissionsByPhysician(uuid.MustParse(physicianId))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := GetAdmissionsResponse{
		Admissions: make([]AdmissionResponse, 0),
	}

	for _, admission := range admissions {
		patientName, err := handler.patientService.GetPatientName(admission.Patient.ID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response.Admissions = append(response.Admissions, AdmissionResponse{
			Id:            admission.ID,
			Patient:       patientName,
			Department:    admission.Department,
			AdmissionTime: admission.StartTime,
			Status:        admission.Status,
		})
	}

	c.IndentedJSON(http.StatusOK, response)

}

func (handler *HandlerImpl) GetDepartments(c *gin.Context) {
	departments, err := handler.departmentService.GetDepartments()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"departments": departments,
	})
}
