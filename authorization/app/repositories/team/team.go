package team

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

func (repo *RepositoryImpl) Insert(team models.Team) error {
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Create(&team)
		if result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *RepositoryImpl) Get(id uint) (models.Team, error) {
	var team models.Team
	result := repo.db.Where("id = ?", id).First(&team)
	if result.Error != nil {
		return models.Team{}, result.Error
	}

	return team, nil
}

func (repo *RepositoryImpl) GetByName(name string) (models.Team, error) {
	var team models.Team
	result := repo.db.Where("name = ?", name).First(&team)
	if result.Error != nil {
		return models.Team{}, result.Error
	}

	return team, nil
}

func (repo *RepositoryImpl) GetAll() ([]models.Team, error) {
	var teams []models.Team
	result := repo.db.Find(&teams)
	if result.Error != nil {
		return nil, result.Error
	}

	return teams, nil
}
