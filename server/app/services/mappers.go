package services

import (
	"encoding/json"
	"fmt"
	"hospital-system/proto_gen/authorization/v1"
	"hospital-system/server/app/dto"
	"hospital-system/server/models"

	"github.com/google/uuid"
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

func toDtoPatient(patient models.Patient) dto.Patient {
	return dto.Patient{
		ID:                           patient.ID,
		Firstname:                    patient.Firstname,
		Lastname:                     patient.Lastname,
		NationalIdentificationNumber: patient.NationalIdentificationNumber,
		MedicalRecordNumber:          patient.MedicalRecordNumber,
		Email:                        patient.Email,
		PhoneNumber:                  patient.PhoneNumber,
	}
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

	return models.Admission{
		Anamnesis: marshalledAnamnesis,
		Logs:      marshalledLogs,
		PatientID: patientId,
	}, nil
}
