package repositories

import (
	"context"
	"errors"

	"hospital-system/server/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *RepositoryImpl) InsertAdmission(ctx context.Context, admission models.Admission) (*models.Admission, error) {
	if err := repo.db.WithContext(ctx).Create(&admission).Error; err != nil {
		return nil, err
	}

	return &admission, nil
}

func (repo *RepositoryImpl) GetAdmission(ctx context.Context, id uuid.UUID) (*models.Admission, error) {
	var admission models.Admission

	if err := repo.db.WithContext(ctx).Where("id = ?", id).First(&admission).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &admission, nil
}

func (repo *RepositoryImpl) UpdateAdmission(ctx context.Context, admission *models.Admission) error {
	return repo.db.WithContext(ctx).Model(&models.Admission{}).Where("id = ?", admission.ID).Updates(map[string]interface{}{
		"start_time":  admission.StartTime,
		"end_time":    admission.EndTime,
		"status":      admission.Status,
		"vitals":      admission.Vitals,
		"diagnosis":   admission.Diagnosis,
		"medications": admission.Medications,
		"logs":        admission.Logs,
	}).Error
}

func (repo *RepositoryImpl) DeleteAdmission(ctx context.Context, id uuid.UUID) error {
	return repo.db.WithContext(ctx).Where("id = ?", id.String()).Delete(&models.Admission{}).Error
}

func (repo *RepositoryImpl) GetAdmissionsByPatientId(ctx context.Context, id uuid.UUID) ([]models.Admission, error) {
	var admissions []models.Admission

	if err := repo.db.WithContext(ctx).Where("patient_id = ?", id).Find(&admissions).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, nil
		}
		return nil, err
	}

	return admissions, nil
}

func (repo *RepositoryImpl) GetAdmissionsByStatuses(ctx context.Context, statuses []string) ([]models.Admission, error) {
	var admissions []models.Admission
	if err := repo.db.WithContext(ctx).Where("status IN (?)", statuses).Order("start_time DESC").Find(&admissions).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, nil
		}
		return nil, err
	}

	return admissions, nil
}

func (repo *RepositoryImpl) GetAdmissionsByIDs(ctx context.Context, ids []uuid.UUID) ([]models.Admission, error) {
	var admissions []models.Admission

	if err := repo.db.WithContext(ctx).Where("id IN (?)", ids).Find(&admissions).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, nil
		}
		return nil, err
	}

	return admissions, nil
}
