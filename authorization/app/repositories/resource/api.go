package resource

import (
	"context"
	"github.com/google/uuid"
	"hospital-system/authorization/models"
)

type Repository interface {
	Insert(resource models.Resource) error
	Get(ctx context.Context, id uuid.UUID) (*models.Resource, error)
	GetByIDs(ids []string) ([]models.Resource, error)
	GetByActorID(ctx context.Context, actorId uuid.UUID, archived bool) ([]models.Resource, error)
	UpdateResource(ctx context.Context, resource *models.Resource) error
	UpdateArchived(ctx context.Context, id uuid.UUID) error
}
