package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hospital-system/server/models"
)

func (repo *RepositoryImpl) InsertPatient(patient models.Patient) (*models.Patient, error) {
	if err := repo.db.Create(&patient).Error; err != nil {
		return nil, err
	}

	return &patient, nil
}

func (repo *RepositoryImpl) GetPatient(id uuid.UUID) (*models.Patient, error) {
	var patient models.Patient

	if err := repo.db.Where("id = ?", id).First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &patient, nil

}

func (repo *RepositoryImpl) GetPatientByPersonalID(personalID string) (*models.Patient, error) {
	var patient models.Patient

	if err := repo.db.Where("national_identification_number = ?", personalID).
		Or("medical_record_number = ?", personalID).First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &patient, nil
}
