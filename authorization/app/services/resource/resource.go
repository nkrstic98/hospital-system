package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"hospital-system/authorization/app/repositories/resource"
	"hospital-system/authorization/app/services/actor"
	"hospital-system/authorization/models"
	"time"
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
		return stdErr
	}

	var marshalledPendingTransfer []byte
	if request.PendingTransfer != nil {
		marshalledPendingTransfer, stdErr = json.Marshal(request.PendingTransfer)
		if stdErr != nil {
			slog.Error("Failed to marshal pending transfer", request.ID, err)
			return stdErr
		}
	}

	return s.resourceRepo.Insert(models.Resource{
		ID:              request.ID,
		Team:            request.Team,
		TeamLead:        request.TeamLead,
		TeamAssignments: marshalledAssignments,
		PendingTransfer: marshalledPendingTransfer,
	})
}

func (s *ServiceImpl) GetResource(ctx context.Context, id uuid.UUID) (*Resource, error) {
	resource, err := s.resourceRepo.Get(ctx, id)
	if err != nil {
		slog.Error("Failed to fetch resource", id, err)
		return nil, err
	}

	var assignments []Assignment
	if err := json.Unmarshal(resource.TeamAssignments, &assignments); err != nil {
		slog.Error("Failed to unmarshal assignments", id, err)
		return nil, err
	}

	var journey []JourneyStep
	if resource.Journey != nil {
		if err := json.Unmarshal(resource.Journey, &journey); err != nil {
			slog.Error("Failed to unmarshal journey", id, err)
			return nil, err
		}
	}

	var pendingTransfer *JourneyStep
	if resource.PendingTransfer != nil {
		if err = json.Unmarshal(resource.PendingTransfer, &pendingTransfer); err != nil {
			slog.Error("Failed to unmarshal pending transfer", id, err)
			return nil, err
		}
	}

	return &Resource{
		ID:              resource.ID,
		Team:            resource.Team,
		TeamLead:        resource.TeamLead,
		Assignments:     assignments,
		Journey:         journey,
		PendingTransfer: pendingTransfer,
		Archived:        lo.Ternary(resource.Archived.Valid, &resource.Archived.Bool, nil),
	}, nil

}

func (s *ServiceImpl) GetResources(ctx context.Context, ids *[]string, actorId *uuid.UUID, archived bool) ([]Resource, error) {
	var resources []models.Resource
	var err error

	if ids != nil {
		resources, err = s.resourceRepo.GetByIDs(*ids)
	} else if actorId != nil {
		resources, err = s.resourceRepo.GetByActorID(ctx, *actorId, archived)
	}
	if err != nil {
		slog.Error("Failed to fetch resources", err)
		return nil, err
	}

	result := make([]Resource, 0, len(resources))
	for _, r := range resources {
		var pendingTransfer *JourneyStep
		if r.PendingTransfer != nil {
			if err := json.Unmarshal(r.PendingTransfer, &pendingTransfer); err != nil {
				slog.Error("Failed to unmarshal pending transfer", r.ID, err)
				return nil, err
			}
		}

		result = append(result, Resource{
			ID:              r.ID,
			Team:            r.Team,
			TeamLead:        r.TeamLead,
			PendingTransfer: pendingTransfer,
		})
	}

	return result, nil
}

func (s *ServiceImpl) TransferResource(ctx context.Context, id uuid.UUID) error {
	resourceResult, err := s.resourceRepo.Get(ctx, id)
	if err != nil {
		slog.Error("Failed to fetch resource", id, err)
		return err
	}
	if resourceResult == nil {
		return fmt.Errorf("resource %s not found", id)
	}

	if resourceResult.PendingTransfer != nil {
		var journey []JourneyStep
		if resourceResult.Journey != nil {
			if err = json.Unmarshal(resourceResult.Journey, &journey); err != nil {
				slog.Error("Failed to unmarshal journey")
				return err
			}
		}

		if len(journey) == 0 {
			journey = make([]JourneyStep, 0)
		}

		var pendingTransfer JourneyStep

		if err = json.Unmarshal(resourceResult.PendingTransfer, &pendingTransfer); err != nil {
			slog.Error("Failed to unmarshal pending transfer")
			return err
		}

		pendingTransfer.TransferTime = time.Now().Format(time.DateTime)

		journey = append(journey, pendingTransfer)

		marshalledJourney, err := json.Marshal(journey)
		if err != nil {
			slog.Error("Failed to marshal journey")
			return err
		}

		actorResult, err := s.actorService.GetActor(pendingTransfer.ToTeamLead)
		if err != nil {
			slog.Error("Failed to fetch actor", pendingTransfer.ToTeamLead, err)
			return err
		}

		teamAssignments := []Assignment{
			{
				ActorID:     actorResult.ActorID,
				Role:        actorResult.Role,
				Permissions: actorResult.Permissions,
			},
		}
		marshalledAssignments, err := json.Marshal(teamAssignments)
		if err != nil {
			slog.Error("Failed to marshal assignments", pendingTransfer.ToTeamLead, err)
			return err
		}

		resourceResult.Team = pendingTransfer.ToTeam
		resourceResult.TeamLead = pendingTransfer.ToTeamLead
		resourceResult.TeamAssignments = marshalledAssignments
		resourceResult.Journey = marshalledJourney
		resourceResult.PendingTransfer = nil

		return s.resourceRepo.UpdateResource(ctx, resourceResult)
	}

	return fmt.Errorf("resource %s has no pending transfer", id)
}

