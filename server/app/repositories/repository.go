package repositories

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hospital-system/server/models"
)

type Repository interface {
	// Users
	InsertUser(ctx context.Context, user models.User) (*uuid.UUID, error)
	GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUsersByIDs(ctx context.Context, ids []uuid.UUID) ([]models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error

	// Patients
	InsertPatient(ctx context.Context, user models.Patient) (*models.Patient, error)
	GetPatient(ctx context.Context, id uuid.UUID) (*models.Patient, error)
	GetPatientByPersonalID(ctx context.Context, personalID string) (*models.Patient, error)

	// Admissions
	InsertAdmission(ctx context.Context, admission models.Admission) (*models.Admission, error)
	GetAdmission(ctx context.Context, id uuid.UUID) (*models.Admission, error)
	DeleteAdmission(ctx context.Context, id uuid.UUID) error
	GetAdmissionsByPatientId(ctx context.Context, id uuid.UUID) ([]models.Admission, error)
	GetAdmissionsByStatuses(ctx context.Context, statuses []string) ([]models.Admission, error)
	GetAdmissionsByIDs(ctx context.Context, ids []uuid.UUID) ([]models.Admission, error)

	// Labs
	GetLabsByAdmissionID(ctx context.Context, admissionID uuid.UUID) ([]models.Lab, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{db: db}
}
