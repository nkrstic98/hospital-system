package services

import (
	"context"
	"fmt"
	"hospital-system/proto_gen/authorization/v1"
	"hospital-system/server/app/dto"
	"hospital-system/server/models"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type patientRepo interface {
	InsertPatient(ctx context.Context, user models.Patient) (*models.Patient, error)
	GetPatient(ctx context.Context, id uuid.UUID) (*models.Patient, error)
	GetPatientByPersonalID(ctx context.Context, personalID string) (*models.Patient, error)

	InsertAdmission(ctx context.Context, admission models.Admission) (*models.Admission, error)
	GetAdmission(ctx context.Context, id uuid.UUID) (*models.Admission, error)
	DeleteAdmission(ctx context.Context, id uuid.UUID) error
	GetAdmissionsByPatientId(ctx context.Context, id uuid.UUID) ([]models.Admission, error)
	GetAdmissionsByStatuses(ctx context.Context, statuses []string) ([]models.Admission, error)
	GetAdmissionsByIDs(ctx context.Context, ids []uuid.UUID) ([]models.Admission, error)

	GetLabsByAdmissionID(ctx context.Context, admissionID uuid.UUID) ([]models.Lab, error)
}

type PatientService struct {
	authorizationClient authorization.AuthorizationServiceClient
	repo                patientRepo
	userService         *UserService
}

func NewPatientService(authorizationClient authorization.AuthorizationServiceClient, repo patientRepo, userService *UserService) *PatientService {
	return &PatientService{
		authorizationClient: authorizationClient,
		repo:                repo,
		userService:         userService,
	}
}

func (s *PatientService) RegisterPatient(ctx context.Context, patient dto.Patient) (*dto.Patient, error) {
	patientResponse, err := s.repo.InsertPatient(ctx, models.Patient{
		Firstname:                    patient.Firstname,
		Lastname:                     patient.Lastname,
		NationalIdentificationNumber: patient.NationalIdentificationNumber,
		MedicalRecordNumber:          patient.MedicalRecordNumber,
		Email:                        patient.Email,
		PhoneNumber:                  patient.PhoneNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert patient: %w", err)
	}
	if patientResponse == nil {
		return nil, fmt.Errorf("patient insert returned nil response")
	}

	dtoPatient := toDtoPatient(*patientResponse)

	return &dtoPatient, nil
}

func (s *PatientService) GetPatient(ctx context.Context, id string) (*dto.Patient, error) {
	patient, err := s.repo.GetPatientByPersonalID(ctx, id)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, fmt.Errorf("patient not found")
	}

	dtoPatient := toDtoPatient(*patient)

	return &dtoPatient, nil
}

func (s *PatientService) RegisterPatientAdmission(ctx context.Context, patientId uuid.UUID, admission dto.AdmissionDetails) error {
	// Add admission log
	admission.Logs = []dto.Log{
		{
			Timestamp:   time.Now(),
			Action:      "Patient Admission",
			Message:     "Patient admitted to the hospital",
			PerformedBy: admission.Physician,
		},
	}

	admissionModel, err := toAdmissionModel(patientId, admission)
	if err != nil {
		return fmt.Errorf("failed to convert admission dto to model: %w", err)
	}

	result, err := s.repo.InsertAdmission(ctx, admissionModel)
	if err != nil {
		return err
	}

	if _, err = s.authorizationClient.AddResource(context.Background(), &authorization.AddResourceRequest{
		Id:       result.ID.String(),
		Team:     admission.Department,
		TeamLead: admission.Physician.String(),
	}); err != nil {
		deleteErr := s.repo.DeleteAdmission(ctx, result.ID)
		if deleteErr != nil {
			return fmt.Errorf("failed to delete admission: %w", deleteErr)
		}

		return fmt.Errorf("failed to add resource: %w", err)
	}

	return nil
}

func (s *PatientService) GetActiveAdmissions(ctx context.Context) ([]dto.Admission, error) {
	admissions, err := s.repo.GetAdmissionsByStatuses(ctx, []string{models.StatusPending, models.StatusAdmitted})
	if err != nil {
		return nil, fmt.Errorf("failed to get admissions: %w", err)
	}
	if len(admissions) == 0 {
		return nil, nil
	}

	resources, err := s.authorizationClient.GetResources(context.Background(), &authorization.GetResourcesRequest{
		Ids: lo.Map(admissions, func(admission models.Admission, _ int) string {
			return admission.ID.String()
		}),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get resources: %w", err)
	}

	resultAdmissions := make([]dto.Admission, 0, len(admissions))

	for _, a := range admissions {
		patient, err := s.repo.GetPatient(ctx, a.PatientID)
		if err != nil {
			return nil, fmt.Errorf("failed to get patient: %w", err)
		}
		if patient == nil {
			return nil, fmt.Errorf("patient with id %v not found", a.PatientID)
		}

		physicianId := uuid.MustParse(lo.FindOrElse(resources.GetResources(), &authorization.Resource{}, func(item *authorization.Resource) bool {
			return item.Id == a.ID.String()
		}).TeamLead)

		physician, err := s.userService.GetUser(ctx, physicianId)
		if err != nil {
			return nil, fmt.Errorf("failed to get physician: %w", err)
		}
		if physician == nil {
			return nil, fmt.Errorf("physician with id %v not found", physicianId)
		}

		resultAdmissions = append(resultAdmissions, dto.Admission{
			ID:        a.ID,
			StartTime: a.StartTime,
			EndTime:   a.EndTime,
			Status:    a.Status,
			Patient:   fmt.Sprintf("%s, %s", patient.Lastname, patient.Firstname),
			Department: lo.FindOrElse(resources.GetResources(), &authorization.Resource{}, func(item *authorization.Resource) bool {
				return item.Id == a.ID.String()
			}).Team.DisplayName,
			Physician: fmt.Sprintf("Doctor %s, %s, MD", physician.Lastname, physician.Firstname),
		})
	}

	return resultAdmissions, nil
}
