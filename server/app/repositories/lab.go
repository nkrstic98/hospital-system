package repositories

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"hospital-system/server/models"
)

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
