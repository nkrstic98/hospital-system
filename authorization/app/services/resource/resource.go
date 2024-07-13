package resource

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"hospital-system/authorization/app/repositories/resource"
	"hospital-system/authorization/app/services/actor"
	"hospital-system/authorization/models"
)

type ServiceImpl struct {
	resourceRepo resource.Repository
	actorService actor.Service
}

func NewService(resourceRepo resource.Repository, ctorService actor.Service) *ServiceImpl {
	return &ServiceImpl{
		resourceRepo: resourceRepo,
		actorService: ctorService,
	}
}

func (s *ServiceImpl) AddResource(request Resource) error {
	actor, err := s.actorService.GetActor(request.TeamLead)
	if err != nil {
		slog.Error("Failed to fetch actor", request.ID, err)
		return err
	}

	assignments := []Assignment{
		{
			ActorID:     request.TeamLead,
			Role:        actor.Role,
			Permissions: actor.Permissions,
		},
	}

	marshalledAssignments, stdErr := json.Marshal(assignments)
	if stdErr != nil {
		slog.Error("Failed to marshal assignments", request.ID, err)
	}

	return s.resourceRepo.Insert(models.Resource{
		ID:              request.ID,
		Team:            request.Team,
		TeamLead:        request.TeamLead,
		TeamAssignments: marshalledAssignments,
	})
}

func (s *ServiceImpl) GetResources(ids *[]string) ([]Resource, error) {
	var resources []models.Resource
	var err error

	if ids != nil {
		resources, err = s.resourceRepo.GetByIDs(*ids)
	} else {
		resources, err = s.resourceRepo.GetAll()
	}
	if err != nil {
		slog.Error("Failed to fetch resources", err)
		return nil, err
	}

	return lo.Map(resources, func(r models.Resource, _ int) Resource {
		return Resource{
			ID:       r.ID,
			Team:     r.Team,
			TeamLead: r.TeamLead,
		}
	}), nil
}

func (s *ServiceImpl) ArchiveResource(id uuid.UUID) error {
	return s.resourceRepo.UpdateArchived(id)
}
