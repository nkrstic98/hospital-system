package team

import (
	"github.com/samber/lo"
	team_repo "hospital-system/authorization/app/repositories/team"
	"hospital-system/authorization/app/services/actor"
)

type ServiceImpl struct {
	teamRepo     team_repo.Repository
	actorService actor.Service
}

func NewService(teamRepo team_repo.Repository, actorService actor.Service) *ServiceImpl {
	return &ServiceImpl{
		teamRepo:     teamRepo,
		actorService: actorService,
	}
}

func (service *ServiceImpl) GetTeam(team string) (Team, error) {
	result, err := service.teamRepo.Get(team)
	if err != nil {
		return Team{}, err
	}

	actors, err := service.actorService.GetActorsByTeamID(team)
	if err != nil {
		return Team{}, err
	}

	return Team{
		Name:        result.ID,
		DisplayName: result.Name,
		Actors: lo.Map(actors, func(a actor.Actor, _ int) actor.Actor {
			return actor.Actor{
				ActorID:     a.ActorID,
				Role:        a.Role,
				Team:        a.Team,
				Permissions: a.Permissions,
			}
		}),
	}, nil
}

func (service *ServiceImpl) GetTeams() ([]Team, error) {
	teams, err := service.teamRepo.GetAll()
	if err != nil {
		return nil, err
	}

	resultTeams := make([]Team, 0)

	for _, team := range teams {
		actors, err := service.actorService.GetActorsByTeamID(team.ID)
		if err != nil {
			return nil, err
		}

		resultTeams = append(resultTeams, Team{
			Name:        team.ID,
			DisplayName: team.Name,
			Actors:      actors,
		})
	}

	return resultTeams, nil
}
