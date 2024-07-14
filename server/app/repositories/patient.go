package repositories

import (
	"context"
	"errors"
	"hospital-system/server/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *RepositoryImpl) InsertPatient(ctx context.Context, patient models.Patient) (*models.Patient, error) {
	if err := repo.db.WithContext(ctx).Create(&patient).Error; err != nil {
		return nil, err
	}

	return &patient, nil
}

func (repo *RepositoryImpl) GetPatient(ctx context.Context, id uuid.UUID) (*models.Patient, error) {
	var patient models.Patient

	if err := repo.db.WithContext(ctx).Where("id = ?", id).First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &patient, nil

}

func (repo *RepositoryImpl) GetPatientByPersonalID(ctx context.Context, personalID string) (*models.Patient, error) {
	var patient models.Patient

	if err := repo.db.WithContext(ctx).Where("national_identification_number = ?", personalID).
		Or("medical_record_number = ?", personalID).First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &patient, nil
}
