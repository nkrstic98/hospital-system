package resource

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	AddResource(request Resource) error
	GetResource(ctx context.Context, id uuid.UUID) (*Resource, error)
	GetResources(ctx context.Context, ids *[]string, actorId *uuid.UUID, archived bool) ([]Resource, error)
	TransferResource(ctx context.Context, id uuid.UUID) error
	DeclineTransfer(ctx context.Context, id uuid.UUID) error
	UpdateResourceAssignment(ctx context.Context, actorId, resourceId uuid.UUID, add bool) error
	AddPermission(ctx context.Context, actorId, resourceId uuid.UUID, section, permission string) error
	RemovePermission(ctx context.Context, actorId, resourceId uuid.UUID, section string) error
	RequestResourceTransfer(ctx context.Context, resourceId uuid.UUID, toTeam string, toTeamLead uuid.UUID) error
	ArchiveResource(ctx context.Context, id uuid.UUID) error
}
