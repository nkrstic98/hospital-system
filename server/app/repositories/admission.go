package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hospital-system/server/models"
)

func (repo *RepositoryImpl) InsertAdmission(admission models.Admission) (*models.Admission, error) {
	if err := repo.db.Create(&admission).Error; err != nil {
		return nil, err
	}

	return &admission, nil
}

func (repo *RepositoryImpl) GetAdmission(id uuid.UUID) (*models.Admission, error) {
	var admission models.Admission

	if err := repo.db.Where("id = ?", id).First(&admission).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &admission, nil
}

func (repo *RepositoryImpl) DeleteAdmission(id uuid.UUID) error {
	return repo.db.Where("id = ?", id.String()).Delete(&models.Admission{}).Error
}

func (repo *RepositoryImpl) GetAdmissionsByPatientId(id uuid.UUID) ([]models.Admission, error) {
	var admissions []models.Admission

	if err := repo.db.Where("patient_id = ?", id).Find(&admissions).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, nil
		}
		return nil, err
	}

	return admissions, nil
}

func (repo *RepositoryImpl) GetAdmissionsByStatuses(statuses []string) ([]models.Admission, error) {
	var admissions []models.Admission
	if err := repo.db.Where("status IN (?)", statuses).Find(&admissions).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, nil
		}
		return nil, err
	}

	return admissions, nil
}

func (repo *RepositoryImpl) GetAdmissionsByIDs(ids []uuid.UUID) ([]models.Admission, error) {
	var admissions []models.Admission

	if err := repo.db.Where("id IN (?)", ids).Find(&admissions).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, nil
		}
		return nil, err
	}

	return admissions, nil
}
