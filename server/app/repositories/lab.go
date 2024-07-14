package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hospital-system/server/models"
)

func (repo *RepositoryImpl) GetLabsByAdmissionID(admissionID uuid.UUID) ([]models.Lab, error) {
	var labs []models.Lab

	if err := repo.db.Where("admission_id = ?", admissionID).Find(&labs).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, nil
		}
		return nil, err
	}

	return labs, nil
}
