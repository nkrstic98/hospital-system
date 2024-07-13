package admission

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hospital-system/server/models"
	"time"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

func (repo *RepositoryImpl) Insert(admission models.Admission) (*models.Admission, error) {
	admission.ID = uuid.New()
	admission.StartTime = time.Now()
	admission.Status = models.StatusPending

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Create(&admission)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
	if err != nil {
		return nil, err

	}

	return &admission, nil
}

func (repo *RepositoryImpl) Get(id uuid.UUID) (*models.Admission, error) {
	var admission models.Admission
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Clauses(
			clause.Locking{Strength: clause.LockingStrengthShare},
		).Where("id = ?", id.String()).First(&admission)
		if result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &admission, nil
}

func (repo *RepositoryImpl) Update(admission *models.Admission) (*models.Admission, error) {
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Save(admission)
		if result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return admission, nil
}

func (repo *RepositoryImpl) Delete(id uuid.UUID) error {
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Where("id = ?", id.String()).Delete(&models.Admission{})
		if result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *RepositoryImpl) GetByPatientId(id uuid.UUID) ([]models.Admission, error) {
	var admissions []models.Admission
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Clauses(
			clause.Locking{Strength: clause.LockingStrengthShare},
		).Where("patient_id = ?", id.String()).Find(&admissions)
		if result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return admissions, nil
}

func (repo *RepositoryImpl) GetByStatuses(statuses []string) ([]models.Admission, error) {
	var admissions []models.Admission
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Clauses(
			clause.Locking{Strength: clause.LockingStrengthShare},
		).Where("status IN (?)", statuses).Find(&admissions)
		if result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return admissions, nil
}

func (repo *RepositoryImpl) GetByIDs(ids []uuid.UUID) ([]models.Admission, error) {
	var admissions []models.Admission
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Clauses(
			clause.Locking{Strength: clause.LockingStrengthShare},
		).Where("id IN (?)", ids).Find(&admissions)
		if result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return admissions, nil
}

func (repo *RepositoryImpl) GetLabsByAdmissionID(admissionID uuid.UUID) ([]models.Lab, error) {
	var labs []models.Lab
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Clauses(
			clause.Locking{Strength: clause.LockingStrengthShare},
		).Where("admission_id = ?", admissionID.String()).Find(&labs)
		if result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return labs, nil
}
