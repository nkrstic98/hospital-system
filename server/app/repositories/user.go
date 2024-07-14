package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hospital-system/server/models"
)

func (repo *RepositoryImpl) InsertUser(user models.User) error {
	return repo.db.Create(&user).Error
}

func (repo *RepositoryImpl) GetUser(id uuid.UUID) (*models.User, error) {
	var user models.User

	if err := repo.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *RepositoryImpl) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *RepositoryImpl) GetUserByIDs(ids []uuid.UUID) ([]models.User, error) {
	var users []models.User

	if err := repo.db.Where("id IN ?", ids).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}

func (repo *RepositoryImpl) GetAllUsers() ([]models.User, error) {
	var users []models.User

	if err := repo.db.Order("username ASC").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}

func (repo *RepositoryImpl) DeleteUser(id uuid.UUID) error {
	return repo.db.Clauses(
		clause.Locking{Strength: clause.LockingStrengthUpdate},
	).Where("id = ?", id.String()).Delete(&models.User{}).Error
}
