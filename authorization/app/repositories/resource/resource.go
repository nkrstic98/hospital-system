package resource

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/samber/lo"
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

func (r *RepositoryImpl) Get(ctx context.Context, id uuid.UUID) (*models.Resource, error) {
	var resource models.Resource
	if err := r.db.WithContext(ctx).Where("id = ?", id.String()).First(&resource).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &resource, nil
}

func (r *RepositoryImpl) GetByIDs(ids []string) ([]models.Resource, error) {
	var resources []models.Resource
	err := r.db.Where("id IN ?", ids).Find(&resources).Error
	return resources, err
}

func (r *RepositoryImpl) GetByActorID(ctx context.Context, actorId uuid.UUID, archived bool) ([]models.Resource, error) {
	var resources []models.Resource

	sql := "SELECT * FROM resources r, jsonb_array_elements(r.team_assignments::jsonb) AS ta WHERE (ta::jsonb->>'actor_id' = ? OR r.pending_transfer->>'to_team_lead' = ?)"
	sql += lo.Ternary(archived, " AND r.archived IS NOT NULL", " AND r.archived IS NULL")

	err := r.db.WithContext(ctx).Raw(sql, actorId, actorId).Scan(&resources).Error
	if err != nil {
		return nil, err
	}

	return resources, err
}

func (r *RepositoryImpl) UpdateResource(ctx context.Context, resource *models.Resource) error {
	return r.db.WithContext(ctx).Model(&models.Resource{}).Where("id = ?", resource.ID.String()).Updates(map[string]interface{}{
		"team":             resource.Team,
		"team_lead":        resource.TeamLead,
		"team_assignments": resource.TeamAssignments,
		"journey":          resource.Journey,
		"pending_transfer": resource.PendingTransfer,
	}).Error
}

func (r *RepositoryImpl) UpdateArchived(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&models.Resource{}).Where("id = ?", id.String()).Update("archived", true).Error
}
