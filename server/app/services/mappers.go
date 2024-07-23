package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"hospital-system/proto_gen/authorization/v1"
	"hospital-system/server/app/dto"
	"hospital-system/server/models"
)

func toDtoUser(user models.User, actor *authorization.Actor) dto.User {
	dtoUser := dto.User{
		ID:                           user.ID,
		Firstname:                    user.Firstname,
		Lastname:                     user.Lastname,
		NationalIdentificationNumber: user.NationalIdentificationNumber,
		Username:                     user.Username,
		Email:                        user.Email,
	}

	if actor != nil {
		dtoUser.Role = actor.Role
		dtoUser.Team = actor.Team
		dtoUser.Permissions = actor.Permissions
	}

	return dtoUser
}

func toDtoPatient(patient models.Patient, admissions []models.Admission) dto.Patient {
	var adms []dto.Admission
	if len(admissions) > 0 {
		adms = lo.Map(admissions, func(a models.Admission, _ int) dto.Admission {
			return dto.Admission{
				ID:        a.ID,
				StartTime: a.StartTime,
				EndTime:   lo.Ternary(a.EndTime.Valid, &a.EndTime.Time, nil),
				Status:    a.Status,
			}
		})
	}

	return dto.Patient{
		ID:                           patient.ID,
		Firstname:                    patient.Firstname,
		Lastname:                     patient.Lastname,
		NationalIdentificationNumber: patient.NationalIdentificationNumber,
		MedicalRecordNumber:          patient.MedicalRecordNumber,
		Email:                        patient.Email,
		PhoneNumber:                  patient.PhoneNumber,
		Admissions:                   adms,
	}
}

func toDtoLabs(labs []models.Lab) ([]dto.Lab, error) {
	if len(labs) == 0 {
		return nil, nil
	}

	dtoLabs := make([]dto.Lab, 0, len(labs))

	for _, lab := range labs {
		var testResult *[]dto.LabTest
		if lab.TestResults != nil {
			if err := json.Unmarshal(lab.TestResults, &testResult); err != nil {
				return nil, fmt.Errorf("failed to unmarshal test results: %w", err)
			}
		}

		dtoLabs = append(dtoLabs, dto.Lab{
			ID:          lab.ID,
			RequestedAt: lab.RequestedAt,
			ProcessedAt: lo.Ternary(lab.ProcessedAt.Valid, &lab.ProcessedAt.Time, nil),
			TestType:    lab.TestType,
			TestResults: testResult,
			RequestedBy: lab.RequestedBy,
			ProcessedBy: lab.ProcessedBy,
		})
	}

	return dtoLabs, nil
}

func toDtoAdmissionDetails(admission models.Admission, labs []models.Lab, patient models.Patient) (*dto.AdmissionDetails, error) {
	var anamnesis dto.Anamnesis
	if err := json.Unmarshal(admission.Anamnesis, &anamnesis); err != nil {
		return nil, fmt.Errorf("failed to unmarshal anamnesis: %w", err)
	}

	var vitals dto.Vitals
	if admission.Vitals != nil {
		if err := json.Unmarshal(admission.Vitals, &vitals); err != nil {
			return nil, fmt.Errorf("failed to unmarshal vitals: %w", err)
		}
	}

	var medications []dto.MedicationInfo
	if admission.Medications != nil {
		if err := json.Unmarshal(admission.Medications, &medications); err != nil {
			return nil, fmt.Errorf("failed to unmarshal medications: %w", err)
		}
	}

	var logs []dto.Log
	if admission.Logs != nil {
		if err := json.Unmarshal(admission.Logs, &logs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal logs: %w", err)
		}
	}

	dtoLabs, err := toDtoLabs(labs)
	if err != nil {
		return nil, fmt.Errorf("failed to convert labs to dto: %w", err)
	}

	return &dto.AdmissionDetails{
		ID:          admission.ID,
		StartTime:   admission.StartTime,
		EndTime:     lo.Ternary(admission.EndTime.Valid, &admission.EndTime.Time, nil),
		Status:      admission.Status,
		Anamnesis:   anamnesis,
		Vitals:      vitals,
		Diagnosis:   lo.Ternary(admission.Diagnosis.Valid, &admission.Diagnosis.String, nil),
		Medications: medications,
		Labs:        dtoLabs,
		Logs:        logs,
		Patient:     toDtoPatient(patient, nil),
	}, nil
}

func toAdmissionModel(patientId uuid.UUID, admission dto.AdmissionDetails) (models.Admission, error) {
	marshalledAnamnesis, stdErr := json.Marshal(admission.Anamnesis)
	if stdErr != nil {
		return models.Admission{}, fmt.Errorf("failed to marshal anamnesis: %w", stdErr)
	}

	marshalledLogs, stdErr := json.Marshal(admission.Logs)
	if stdErr != nil {
		return models.Admission{}, fmt.Errorf("failed to marshal logs: %w", stdErr)
	}

	marshalledVitals, stdErr := json.Marshal(admission.Vitals)
	if stdErr != nil {
		return models.Admission{}, fmt.Errorf("failed to marshal vitals: %w", stdErr)
	}

	if admission.Medications == nil {
		admission.Medications = []dto.MedicationInfo{}
	}
	marshalledMedications, stdErr := json.Marshal(admission.Medications)
	if stdErr != nil {
		return models.Admission{}, fmt.Errorf("failed to marshal medications: %w", stdErr)
	}

	endTime := sql.NullTime{}
	if admission.EndTime != nil {
		endTime = sql.NullTime{
			Time:  *admission.EndTime,
			Valid: true,
		}
	}

	diagnosis := sql.NullString{}
	if admission.Diagnosis != nil {
		diagnosis = sql.NullString{
			String: *admission.Diagnosis,
			Valid:  true,
		}
	}

	return models.Admission{
		ID:          admission.ID,
		StartTime:   admission.StartTime,
		EndTime:     endTime,
		Status:      admission.Status,
		Anamnesis:   marshalledAnamnesis,
		Vitals:      marshalledVitals,
		Diagnosis:   diagnosis,
		Medications: marshalledMedications,
		Logs:        marshalledLogs,
		PatientID:   patientId,
	}, nil
}
