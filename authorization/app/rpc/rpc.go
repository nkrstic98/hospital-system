package rpc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"hospital-system/authorization/app/services/actor"
	"hospital-system/authorization/app/services/resource"
	"hospital-system/authorization/app/services/team"
	"hospital-system/proto_gen/authorization/v1"
)

type Service struct {
	authorization.AuthorizationServiceServer
	actorService    actor.Service
	teamService     team.Service
	resourceService resource.Service
}

func NewService(actorService actor.Service, teamService team.Service, service resource.Service) *Service {
	return &Service{
		actorService:    actorService,
		teamService:     teamService,
		resourceService: service,
	}
}

func (s *Service) AddActor(_ context.Context, request *authorization.AddActorRequest) (*authorization.AddActorResponse, error) {
	if err := s.actorService.AddActor(actor.Actor{
		ActorID: uuid.MustParse(request.ActorId),
		Role:    request.Role,
		Team:    request.Team,
	}); err != nil {
		return nil, err
	}

	return &authorization.AddActorResponse{}, nil
}

func (s *Service) GetActor(_ context.Context, request *authorization.GetActorRequest) (*authorization.GetActorResponse, error) {
	actor, err := s.actorService.GetActor(uuid.MustParse(request.ActorId))
	if err != nil {
		return nil, err
	}

	return &authorization.GetActorResponse{
		Actor: &authorization.Actor{
			ActorId:     actor.ActorID.String(),
			Role:        actor.Role,
			Team:        actor.Team,
			Permissions: actor.Permissions,
		},
	}, nil
}

func (s *Service) GetActors(_ context.Context, _ *authorization.GetActorsRequest) (*authorization.GetActorsResponse, error) {
	actors, err := s.actorService.GetActors()
	if err != nil {
		return nil, err
	}

	return &authorization.GetActorsResponse{
		Actors: lo.Map(actors, func(actor actor.Actor, _ int) *authorization.Actor {
			return &authorization.Actor{
				ActorId:     actor.ActorID.String(),
				Role:        actor.Role,
				Team:        actor.Team,
				Permissions: actor.Permissions,
			}
		}),
	}, nil
}

func (s *Service) GetTeams(_ context.Context, _ *authorization.GetTeamsRequest) (*authorization.GetTeamsResponse, error) {
	teams, err := s.teamService.GetTeams()
	if err != nil {
		return nil, err
	}

	return &authorization.GetTeamsResponse{
		Teams: lo.Map(teams, func(team team.Team, _ int) *authorization.Team {
			return &authorization.Team{
				Name:        team.Name,
				DisplayName: team.DisplayName,
				Actors: lo.Map(team.Actors, func(actor actor.Actor, _ int) *authorization.Actor {
					return &authorization.Actor{
						ActorId:     actor.ActorID.String(),
						Role:        actor.Role,
						Team:        actor.Team,
						Permissions: actor.Permissions,
					}
				}),
			}
		}),
	}, nil
}

func (s *Service) AddResource(_ context.Context, request *authorization.AddResourceRequest) (*authorization.AddResourceResponse, error) {
	var pendingTransfer *resource.JourneyStep
	if request.PendingTransfer != nil {
		pendingTransfer = &resource.JourneyStep{
			FromTeam:     request.PendingTransfer.FromTeam,
			ToTeam:       request.PendingTransfer.ToTeam,
			FromTeamLead: uuid.MustParse(request.PendingTransfer.FromTeamLead),
			ToTeamLead:   uuid.MustParse(request.PendingTransfer.ToTeamLead),
		}
	}

	if err := s.resourceService.AddResource(resource.Resource{
		ID:              uuid.MustParse(request.Id),
		Team:            request.Team,
		TeamLead:        uuid.MustParse(request.TeamLead),
		PendingTransfer: pendingTransfer,
	}); err != nil {
		return nil, err
	}

	return &authorization.AddResourceResponse{}, nil
}

func (s *Service) GetResources(ctx context.Context, request *authorization.GetResourcesRequest) (*authorization.GetResourcesResponse, error) {
	actorIdString := lo.FromPtrOr(request.ActorId, "")
	var actorIdPtr *uuid.UUID
	if actorIdString != "" {
		actorId := uuid.MustParse(actorIdString)
		actorIdPtr = &actorId
	}

	resources, err := s.resourceService.GetResources(
		ctx,
		lo.Ternary(len(request.Ids) > 0, &request.Ids, nil),
		actorIdPtr,
		request.Archived,
	)
	if err != nil {
		return nil, err
	}

	var result []*authorization.Resource
	for _, r := range resources {
		teamResult, err := s.teamService.GetTeam(r.Team)
		if err != nil {
			slog.Error("Failed to get team ", r.Team, err)
			return nil, err
		}

		var pendingTransfer *authorization.JourneyStep
		if r.PendingTransfer != nil {
			pendingTransfer = &authorization.JourneyStep{
				FromTeam:     r.PendingTransfer.FromTeam,
				ToTeam:       r.PendingTransfer.ToTeam,
				FromTeamLead: r.PendingTransfer.FromTeamLead.String(),
				ToTeamLead:   r.PendingTransfer.ToTeamLead.String(),
			}
		}

		result = append(result, &authorization.Resource{
			Id: r.ID.String(),
			Team: &authorization.Team{
				Name:        teamResult.Name,
				DisplayName: teamResult.DisplayName,
				Actors: lo.Map(teamResult.Actors, func(actor actor.Actor, _ int) *authorization.Actor {
					return &authorization.Actor{
						ActorId:     actor.ActorID.String(),
						Role:        actor.Role,
						Team:        actor.Team,
						Permissions: actor.Permissions,
					}
				}),
			},
			TeamLead:        r.TeamLead.String(),
			PendingTransfer: pendingTransfer,
		})
	}

	return &authorization.GetResourcesResponse{
		Resources: result,
	}, nil
}

