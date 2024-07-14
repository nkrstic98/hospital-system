package patient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"hospital-system/proto_gen/authorization/v1"
	"hospital-system/server/app/dto"
	"hospital-system/server/app/repositories/admission"
	"hospital-system/server/app/repositories/patient"
	user_service "hospital-system/server/app/services/user"
	"hospital-system/server/models"
	"time"
)

type ServiceImpl struct {
	authorizationClient authorization.AuthorizationServiceClient
	patientRepo         patient.Repository
	admissionRepo       admission.Repository
	userService         user_service.Service
}

func NewService(authorizationClient authorization.AuthorizationServiceClient, patientRepo patient.Repository, admissionRepo admission.Repository, userService user_service.Service) *ServiceImpl {
	return &ServiceImpl{
		authorizationClient: authorizationClient,
		patientRepo:         patientRepo,
		admissionRepo:       admissionRepo,
		userService:         userService,
	}
}

func (service *ServiceImpl) RegisterPatient(patient dto.Patient) (*dto.Patient, error) {
	id := uuid.New()

	patientResponse, err := service.patientRepo.Insert(models.Patient{
		ID:                           id,
		Firstname:                    patient.Firstname,
		Lastname:                     patient.Lastname,
		NationalIdentificationNumber: patient.NationalIdentificationNumber,
		MedicalRecordNumber:          patient.MedicalRecordNumber,
		Email:                        patient.Email,
		PhoneNumber:                  patient.PhoneNumber,
	})
	if err != nil {
		return nil, err
	}

	return &dto.Patient{
		ID:                           patientResponse.ID,
		Firstname:                    patientResponse.Firstname,
		Lastname:                     patientResponse.Lastname,
		NationalIdentificationNumber: patientResponse.NationalIdentificationNumber,
		MedicalRecordNumber:          patientResponse.MedicalRecordNumber,
		Email:                        patientResponse.Email,
		PhoneNumber:                  patientResponse.PhoneNumber,
		Admissions:                   nil,
	}, nil
}

func (service *ServiceImpl) GetPatient(id string) (*dto.Patient, error) {
	patient, err := service.patientRepo.GetByPersonalID(id)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, nil
	}

	admissions, err := service.admissionRepo.GetByPatientId(patient.ID)
	if err != nil {
		return nil, err
	}

	return &dto.Patient{
		ID:                           patient.ID,
		Firstname:                    patient.Firstname,
		Lastname:                     patient.Lastname,
		NationalIdentificationNumber: patient.NationalIdentificationNumber,
		MedicalRecordNumber:          patient.MedicalRecordNumber,
		Email:                        patient.Email,
		PhoneNumber:                  patient.PhoneNumber,
		Admissions: lo.Map(admissions, func(admission models.Admission, _ int) dto.Admission {
			return dto.Admission{
				ID:        admission.ID,
				StartTime: admission.StartTime,
				EndTime:   admission.EndTime,
				Status:    admission.Status,
			}
		}),
	}, nil
}

func (service *ServiceImpl) GetPatientName(id uuid.UUID) (string, error) {
	patient, err := service.patientRepo.Get(id)
	if err != nil {
		return "", err

	}

	return fmt.Sprintf("%s, %s", patient.Lastname, patient.Firstname), nil
}

func (service *ServiceImpl) RegisterPatientAdmission(patientId uuid.UUID, admission dto.AdmissionDetails) error {
	marshalledAnamnesis, stdErr := json.Marshal(admission.Anamnesis)
	if stdErr != nil {
		slog.Error("Failed to marshal intake info")
		return stdErr
	}

	result, err := service.admissionRepo.Insert(models.Admission{
		ID:        uuid.New(),
		StartTime: time.Now(),
		Anamnesis: marshalledAnamnesis,
		PatientID: patientId,
	})
	if err != nil {
		return err
	}

	_, err = service.authorizationClient.AddResource(context.Background(), &authorization.AddResourceRequest{
		Id:       result.ID.String(),
		Team:     admission.Department,
		TeamLead: admission.Physician.String(),
	})
	if err != nil {
		slog.Error("Failed to add resource to team", err)
		deleteErr := service.admissionRepo.Delete(result.ID)
		if deleteErr != nil {
			slog.Error("Failed to delete admission with id: ", admission.ID, err)
			return err
		}

		return err
	}

	return nil
}

