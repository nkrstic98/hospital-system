package actor

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hospital-system/authorization/models"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

func (repo *RepositoryImpl) Insert(actor models.Actor) error {
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Create(&actor)
		if result.Error != nil {
			slog.Error("Error creating actor: ", result.Error.Error())
			return result.Error
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *RepositoryImpl) Get(id uuid.UUID) (models.Actor, error) {
	var actor models.Actor
	result := repo.db.Clauses(clause.Locking{
		Strength: clause.LockingStrengthShare,
	}).Where("id = ?", id.String()).First(&actor)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("Error fetching actor with id %s: %s", id.String(), result.Error.Error()))
		return models.Actor{}, result.Error
	}

	return actor, nil
}

func (repo *RepositoryImpl) GetAll() ([]models.Actor, error) {
	var actors []models.Actor
	result := repo.db.Find(&actors)
	if result.Error != nil {
		slog.Error("Error fetching actors: ", result.Error.Error())
		return nil, result.Error
	}

	return actors, nil
}

func (repo *RepositoryImpl) GetByTeamID(teamID uint) ([]models.Actor, error) {
	var actors []models.Actor
	result := repo.db.Where("team_id = ?", teamID).Find(&actors)
	if result.Error != nil {
		return nil, result.Error
	}

	return actors, nil
}
