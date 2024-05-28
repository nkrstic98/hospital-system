package resource

import (
	"github.com/google/uuid"
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

func (r *RepositoryImpl) Insert(resource models.Resource) error {
	return r.db.Create(&resource).Error
}

func (r *RepositoryImpl) GetByIDs(ids []string) ([]models.Resource, error) {
	var resources []models.Resource
	err := r.db.Where("id IN ?", ids).Find(&resources).Error
	return resources, err
}

func (r *RepositoryImpl) GetAll() ([]models.Resource, error) {
	var resources []models.Resource
	err := r.db.Find(&resources).Error
	return resources, err
}

func (r *RepositoryImpl) UpdateArchived(id uuid.UUID) error {
	return r.db.Model(&models.Resource{}).Where("id = ?", id.String()).Update("archived", true).Error
}
