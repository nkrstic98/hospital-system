package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"hospital-system/server/models"

	"github.com/google/uuid"
)

func (repo *RepositoryImpl) InsertLab(ctx context.Context, lab models.Lab) error {
	return repo.db.WithContext(ctx).Create(&lab).Error
}

func (repo *RepositoryImpl) GetLabs(ctx context.Context) ([]models.Lab, error) {
	var labs []models.Lab

	if err := repo.db.WithContext(ctx).Order("requested_at DESC").Find(&labs).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, nil
		}
		return nil, err
	}

	return labs, nil
}

func (repo *RepositoryImpl) GetLab(ctx context.Context, id uuid.UUID) (*models.Lab, error) {
	var lab models.Lab

	if err := repo.db.WithContext(ctx).Where("id = ?", id).First(&lab).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &lab, nil
}

func (repo *RepositoryImpl) GetLabsByAdmissionID(ctx context.Context, admissionID uuid.UUID) ([]models.Lab, error) {
	var labs []models.Lab

	if err := repo.db.WithContext(ctx).Where("admission_id = ?", admissionID).Find(&labs).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, nil
		}
		return nil, err
	}

	return labs, nil
}

func (repo *RepositoryImpl) UpdateLab(ctx context.Context, lab *models.Lab) error {
	return repo.db.WithContext(ctx).Model(lab).Updates(map[string]interface{}{
		"processed_at": lab.ProcessedAt,
		"test_results": lab.TestResults,
		"processed_by": lab.ProcessedBy,
	}).Error
}
