package rpc

import (
	"context"
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
	if err := s.resourceService.AddResource(resource.Resource{
		ID:       uuid.MustParse(request.Id),
		Team:     request.Team,
		TeamLead: uuid.MustParse(request.TeamLead),
	}); err != nil {
		return nil, err
	}

	return &authorization.AddResourceResponse{}, nil
}

func (s *Service) GetResources(_ context.Context, request *authorization.GetResourcesRequest) (*authorization.GetResourcesResponse, error) {
	resources, err := s.resourceService.GetResources(lo.Ternary(len(request.Ids) > 0, &request.Ids, nil))
	if err != nil {
		return nil, err
	}

	var result []*authorization.Resource
	for _, resource := range resources {
		teamResult, err := s.teamService.GetTeam(resource.Team)
		if err != nil {
			slog.Error("Failed to get team ", resource.Team, err)
			return nil, err
		}

		result = append(result, &authorization.Resource{
			Id: resource.ID.String(),
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
			TeamLead: resource.TeamLead.String(),
		})
	}

	return &authorization.GetResourcesResponse{
		Resources: result,
	}, nil
}

func (s *Service) ArchiveResource(_ context.Context, request *authorization.ArchiveResourceRequest) (*authorization.ArchiveResourceResponse, error) {
	if err := s.resourceService.ArchiveResource(uuid.MustParse(request.Id)); err != nil {
		return nil, err
	}

	return &authorization.ArchiveResourceResponse{}, nil
}
