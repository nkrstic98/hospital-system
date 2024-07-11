package user

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
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

func (repo *RepositoryImpl) Insert(user models.User) error {
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		if result := repo.db.Create(&user); result.Error != nil {
			slog.Error(fmt.Sprintf("Failed to create user with email %s: %s", user.Email, result.Error.Error()))
			return result.Error
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *RepositoryImpl) Get(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Clauses(
			clause.Locking{Strength: clause.LockingStrengthShare},
		).Where("id = ?", id.String()).First(&user)
		if result.Error != nil {
			slog.Error(fmt.Sprintf("Failed to get user with id %v: %s", id, result.Error.Error()))
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *RepositoryImpl) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Clauses(
			clause.Locking{Strength: clause.LockingStrengthShare},
		).Where("username = ?", username).First(&user)
		if result.Error != nil {
			slog.Error(fmt.Sprintf("Failed to get user with username %s: %s", username, result.Error.Error()))
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *RepositoryImpl) GetByIDs(ids []uuid.UUID) ([]models.User, error) {
	var users []models.User
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Clauses(
			clause.Locking{Strength: clause.LockingStrengthShare},
		).Where("id IN ?", ids).Find(&users)
		if result.Error != nil {
			slog.Error(fmt.Sprintf("Failed to get users with ids %v: %s", ids, result.Error.Error()))
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *RepositoryImpl) GetAll() ([]models.User, error) {
	var users []models.User
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Order("username ASC").Find(&users)
		if result.Error != nil {
			slog.Error(fmt.Sprintf("Failed to get all users: %s", result.Error.Error()))
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *RepositoryImpl) Delete(id uuid.UUID) error {
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Clauses(
			clause.Locking{Strength: clause.LockingStrengthUpdate},
		).Where("id = ?", id.String()).Delete(&models.User{})
		if result.Error != nil {
			slog.Error(fmt.Sprintf("Failed to delete user with id %v: %s", id, result.Error.Error()))
			return result.Error
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