func (service *ServiceImpl) GetAdmissionsByStatuses(statuses []string) ([]dto.AdmissionDetails, error) {
	admissions, err := service.admissionRepo.GetByStatuses(statuses)
	if err != nil {
		return nil, err
	}

	resources, err := service.authorizationClient.GetResources(context.Background(), &authorization.GetResourcesRequest{
		Ids: lo.Map(admissions, func(admission models.Admission, _ int) string {
			return admission.ID.String()
		}),
	})
	if err != nil {
		slog.Error("Failed to get resources", err)
		return nil, err
	}

	return lo.Map(admissions, func(a models.Admission, _ int) dto.AdmissionDetails {
		patient, err := service.patientRepo.Get(a.PatientID)
		if err != nil {
			return dto.AdmissionDetails{}
		}

		return dto.AdmissionDetails{
			ID:        a.ID,
			StartTime: a.StartTime,
			Status:    a.Status,
			Patient: dto.Patient{
				ID:                           patient.ID,
				Firstname:                    patient.Firstname,
				Lastname:                     patient.Lastname,
				NationalIdentificationNumber: patient.NationalIdentificationNumber,
				MedicalRecordNumber:          patient.MedicalRecordNumber,
				Email:                        patient.Email,
				PhoneNumber:                  patient.PhoneNumber,
			},
			Department: lo.FindOrElse(resources.GetResources(), &authorization.Resource{}, func(item *authorization.Resource) bool {
				return item.Id == a.ID.String()
			}).Team.DisplayName,
			Physician: uuid.MustParse(lo.FindOrElse(resources.GetResources(), &authorization.Resource{}, func(item *authorization.Resource) bool {
				return item.Id == a.ID.String()
			}).TeamLead),
		}
	}), nil
}

func (service *ServiceImpl) GetAdmission(id uuid.UUID) (*dto.AdmissionDetails, error) {
	admission, err := service.admissionRepo.Get(id)
	if err != nil {
		return nil, err
	}
	if admission == nil {
		return nil, nil
	}

	resources, err := service.authorizationClient.GetResources(context.Background(), &authorization.GetResourcesRequest{
		Ids: []string{admission.ID.String()},
	})
	if err != nil {
		slog.Error("Failed to get resources", err)
		return nil, err
	}

	patient, err := service.patientRepo.Get(admission.PatientID)
	if err != nil {
		return nil, err
	}

	return &dto.AdmissionDetails{
		ID:        admission.ID,
		StartTime: admission.StartTime,
		Status:    admission.Status,
		Patient: dto.Patient{
			ID: patient.ID,
		},
		Department: lo.FindOrElse(resources.GetResources(), &authorization.Resource{}, func(item *authorization.Resource) bool {
			return item.Id == admission.ID.String()
		}).Team.DisplayName,
		Physician: uuid.MustParse(lo.FindOrElse(resources.GetResources(), &authorization.Resource{}, func(item *authorization.Resource) bool {
			return item.Id == admission.ID.String()
		}).TeamLead),
	}, nil
}

