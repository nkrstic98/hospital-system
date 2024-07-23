package services

import (
	"context"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
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
	UpdateAdmission(ctx context.Context, admission *models.Admission) error
	GetAdmission(ctx context.Context, id uuid.UUID) (*models.Admission, error)
	DeleteAdmission(ctx context.Context, id uuid.UUID) error
	GetAdmissionsByPatientId(ctx context.Context, id uuid.UUID) ([]models.Admission, error)
	GetAdmissionsByStatuses(ctx context.Context, statuses []string) ([]models.Admission, error)
	GetAdmissionsByIDs(ctx context.Context, ids []uuid.UUID) ([]models.Admission, error)

	GetLabsByAdmissionID(ctx context.Context, admissionID uuid.UUID) ([]models.Lab, error)

	InsertLab(ctx context.Context, lab models.Lab) error
}

type PatientService struct {
	log                 *zap.Logger
	authorizationClient authorization.AuthorizationServiceClient
	repo                patientRepo
	userService         *UserService
}

func NewPatientService(log *zap.Logger, authorizationClient authorization.AuthorizationServiceClient, repo patientRepo, userService *UserService) *PatientService {
	return &PatientService{
		log:                 log,
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

	dtoPatient := toDtoPatient(*patientResponse, nil)

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

	admissions, err := s.repo.GetAdmissionsByPatientId(ctx, patient.ID)
	if err != nil {
		return nil, err
	}

	dtoPatient := toDtoPatient(*patient, admissions)

	return &dtoPatient, nil
}

func (s *PatientService) RegisterPatientAdmission(ctx context.Context, patientId uuid.UUID, admission dto.AdmissionDetails) error {
	teamLead, err := s.userService.GetUser(ctx, admission.CareTeam.TeamLead)
	if err != nil {
		return fmt.Errorf("failed to get team lead: %w", err)
	}

	// Add admission log
	admission.Logs = []dto.Log{
		{
			Timestamp:   time.Now(),
			Action:      "Admit Patient",
			Message:     "Patient admitted to the hospital",
			PerformedBy: fmt.Sprintf("%s %s", teamLead.Firstname, teamLead.Lastname),
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

	var pendingTransfer *authorization.JourneyStep
	if admission.PendingTransfer != nil {
		pendingTransfer = &authorization.JourneyStep{
			FromTeam:     admission.PendingTransfer.FromTeam,
			ToTeam:       admission.PendingTransfer.ToTeam,
			FromTeamLead: admission.PendingTransfer.FromTeamLead.String(),
			ToTeamLead:   admission.PendingTransfer.ToTeamLead.String(),
		}
	}

	if _, err = s.authorizationClient.AddResource(ctx, &authorization.AddResourceRequest{
		Id:              result.ID.String(),
		Team:            admission.CareTeam.Department,
		TeamLead:        admission.CareTeam.TeamLead.String(),
		PendingTransfer: pendingTransfer,
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

	resources, err := s.authorizationClient.GetResources(ctx, &authorization.GetResourcesRequest{
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
			EndTime:   lo.Ternary(a.EndTime.Valid, &a.EndTime.Time, nil),
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

func (s *PatientService) GetActiveAdmissionsByUserId(ctx context.Context, userId string) ([]dto.Admission, error) {
	resources, err := s.authorizationClient.GetResources(ctx, &authorization.GetResourcesRequest{
		ActorId:  &userId,
		Archived: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get resources: %w", err)
	}

	admissions, err := s.repo.GetAdmissionsByIDs(ctx, lo.Map(resources.GetResources(), func(r *authorization.Resource, _ int) uuid.UUID {
		return uuid.MustParse(r.Id)
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to get admissions: %w", err)
	}

	resultAdmissions := make([]dto.Admission, 0, len(admissions))
	for _, a := range admissions {
		res, _ := lo.Find(resources.GetResources(), func(r *authorization.Resource) bool {
			return r.Id == a.ID.String()
		})

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

		status := a.Status
		if status == models.StatusAdmitted && res.PendingTransfer != nil && res.PendingTransfer.ToTeamLead == userId {
			status = models.StatusAwaitingTransfer
		}

		resultAdmissions = append(resultAdmissions, dto.Admission{
			ID:        a.ID,
			StartTime: a.StartTime,
			EndTime:   lo.Ternary(a.EndTime.Valid, &a.EndTime.Time, nil),
			Status:    status,
			Patient:   fmt.Sprintf("%s, %s", patient.Lastname, patient.Firstname),
			Department: lo.FindOrElse(resources.GetResources(), &authorization.Resource{}, func(item *authorization.Resource) bool {
				return item.Id == a.ID.String()
			}).Team.DisplayName,
			Physician: fmt.Sprintf("Doctor %s, %s, MD", physician.Lastname, physician.Firstname),
		})
	}

	return resultAdmissions, nil
}

func (s *PatientService) GetAdmissionDetails(ctx context.Context, id uuid.UUID) (*dto.AdmissionDetails, error) {
	admission, err := s.repo.GetAdmission(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get admission: %w", err)
	}
	if admission == nil {
		return nil, fmt.Errorf("admission with id %v not found", id)
	}

	labs, err := s.repo.GetLabsByAdmissionID(ctx, admission.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get labs: %w", err)
	}

	patient, err := s.repo.GetPatient(ctx, admission.PatientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get patient: %w", err)
	}
	if patient == nil {
		return nil, fmt.Errorf("patient with id %v not found", admission.PatientID)
	}

	admissionDetails, err := toDtoAdmissionDetails(*admission, labs, *patient)
	if err != nil {
		return nil, fmt.Errorf("failed to convert admission to dto: %w", err)
	}

	resource, err := s.authorizationClient.GetResource(ctx, &authorization.GetResourceRequest{Id: admission.ID.String()})
	if err != nil {
		return nil, fmt.Errorf("failed to get resource from authorization service: %w", err)
	}

	assignments := make(map[uuid.UUID]dto.User)
	for _, a := range resource.Resource.Assignments {
		user, err := s.userService.GetUser(ctx, uuid.MustParse(a.ActorId))
		if err != nil {
			return nil, fmt.Errorf("failed to get user: %w", err)
		}
		if user == nil {
			return nil, fmt.Errorf("user with id %v not found from resource assignments", a.ActorId)
		}

		user.Role = a.Role
		user.Permissions = a.Permissions

		assignments[uuid.MustParse(a.ActorId)] = *user
	}

	var pendingTransfer *dto.JourneyStep
	if resource.Resource.PendingTransfer != nil {
		pendingTransfer = &dto.JourneyStep{
			FromTeam:     resource.Resource.PendingTransfer.FromTeam,
			ToTeam:       resource.Resource.PendingTransfer.ToTeam,
			FromTeamLead: uuid.MustParse(resource.Resource.PendingTransfer.FromTeamLead),
			ToTeamLead:   uuid.MustParse(resource.Resource.PendingTransfer.ToTeamLead),
		}
	}

	admissionDetails.CareTeam = dto.CareTeam{
		Team:        resource.Resource.Team.Name,
		Department:  resource.Resource.Team.DisplayName,
		TeamLead:    uuid.MustParse(resource.Resource.TeamLead),
		Assignments: assignments,
		Journey: lo.Map(resource.Resource.Journey, func(js *authorization.JourneyStep, index int) dto.JourneyStep {
			parsedTime, _ := time.Parse(time.DateTime, js.TransferTime)

			return dto.JourneyStep{
				TransferTime: parsedTime,
				FromTeam:     js.FromTeam,
				ToTeam:       js.ToTeam,
				FromTeamLead: uuid.MustParse(js.FromTeamLead),
				ToTeamLead:   uuid.MustParse(js.ToTeamLead),
			}
		}),
		PendingTransfer: pendingTransfer,
	}

	return admissionDetails, nil
}

func (s *PatientService) AcceptTransferRequest(ctx context.Context, userId, admissionId uuid.UUID, accept bool) error {
	if _, err := s.authorizationClient.TransferResource(ctx, &authorization.TransferResourceRequest{
		Id:             admissionId.String(),
		ActorId:        userId.String(),
		AcceptTransfer: accept,
	}); err != nil {
		return fmt.Errorf("failed to accept transfer: %w", err)
	}

	return nil
}

func (s *PatientService) AcceptAdmissionRequest(ctx context.Context, userId, admissionId uuid.UUID) error {
	admission, err := s.repo.GetAdmission(ctx, admissionId)
	if err != nil {
		return fmt.Errorf("failed to get admission: %w", err)
	}
	if admission == nil {
		return fmt.Errorf("admission with id %v not found", admissionId)
	}

	admission.Status = models.StatusAdmitted

	if err = s.repo.UpdateAdmission(ctx, admission); err != nil {
		return fmt.Errorf("failed to update admission: %w", err)
	}

	if _, err := s.authorizationClient.TransferResource(ctx, &authorization.TransferResourceRequest{
		Id:             admissionId.String(),
		ActorId:        userId.String(),
		AcceptTransfer: true,
	}); err != nil {
		return fmt.Errorf("failed to accept transfer: %w", err)
	}

	return nil
}

func (s *PatientService) UpdateAdmission(ctx context.Context, admission dto.AdmissionDetails) (*dto.AdmissionDetails, error) {
	admissionModel, err := toAdmissionModel(admission.Patient.ID, admission)
	if err != nil {
		return nil, fmt.Errorf("failed to convert admission dto to model: %w", err)
	}

	if err = s.repo.UpdateAdmission(ctx, &admissionModel); err != nil {
		return nil, fmt.Errorf("failed to update admission: %w", err)
	}

	result, err := s.GetAdmissionDetails(ctx, admissionModel.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get admission details: %w", err)
	}
	if result == nil {
		return nil, fmt.Errorf("admission details returned nil")
	}

	return result, nil
}

func (s *PatientService) OrderLabTest(ctx context.Context, userId uuid.UUID, request dto.OrderLabTestRequest) error {
	// Check if admission exists
	admission, err := s.repo.GetAdmission(ctx, request.AdmissionId)
	if err != nil {
		return fmt.Errorf("failed to get admission: %w", err)
	}
	if admission == nil {
		return fmt.Errorf("admission with id %v not found", request.AdmissionId)
	}

	if err = s.repo.InsertLab(ctx, models.Lab{
		RequestedAt: time.Now(),
		TestType:    request.LabTest,
		AdmissionID: request.AdmissionId,
		RequestedBy: userId,
	}); err != nil {
		return fmt.Errorf("failed to insert lab test: %w", err)
	}

	return nil
}

func (s *PatientService) AddTeamMember(ctx context.Context, userId uuid.UUID, admissionId uuid.UUID) error {
	if _, err := s.authorizationClient.UpdateResourceAssignment(ctx, &authorization.UpdateResourceAssignmentRequest{
		ResourceId: admissionId.String(),
		ActorId:    userId.String(),
		Add:        true,
	}); err != nil {
		return fmt.Errorf("failed to add assignment: %w", err)
	}

	return nil
}

func (s *PatientService) RemoveTeamMember(ctx context.Context, userId uuid.UUID, admissionId uuid.UUID) error {
	if _, err := s.authorizationClient.UpdateResourceAssignment(ctx, &authorization.UpdateResourceAssignmentRequest{
		ResourceId: admissionId.String(),
		ActorId:    userId.String(),
		Add:        false,
	}); err != nil {
		return fmt.Errorf("failed to remove assignment: %w", err)
	}

	return nil
}

func (s *PatientService) AddTeamMemberPermissions(ctx context.Context, userId, admissionId uuid.UUID, section, permission string) error {
	if _, err := s.authorizationClient.AddPermission(ctx, &authorization.AddPermissionRequest{
		ActorId:    userId.String(),
		ResourceId: admissionId.String(),
		Section:    section,
		Permission: permission,
	}); err != nil {
		return fmt.Errorf("failed to add team member permissions: %w", err)
	}

	return nil
}

func (s *PatientService) RemoveTeamMemberPermissions(ctx context.Context, userId, admissionId uuid.UUID, section string) error {
	if _, err := s.authorizationClient.RemovePermission(ctx, &authorization.RemovePermissionRequest{
		ResourceId: admissionId.String(),
		ActorId:    userId.String(),
		Section:    section,
	}); err != nil {
		return fmt.Errorf("failed to remove team member permissions: %w", err)
	}

	return nil
}

func (s *PatientService) RequestPatientTransfer(ctx context.Context, admissionId uuid.UUID, toTeam string, toTeamLead uuid.UUID) error {
	if _, err := s.authorizationClient.RequestResourceTransfer(ctx, &authorization.RequestResourceTransferRequest{
		ResourceId: admissionId.String(),
		ToTeam:     toTeam,
		ToTeamLead: toTeamLead.String(),
	}); err != nil {
		return fmt.Errorf("failed to request transfer: %w", err)
	}

	return nil
}

func (s *PatientService) DischargePatient(ctx context.Context, admissionId uuid.UUID) error {
	admission, err := s.repo.GetAdmission(ctx, admissionId)
	if err != nil {
		return fmt.Errorf("failed to get admission: %w", err)
	}
	if admission == nil {
		return fmt.Errorf("admission with id %v not found", admissionId)
	}

	admission.Status = models.StatusDischarged
	admission.EndTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	if err = s.repo.UpdateAdmission(ctx, admission); err != nil {
		return fmt.Errorf("failed to update admission: %w", err)
	}

	if _, err = s.authorizationClient.ArchiveResource(ctx, &authorization.ArchiveResourceRequest{Id: admissionId.String()}); err != nil {
		return fmt.Errorf("failed to archive resource: %w", err)
	}

	return nil
}
