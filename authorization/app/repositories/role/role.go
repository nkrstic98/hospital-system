package role

import (
	"gorm.io/gorm"
	"hospital-system/authorization/models"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

func (repo *RepositoryImpl) Insert(role models.Role) error {
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Create(&role)
		if result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *RepositoryImpl) Get(id uint) (models.Role, error) {
	var role models.Role
	result := repo.db.Where("id = ?", id).First(&role)
	if result.Error != nil {
		return models.Role{}, result.Error
	}

	return role, nil
}

func (repo *RepositoryImpl) GetByName(name string) (models.Role, error) {
	var role models.Role
	result := repo.db.Where("name = ?", name).First(&role)
	if result.Error != nil {
		return models.Role{}, result.Error
	}

	return role, nil
}

func (repo *RepositoryImpl) GetAll() ([]models.Role, error) {
	var roles []models.Role
	result := repo.db.Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}
