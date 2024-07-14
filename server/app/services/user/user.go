package user

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
	"hospital-system/server/utils"
	"strings"
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

func (service *ServiceImpl) CreateUser(user dto.User) error {
	id := uuid.New()

	// Per default, password is equal to the national identification number and can be changed later
	hashedPassword, err := utils.HashPassword(user.NationalIdentificationNumber)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to hash password: %v for user %s", err, user.Email))
		return err
	}

	username, err := extractUsernameFromEmail(user.Email)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to extract username from email: %v for user %s", err, user.Email))
		return err
	}

	if err = service.userRepo.Insert(models.User{
		ID:                           id,
		Firstname:                    user.Firstname,
		Lastname:                     user.Lastname,
		NationalIdentificationNumber: user.NationalIdentificationNumber,
		// Per default, the username is the email, it can be changed later
		Username: username,
		Email:    user.Email,
		Password: hashedPassword,
	}); err != nil {
		return err
	}

	_, err = service.authorizationClient.AddActor(context.Background(), &authorization.AddActorRequest{
		ActorId: id.String(),
		Role:    user.Role,
		Team:    user.Team,
	})
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to add user %v to authorization database: %v", id, err))

		if deleteErr := service.userRepo.Delete(id); err != nil {
			slog.Error(fmt.Sprintf("Failed to delete user %v: %v", id, err))
			return deleteErr
		}

		return err
	}

	return nil
}

func (service *ServiceImpl) GetUser(id uuid.UUID) (*dto.User, error) {
	user, err := service.userRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return &dto.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}, nil
}

func (service *ServiceImpl) GetByUsername(username string) (*dto.User, error) {
	user, err := service.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	actorResponse, err := service.authorizationClient.GetActor(context.Background(), &authorization.GetActorRequest{
		ActorId: user.ID.String(),
	})
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get actor %v from authorization database: %v", user.ID, err))
		return nil, err
	}

	return &dto.User{
		ID:                           user.ID,
		Firstname:                    user.Firstname,
		Lastname:                     user.Lastname,
		NationalIdentificationNumber: user.NationalIdentificationNumber,
		Username:                     user.Username,
		Email:                        user.Email,
		Role:                         actorResponse.GetActor().Role,
		Team:                         actorResponse.GetActor().Team,
		Permissions:                  actorResponse.GetActor().Permissions,
	}, nil
}

func (service *ServiceImpl) ValidateUserPassword(userId uuid.UUID, password string) (bool, error) {
	user, err := service.userRepo.Get(userId)
	if err != nil {
		return false, err
	}

	return utils.CheckPasswordHash(password, user.Password), nil
}

func (service *ServiceImpl) GetUsers() ([]dto.User, error) {
	users, err := service.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	getActorsResponse, err := service.authorizationClient.GetActors(context.Background(), &authorization.GetActorsRequest{})
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get actors from authorization database: %v", err))
		return nil, err
	}

	userList := make([]dto.User, 0)
	for _, user := range users {
		actor, found := lo.Find(getActorsResponse.GetActors(), func(actor *authorization.Actor) bool {
			return actor.ActorId == user.ID.String()
		})
		if !found {
			slog.Error(fmt.Sprintf("Actor not found for user %v, there are inconsistencies between server and authorization dbs", user.ID))
			return nil, fmt.Errorf("inconsistencies found between server and authorization databases")
		}

		newUser := dto.User{
			ID:                           user.ID,
			Firstname:                    user.Firstname,
			Lastname:                     user.Lastname,
			NationalIdentificationNumber: user.NationalIdentificationNumber,
			Username:                     user.Username,
			Email:                        user.Email,
			Role:                         actor.Role,
			Team:                         actor.Team,
			Permissions:                  actor.Permissions,
		}

		userList = append(userList, newUser)
	}

	return userList, nil
}

func extractUsernameFromEmail(email string) (string, error) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid email format")
	}
	return parts[0], nil
}