func (s *ServiceImpl) DeclineTransfer(ctx context.Context, id uuid.UUID) error {
	resourceResult, err := s.resourceRepo.Get(ctx, id)
	if err != nil {
		slog.Error("Failed to fetch resource", id, err)
		return err
	}
	if resourceResult == nil {
		return fmt.Errorf("resource %s not found", id)
	}

	resourceResult.PendingTransfer = nil

	return s.resourceRepo.UpdateResource(ctx, resourceResult)
}

func (s *ServiceImpl) UpdateResourceAssignment(ctx context.Context, actorId, resourceId uuid.UUID, add bool) error {
	resourceResult, err := s.resourceRepo.Get(ctx, resourceId)
	if err != nil {
		slog.Error("Failed to fetch resource", resourceId)
		return err
	}
	if resourceResult == nil {
		return fmt.Errorf("resource %s not found", resourceId)
	}

	var assignments []Assignment
	if err := json.Unmarshal(resourceResult.TeamAssignments, &assignments); err != nil {
		slog.Error("Failed to unmarshal assignments", resourceId, err)
		return err
	}

	actorResult, err := s.actorService.GetActor(actorId)
	if err != nil {
		slog.Error("Failed to fetch actor", actorId)
		return err
	}

	if add {
		assignments = append(assignments, Assignment{
			ActorID:     actorId,
			Role:        actorResult.Role,
			Permissions: actorResult.Permissions,
		})
	} else {
		for i, a := range assignments {
			if a.ActorID == actorId {
				assignments = append(assignments[:i], assignments[i+1:]...)
				break
			}
		}
	}

	marshalledAssignments, err := json.Marshal(assignments)
	if err != nil {
		slog.Error("Failed to marshal assignments", resourceId, err)
		return err
	}

	resourceResult.TeamAssignments = marshalledAssignments

	return s.resourceRepo.UpdateResource(ctx, resourceResult)
}

func (s *ServiceImpl) AddPermission(ctx context.Context, actorId, resourceId uuid.UUID, section, permission string) error {
	resourceResult, err := s.resourceRepo.Get(ctx, resourceId)
	if err != nil {
		slog.Error("Failed to fetch resource", resourceId)
		return err
	}
	if resourceResult == nil {
		return fmt.Errorf("resource %s not found", resourceId)
	}

	var teamAssignments []Assignment
	if err := json.Unmarshal(resourceResult.TeamAssignments, &teamAssignments); err != nil {
		slog.Error("Failed to unmarshal team assignments", resourceId, err)
		return err
	}

	for _, t := range teamAssignments {
		if t.ActorID == actorId {
			t.Permissions[section] = permission
			break
		}
	}

	marshalledAssignments, err := json.Marshal(teamAssignments)
	if err != nil {
		slog.Error("Failed to marshal team assignments", resourceId)
		return err
	}

	resourceResult.TeamAssignments = marshalledAssignments

	return s.resourceRepo.UpdateResource(ctx, resourceResult)
}

func (s *ServiceImpl) RemovePermission(ctx context.Context, actorId, resourceId uuid.UUID, section string) error {
	resourceResult, err := s.resourceRepo.Get(ctx, resourceId)
	if err != nil {
		slog.Error("Failed to fetch resource", resourceId)
		return err
	}
	if resourceResult == nil {
		return fmt.Errorf("resource %s not found", resourceId)
	}

	var teamAssignments []Assignment
	if err := json.Unmarshal(resourceResult.TeamAssignments, &teamAssignments); err != nil {
		slog.Error("Failed to unmarshal team assignments", resourceId, err)
		return err
	}

	for _, t := range teamAssignments {
		if t.ActorID == actorId {
			delete(t.Permissions, section)
			break
		}
	}

	marshalledAssignments, err := json.Marshal(teamAssignments)
	if err != nil {
		slog.Error("Failed to marshal team assignments", resourceId)
		return err
	}

	resourceResult.TeamAssignments = marshalledAssignments

	return s.resourceRepo.UpdateResource(ctx, resourceResult)
}

func (s *ServiceImpl) RequestResourceTransfer(ctx context.Context, resourceId uuid.UUID, toTeam string, toTeamLead uuid.UUID) error {
	resourceResult, err := s.resourceRepo.Get(ctx, resourceId)
	if err != nil {
		slog.Error("Failed to fetch resource", resourceId)
		return err
	}
	if resourceResult == nil {
		return fmt.Errorf("resource %s not found", resourceId)
	}

	if resourceResult.PendingTransfer != nil {
		return fmt.Errorf("resource %s has a pending transfer", resourceId)
	}

	_, err = s.actorService.GetActor(toTeamLead)
	if err != nil {
		slog.Error("Failed to fetch actor", toTeamLead)
		return err
	}

	pendingTransfer := JourneyStep{
		FromTeam:     resourceResult.Team,
		ToTeam:       toTeam,
		FromTeamLead: resourceResult.TeamLead,
		ToTeamLead:   toTeamLead,
	}

	marshalledPendingTransfer, err := json.Marshal(pendingTransfer)
	if err != nil {
		slog.Error("Failed to marshal pending transfer", resourceId)
		return err
	}

	resourceResult.PendingTransfer = marshalledPendingTransfer

	return s.resourceRepo.UpdateResource(ctx, resourceResult)
}

func (s *ServiceImpl) ArchiveResource(ctx context.Context, id uuid.UUID) error {
	return s.resourceRepo.UpdateArchived(ctx, id)
}
