package patient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"hospital-system/proto_gen/authorization/v1"
	"hospital-system/server/app/repositories/admission"
	"hospital-system/server/app/repositories/patient"
	"hospital-system/server/models"
)

type ServiceImpl struct {
	authorizationClient authorization.AuthorizationServiceClient
	patientRepo         patient.Repository
	admissionRepo       admission.Repository
}

func NewService(authorizationClient authorization.AuthorizationServiceClient, patientRepo patient.Repository, admissionRepo admission.Repository) *ServiceImpl {
	return &ServiceImpl{
		authorizationClient: authorizationClient,
		patientRepo:         patientRepo,
		admissionRepo:       admissionRepo,
	}
}

func (service *ServiceImpl) CreatePatient(patient Patient) (*Patient, error) {
	id := uuid.New()

	patientResponse, err := service.patientRepo.Insert(models.Patient{
		ID:                           id,
		Firstname:                    patient.Firstname,
		Lastname:                     patient.Lastname,
		NationalIdentificationNumber: patient.NationalIdentificationNumber,
		MedicalRecordNumber:          patient.MedicalRecordNumber,
		Email:                        patient.Email,
		PhoneNumber:                  patient.PhoneNumber,
		Birthday:                     patient.Birthday,
		Gender:                       patient.Gender,
	})
	if err != nil {
		return nil, err
	}

	return &Patient{
		ID:                           patientResponse.ID,
		Firstname:                    patientResponse.Firstname,
		Lastname:                     patientResponse.Lastname,
		NationalIdentificationNumber: patientResponse.NationalIdentificationNumber,
		MedicalRecordNumber:          patientResponse.MedicalRecordNumber,
		Email:                        patientResponse.Email,
		PhoneNumber:                  patientResponse.PhoneNumber,
		Birthday:                     patientResponse.Birthday,
		Gender:                       patientResponse.Gender,
		Admissions:                   nil,
	}, nil
}

func (service *ServiceImpl) GetPatient(id string) (*Patient, error) {
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

	return &Patient{
		ID:                           patient.ID,
		Firstname:                    patient.Firstname,
		Lastname:                     patient.Lastname,
		NationalIdentificationNumber: patient.NationalIdentificationNumber,
		MedicalRecordNumber:          patient.MedicalRecordNumber,
		Email:                        patient.Email,
		PhoneNumber:                  patient.PhoneNumber,
		Birthday:                     patient.Birthday,
		Gender:                       patient.Gender,
		Admissions: lo.Map(admissions, func(admission models.Admission, _ int) Admission {
			return Admission{
				ID:        admission.ID,
				StartTime: admission.StartTime,
				EndTime:   admission.EndTime,
				Status:    admission.Status,
				Diagnosis: admission.Diagnosis,
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

func (service *ServiceImpl) RegisterPatientAdmission(patientId uuid.UUID, admission Admission) error {
	medicationsMarshalled, stdErr := json.Marshal(admission.Medications)
	if stdErr != nil {
		slog.Error("Failed to marshal medications")
		return stdErr
	}

	allergiesMarshalled, stdErr := json.Marshal(admission.Allergies)
	if stdErr != nil {
		slog.Error("Failed to marshal allergies")
		return stdErr
	}

	result, err := service.admissionRepo.Insert(models.Admission{
		PatientID:   patientId,
		Symptoms:    admission.Symptoms,
		Medications: medicationsMarshalled,
		Allergies:   allergiesMarshalled,
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
		err = service.admissionRepo.Delete(result.ID)
		if err != nil {
			slog.Error("Failed to delete admission", err)
			return err
		}

		return err
	}

	return nil
}

func (service *ServiceImpl) GetAdmissionsByStatuses(statuses []string) ([]Admission, error) {
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

	return lo.Map(admissions, func(a models.Admission, _ int) Admission {
		return Admission{
			ID:          a.ID,
			StartTime:   a.StartTime,
			EndTime:     a.EndTime,
			Status:      a.Status,
			Symptoms:    a.Symptoms,
			Medications: []string{},
			Allergies:   []string{},
			Diagnosis:   a.Diagnosis,
			PatientID:   a.PatientID,
			Department: lo.FindOrElse(resources.GetResources(), &authorization.Resource{}, func(item *authorization.Resource) bool {
				return item.Id == a.ID.String()
			}).Team.DisplayName,
			Physician: uuid.MustParse(lo.FindOrElse(resources.GetResources(), &authorization.Resource{}, func(item *authorization.Resource) bool {
				return item.Id == a.ID.String()
			}).TeamLead),
		}
	}), nil
}

func (service *ServiceImpl) Discharge(admissionId uuid.UUID) error {
	admission, err := service.admissionRepo.Get(admissionId)
	if err != nil {
		return err
	}
	if admission == nil {
		return nil
	}

	admission.Status = models.StatusDischarged

	_, err = service.admissionRepo.Update(admission)
	if err != nil {
		return err
	}

	//// TODO: Handle on authorization side, update resource to archived
	//err = service.authorizationClient.ArchiveResource(context.Background(), &authorization.ArchiveResourceRequest{
	//	Id: admissionId.String(),
	//})
	if _, err = service.authorizationClient.ArchiveResource(context.Background(), &authorization.ArchiveResourceRequest{
		Id: admissionId.String(),
	}); err != nil {
		slog.Error("Failed to remove resource from authorization", err)
		return err
	}

	return nil
}
