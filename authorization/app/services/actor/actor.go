package actor

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"hospital-system/authorization/app/repositories/actor"
	"hospital-system/authorization/app/repositories/role"
	"hospital-system/authorization/app/repositories/team"
	"hospital-system/authorization/models"
)

type ServiceImpl struct {
	actorRepo actor.Repository
	roleRepo  role.Repository
	teamRepo  team.Repository
}

func NewService(actorRepo actor.Repository, roleRepo role.Repository, teamRepo team.Repository) *ServiceImpl {
	return &ServiceImpl{
		actorRepo: actorRepo,
		roleRepo:  roleRepo,
		teamRepo:  teamRepo,
	}
}

func (service *ServiceImpl) AddActor(actor Actor) error {
	role, err := service.roleRepo.GetByName(actor.Role)
	if err != nil {
		slog.Error("Error fetching role: ", err)
		return err
	}

	var teamID *uint
	if actor.Team != nil {
		team, err := service.teamRepo.GetByName(*actor.Team)
		if err != nil {
			slog.Error("Error fetching team: ", err)
			return err
		}
		teamID = &team.ID
	}

	if err = service.actorRepo.Insert(models.Actor{
		ID:     actor.ActorID,
		RoleID: role.ID,
		TeamID: teamID,
	}); err != nil {
		slog.Error("Error inserting actor: ", err)
		return err
	}

	return nil
}

func (service *ServiceImpl) GetActor(id uuid.UUID) (Actor, error) {
	actor, err := service.actorRepo.Get(id)
	if err != nil {
		slog.Error("Error fetching actor: ", err)
		return Actor{}, err
	}

	role, err := service.roleRepo.Get(actor.RoleID)
	if err != nil {
		slog.Error("Error fetching role: ", err)
		return Actor{}, err
	}

	var teamName *string
	if actor.TeamID != nil {
		team, err := service.teamRepo.Get(*actor.TeamID)
		if err != nil {
			slog.Error("Error fetching team: ", err)
			return Actor{}, err
		}

		teamName = lo.ToPtr(team.Name)
	}

	return Actor{
		ActorID: actor.ID,
		Role:    role.Name,
		Team:    teamName,
		//Permissions: lo.Map(role.Permissions, func(permission models.Permission, _ int) string {
		//	return permission.Name
		//}),
	}, nil
}

func (service *ServiceImpl) GetActors() ([]Actor, error) {
	actors, err := service.actorRepo.GetAll()
	if err != nil {
		return nil, err
	}

	roles, err := service.roleRepo.GetAll()
	if err != nil {
		return nil, err
	}

	teams, err := service.teamRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return lo.Map(actors, func(actor models.Actor, _ int) Actor {
		role, found := lo.Find(roles, func(role models.Role) bool {
			return role.ID == actor.RoleID
		})
		if !found {
			slog.Error("Role not found for actor ", actor.ID)
			return Actor{}
		}

		var team models.Team
		if actor.TeamID != nil {
			team, found = lo.Find(teams, func(team models.Team) bool {
				return team.ID == *actor.TeamID
			})
			if !found {
				slog.Error("Team not found for actor ", actor.ID)
				return Actor{}
			}
		}

		return Actor{
			ActorID:     actor.ID,
			Role:        role.DisplayName,
			Team:        lo.ToPtr(team.DisplayName),
			Permissions: []string{},
		}
	}), nil
}

func (service *ServiceImpl) GetActorsByTeamID(teamID uint) ([]Actor, error) {
	actors, err := service.actorRepo.GetByTeamID(teamID)
	if err != nil {
		return nil, err
	}

	roles, err := service.roleRepo.GetAll()

	return lo.Map(actors, func(actor models.Actor, _ int) Actor {
		role, found := lo.Find(roles, func(role models.Role) bool {
			return role.ID == actor.RoleID
		})
		if !found {
			slog.Error("Role not found for actor ", actor.ID)
			return Actor{}
		}

		return Actor{
			ActorID:     actor.ID,
			Role:        role.Name,
			Team:        nil,
			Permissions: []string{},
		}
	}), nil
}