func (service *ServiceImpl) GetActiveAdmissionsByPhysician(physicianId uuid.UUID) ([]dto.AdmissionDetails, error) {
	resources, err := service.authorizationClient.GetActorResources(context.Background(), &authorization.GetActorResourcesRequest{
		ActorId:  physicianId.String(),
		Archived: false,
	})
	if err != nil {
		return nil, err
	}

	admissions, err := service.admissionRepo.GetByIDs(lo.Map(resources.GetResources(), func(resource *authorization.Resource, _ int) uuid.UUID {
		return uuid.MustParse(resource.Id)
	}))
	if err != nil {
		return nil, err
	}

	resultAdmissions := make([]dto.AdmissionDetails, len(admissions))
	for _, a := range admissions {
		var anamnesis dto.Anamnesis
		if err := json.Unmarshal(a.Anamnesis, &anamnesis); err != nil {
			slog.Error("Failed to unmarshal anamnesis", err)
			return nil, err
		}

		var vitals dto.Vitals
		if err := json.Unmarshal(a.Vitals, &vitals); err != nil {
			slog.Error("Failed to unmarshal vitals", err)
			return nil, err
		}

		var medications *[]dto.MedicationInfo
		if err := json.Unmarshal(a.Medications, &medications); err != nil {
			slog.Error("Failed to unmarshal medications", err)
			return nil, err
		}

		var logs []dto.Log
		if err := json.Unmarshal(a.Logs, &logs); err != nil {
			slog.Error("Failed to unmarshal logs", err)
			return nil, err
		}

		labs, err := service.admissionRepo.GetLabsByAdmissionID(a.ID)
		if err != nil {
			slog.Error("Failed to get labs", err)
			return nil, err
		}

		labResults := make([]dto.Lab, len(labs))
		for _, lab := range labs {
			var test *[]dto.LabTest
			if err := json.Unmarshal(lab.TestResults, &test); err != nil {
				slog.Error("Failed to unmarshal lab test results", err)
				return nil, err
			}

			labResults = append(labResults, dto.Lab{
				ID:          lab.ID,
				RequestedAt: lab.RequestedAt,
				ProcessedAt: lab.ProcessedAt,
				TestType:    lab.TestType,
				TestResults: test,
			})
		}

		patient, err := service.patientRepo.Get(a.PatientID)
		if err != nil {
			return nil, err
		}

		resultAdmissions = append(resultAdmissions, dto.AdmissionDetails{
			ID:          a.ID,
			StartTime:   a.StartTime,
			EndTime:     a.EndTime,
			Status:      a.Status,
			Anamnesis:   anamnesis,
			Vitals:      vitals,
			Diagnosis:   a.Diagnosis,
			Medications: medications,
			Labs:        &labResults,
			Logs:        logs,
			Patient: dto.Patient{
				ID:                           patient.ID,
				Firstname:                    patient.Firstname,
				Lastname:                     patient.Lastname,
				NationalIdentificationNumber: patient.NationalIdentificationNumber,
				MedicalRecordNumber:          patient.MedicalRecordNumber,
				Email:                        patient.Email,
				PhoneNumber:                  patient.PhoneNumber,
			},
			Department: lo.FindOrElse(resources.GetResources(), &authorization.Resource{}, func(item *authorization.Resource) bool {
				return item.Id == a.ID.String()
			}).Team.DisplayName,
			Physician: uuid.MustParse(lo.FindOrElse(resources.GetResources(), &authorization.Resource{}, func(item *authorization.Resource) bool {
				return item.Id == a.ID.String()
			}).TeamLead),
		})
	}

	return resultAdmissions, nil
}

func (service *ServiceImpl) GetAdmissions() ([]dto.Admission, error) {
	admissions, err := service.admissionRepo.GetByStatuses([]string{models.StatusPending, models.StatusAdmitted})
	if err != nil {
		return nil, err
	}

	resources, err := service.authorizationClient.GetResources(context.Background(), &authorization.GetResourcesRequest{
		Ids: lo.Map(admissions, func(admission models.Admission, _ int) string {
			return admission.ID.String()
		}),
	})
	if err != nil {
		slog.Error("Failed to get resources", err)
		return nil, err
	}

	return lo.Map(admissions, func(a models.Admission, _ int) dto.Admission {
		patient, err := service.GetPatientName(a.PatientID)
		if err != nil {
			return dto.Admission{}
		}

		pid := uuid.MustParse(lo.FindOrElse(resources.GetResources(), &authorization.Resource{}, func(item *authorization.Resource) bool {
			return item.Id == a.ID.String()
		}).TeamLead)
		physician, err := service.userService.GetUser(pid)
		if err != nil {
			return dto.Admission{}
		}

		return dto.Admission{
			ID:        a.ID,
			StartTime: a.StartTime,
			Status:    a.Status,
			Patient:   patient,
			Department: lo.FindOrElse(resources.GetResources(), &authorization.Resource{}, func(item *authorization.Resource) bool {
				return item.Id == a.ID.String()
			}).Team.DisplayName,
			Physician: fmt.Sprintf("Doctor %s, %s, MD", physician.Lastname, physician.Firstname),
		}
	}), nil
}