func (s *Service) GetResource(ctx context.Context, request *authorization.GetResourceRequest) (*authorization.GetResourceResponse, error) {
	resourceResult, err := s.resourceService.GetResource(ctx, uuid.MustParse(request.Id))
	if err != nil {
		return nil, err
	}

	teamResult, err := s.teamService.GetTeam(resourceResult.Team)
	if err != nil {
		slog.Error("Failed to get team ", resourceResult.Team, err)
		return nil, err
	}

	return &authorization.GetResourceResponse{
		Resource: &authorization.Resource{
			Id: resourceResult.ID.String(),
			Team: &authorization.Team{
				Name:        teamResult.Name,
				DisplayName: teamResult.DisplayName,
			},
			TeamLead: resourceResult.TeamLead.String(),
			Assignments: lo.Map(resourceResult.Assignments, func(a resource.Assignment, _ int) *authorization.Assignment {
				return &authorization.Assignment{
					ActorId:     a.ActorID.String(),
					Role:        a.Role,
					Permissions: a.Permissions,
				}
			}),
			Journey: lo.Map(resourceResult.Journey, func(j resource.JourneyStep, _ int) *authorization.JourneyStep {
				return &authorization.JourneyStep{
					TransferTime: j.TransferTime,
					FromTeam:     j.FromTeam,
					ToTeam:       j.ToTeam,
					FromTeamLead: j.FromTeamLead.String(),
					ToTeamLead:   j.ToTeamLead.String(),
				}
			}),
			PendingTransfer: nil,
		},
	}, nil
}

func (s *Service) TransferResource(ctx context.Context, request *authorization.TransferResourceRequest) (*authorization.TransferResourceResponse, error) {
	resourceResult, err := s.resourceService.GetResource(ctx, uuid.MustParse(request.Id))
	if err != nil {
		slog.Error("Failed to get resource ", request.Id, err)
		return nil, err
	}

	if resourceResult.PendingTransfer == nil {
		return nil, fmt.Errorf("no pending transfer for resource %v", request.Id)
	}

	if resourceResult.PendingTransfer.ToTeamLead != uuid.MustParse(request.ActorId) {
		return nil, fmt.Errorf("user %v is not authorized to accept transfer for resource %v", request.ActorId, request.Id)
	}

	if request.AcceptTransfer {
		if err = s.resourceService.TransferResource(ctx, uuid.MustParse(request.Id)); err != nil {
			return nil, err
		}
	} else {
		if err = s.resourceService.DeclineTransfer(ctx, uuid.MustParse(request.Id)); err != nil {
			return nil, err
		}
	}

	return &authorization.TransferResourceResponse{}, nil
}

func (s *Service) UpdateResourceAssignment(ctx context.Context, request *authorization.UpdateResourceAssignmentRequest) (*authorization.UpdateResourceAssignmentResponse, error) {
	if err := s.resourceService.UpdateResourceAssignment(ctx, uuid.MustParse(request.ActorId), uuid.MustParse(request.ResourceId), request.Add); err != nil {
		return nil, err
	}

	return &authorization.UpdateResourceAssignmentResponse{}, nil
}

func (s *Service) AddPermission(ctx context.Context, request *authorization.AddPermissionRequest) (*authorization.AddPermissionResponse, error) {
	if err := s.resourceService.AddPermission(ctx, uuid.MustParse(request.ActorId), uuid.MustParse(request.ResourceId), request.Section, request.Permission); err != nil {
		return nil, err
	}

	return &authorization.AddPermissionResponse{}, nil
}

func (s *Service) RemovePermission(ctx context.Context, request *authorization.RemovePermissionRequest) (*authorization.RemovePermissionResponse, error) {
	if err := s.resourceService.RemovePermission(ctx, uuid.MustParse(request.ActorId), uuid.MustParse(request.ResourceId), request.Section); err != nil {
		return nil, err
	}

	return &authorization.RemovePermissionResponse{}, nil
}

func (s *Service) RequestResourceTransfer(ctx context.Context, request *authorization.RequestResourceTransferRequest) (*authorization.RequestResourceTransferResponse, error) {
	if err := s.resourceService.RequestResourceTransfer(ctx, uuid.MustParse(request.ResourceId), request.ToTeam, uuid.MustParse(request.ToTeamLead)); err != nil {
		return nil, err
	}

	return &authorization.RequestResourceTransferResponse{}, nil
}

func (s *Service) ArchiveResource(ctx context.Context, request *authorization.ArchiveResourceRequest) (*authorization.ArchiveResourceResponse, error) {
	if err := s.resourceService.ArchiveResource(ctx, uuid.MustParse(request.Id)); err != nil {
		return nil, err
	}

	return &authorization.ArchiveResourceResponse{}, nil
}
