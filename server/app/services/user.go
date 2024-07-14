package services

import (
	"context"
	"fmt"
	"hospital-system/proto_gen/authorization/v1"
	"hospital-system/server/app/dto"
	"hospital-system/server/models"
	"hospital-system/server/utils"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type userRepo interface {
	InsertUser(ctx context.Context, user models.User) (*uuid.UUID, error)
	GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUsersByIDs(ctx context.Context, ids []uuid.UUID) ([]models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type UserService struct {
	log                 *zap.Logger
	authorizationClient authorization.AuthorizationServiceClient
	repo                userRepo
}

func NewUserService(authorizationClient authorization.AuthorizationServiceClient, userRepo userRepo) *UserService {
	return &UserService{
		authorizationClient: authorizationClient,
		repo:                userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user dto.User) error {
	// Per default, password is equal to the national identification number and can be changed later
	hashedPassword, err := utils.HashPassword(user.NationalIdentificationNumber)
	if err != nil {
		return fmt.Errorf("failed to hash user password: %w", err)
	}

	username, err := utils.ExtractUsernameFromEmail(user.Email)
	if err != nil {
		return fmt.Errorf("failed to extract username from email: %w", err)
	}

	userId, err := s.repo.InsertUser(ctx, models.User{
		Firstname:                    user.Firstname,
		Lastname:                     user.Lastname,
		NationalIdentificationNumber: user.NationalIdentificationNumber,
		// Per default, the username is the email, it can be changed later
		Username: username,
		Email:    user.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return fmt.Errorf("failed to insert user into database: %w", err)
	}
	if userId == nil {
		return fmt.Errorf("create user returned nil user id")
	}

	if _, err = s.authorizationClient.AddActor(ctx, &authorization.AddActorRequest{
		ActorId: userId.String(),
		Role:    user.Role,
		Team:    user.Team,
	}); err != nil {
		if userDeleteErr := s.repo.DeleteUser(ctx, *userId); userDeleteErr != nil {
			return fmt.Errorf("failed to delete user from database: %w", userDeleteErr)
		}

		return fmt.Errorf("failed to add user to authorization database: %w", err)
	}

	return nil
}

func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*dto.User, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from database: %w", err)
	}
	if user == nil {
		return nil, nil
	}

	return &dto.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}, nil
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (*dto.User, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from database: %w", err)
	}
	if user == nil {
		return nil, nil
	}

	actorResponse, err := s.authorizationClient.GetActor(ctx, &authorization.GetActorRequest{
		ActorId: user.ID.String(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get actor from authorization database: %w", err)
	}

	dtoUser := toDtoUser(*user, actorResponse.Actor)

	return &dtoUser, nil
}

func (s *UserService) ValidateUserPassword(ctx context.Context, userId uuid.UUID, password string) (bool, error) {
	user, err := s.repo.GetUser(ctx, userId)
	if err != nil {
		return false, fmt.Errorf("failed to get user from database: %w", err)
	}
	if user == nil {
		return false, fmt.Errorf("user not found in database: %v", userId)
	}

	return utils.CheckPasswordHash(password, user.Password), nil
}

func (s *UserService) GetUsers(ctx context.Context) ([]dto.User, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get users from database: %w", err)
	}
	if len(users) == 0 {
		return nil, nil
	}

	getActorsResponse, err := s.authorizationClient.GetActors(ctx, &authorization.GetActorsRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to get actors from authorization database: %w", err)
	}

	userList := make([]dto.User, 0)
	for _, u := range users {
		actor, found := lo.Find(getActorsResponse.GetActors(), func(actor *authorization.Actor) bool {
			return actor.ActorId == u.ID.String()
		})
		if !found {
			return nil, fmt.Errorf("inconsistencies found between server and authorization databases")
		}

		user := toDtoUser(u, actor)
		userList = append(userList, user)
	}

	return userList, nil
}

// TODO: Refactor this flow to be generic
func (s *UserService) GetDepartments(ctx context.Context, team, role *string) (map[string]dto.Department, error) {
	teams, err := s.authorizationClient.GetTeams(ctx, &authorization.GetTeamsRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to get teams: %w", err)
	}

	departments := make(map[string]dto.Department)
	for _, t := range teams.Teams {
		users, err := s.repo.GetUsersByIDs(ctx, lo.FilterMap(t.Actors, func(actor *authorization.Actor, _ int) (uuid.UUID, bool) {
			if role != nil && actor.Role != *role {
				return uuid.UUID{}, false
			}

			return uuid.MustParse(actor.ActorId), true
		}))
		if err != nil {
			return nil, fmt.Errorf("failed to get users by ids: %w", err)
		}

		departments[t.Name] = dto.Department{
			DisplayName: t.DisplayName,
			Users: lo.Map(users, func(p models.User, _ int) dto.User {
				return toDtoUser(p, nil)
			}),
		}
	}

	return departments, nil
}
