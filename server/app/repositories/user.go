package repositories

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hospital-system/server/models"
)

func (repo *RepositoryImpl) InsertUser(ctx context.Context, user models.User) (*uuid.UUID, error) {
	if err := repo.db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}

	return &user.ID, nil
}

func (repo *RepositoryImpl) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User

	if err := repo.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *RepositoryImpl) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	if err := repo.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *RepositoryImpl) GetUsersByIDs(ctx context.Context, ids []uuid.UUID) ([]models.User, error) {
	var users []models.User

	if err := repo.db.WithContext(ctx).Where("id IN ?", ids).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}

func (repo *RepositoryImpl) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User

	if err := repo.db.WithContext(ctx).Order("username ASC").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}

func (repo *RepositoryImpl) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return repo.db.WithContext(ctx).Clauses(
		clause.Locking{Strength: clause.LockingStrengthUpdate},
	).Where("id = ?", id.String()).Delete(&models.User{}).Error
}
