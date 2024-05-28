package patient

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hospital-system/server/models"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

func (repo *RepositoryImpl) Insert(patient models.Patient) (*models.Patient, error) {
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		if result := repo.db.Create(&patient); result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &patient, nil
}

func (repo *RepositoryImpl) Get(id uuid.UUID) (*models.Patient, error) {
	var patient models.Patient
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Clauses(
			clause.Locking{Strength: clause.LockingStrengthShare},
		).Where("id = ?", id.String()).First(&patient)
		if result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &patient, nil

}

func (repo *RepositoryImpl) GetByPersonalID(personalID string) (*models.Patient, error) {
	var patient models.Patient
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Clauses(
			clause.Locking{Strength: clause.LockingStrengthShare},
		).Where("national_identification_number = ?", personalID).
			Or("medical_record_number = ?", personalID).First(&patient)
		if result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &patient, nil
}
