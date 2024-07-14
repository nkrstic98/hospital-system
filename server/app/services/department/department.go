package department

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"hospital-system/proto_gen/authorization/v1"
	"hospital-system/server/app/dto"
	"hospital-system/server/app/repositories/user"
	"hospital-system/server/models"
)

type ServiceImpl struct {
	authorizationClient authorization.AuthorizationServiceClient
	userRepo            user.Repository
}

func NewService(authorizationClient authorization.AuthorizationServiceClient, userRepo user.Repository) *ServiceImpl {
	return &ServiceImpl{
		authorizationClient: authorizationClient,
		userRepo:            userRepo,
	}
}

func (s *ServiceImpl) GetDepartments() (map[string]dto.Department, error) {
	teams, err := s.authorizationClient.GetTeams(context.Background(), &authorization.GetTeamsRequest{})
	if err != nil {
		slog.Error("Failed to get teams from authorization service", err)
		return nil, err
	}

	departments := make(map[string]dto.Department)
	for _, team := range teams.Teams {
		// Get physicians
		physicians, err := s.userRepo.GetByIDs(lo.FilterMap(team.Actors, func(actor *authorization.Actor, _ int) (uuid.UUID, bool) {
			if actor.Role != "ATTENDING" {
				return uuid.UUID{}, false
			}

			return uuid.MustParse(actor.ActorId), true
		}))
		if err != nil {
			slog.Error("Failed to get physicians by ids", err)
			return nil, err
		}

		// Get residents
		residents, err := s.userRepo.GetByIDs(lo.FilterMap(team.Actors, func(actor *authorization.Actor, _ int) (uuid.UUID, bool) {
			if actor.Role != "RESIDENT" {
				return uuid.UUID{}, false
			}

			return uuid.MustParse(actor.ActorId), true
		}))
		if err != nil {
			slog.Error("Failed to get residents by ids", err)
			return nil, err
		}

		// Get nurses
		nurses, err := s.userRepo.GetByIDs(lo.FilterMap(team.Actors, func(actor *authorization.Actor, _ int) (uuid.UUID, bool) {
			if actor.Role != "NURSE" {
				return uuid.UUID{}, false
			}

			return uuid.MustParse(actor.ActorId), true
		}))
		if err != nil {
			slog.Error("Failed to get nurses by ids", err)
			return nil, err
		}

		departments[team.Name] = dto.Department{
			DisplayName: team.DisplayName,
			Physicians: lo.Map(physicians, func(p models.User, _ int) dto.Employee {
				return dto.Employee{
					ID:       p.ID,
					FullName: fmt.Sprintf("Doctor %s %s, MD", p.Firstname, p.Lastname),
				}
			}),
			Residents: lo.Map(residents, func(r models.User, _ int) dto.Employee {
				return dto.Employee{
					ID:       r.ID,
					FullName: fmt.Sprintf("Doctor %s %s, MD", r.Firstname, r.Lastname),
				}
			}),
			Nurses: lo.Map(nurses, func(n models.User, _ int) dto.Employee {
				return dto.Employee{
					ID:       n.ID,
					FullName: fmt.Sprintf("Doctor %s %s, MD", n.Firstname, n.Lastname),
				}
			}),
		}
	}

	return departments, nil
}
