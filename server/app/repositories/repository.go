package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hospital-system/server/models"
)

type Repository interface {
	// Users
	InsertUser(user models.User) error
	GetUser(id uuid.UUID) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByIDs(ids []uuid.UUID) ([]models.User, error)
	GetAllUsers() ([]models.User, error)
	DeleteUser(id uuid.UUID) error

	// Patients
	InsertPatient(user models.Patient) (*models.Patient, error)
	GetPatient(id uuid.UUID) (*models.Patient, error)
	GetPatientByPersonalID(personalID string) (*models.Patient, error)

	// Admissions
	InsertAdmission(admission models.Admission) (*models.Admission, error)
	GetAdmission(id uuid.UUID) (*models.Admission, error)
	DeleteAdmission(id uuid.UUID) error
	GetAdmissionsByPatientId(id uuid.UUID) ([]models.Admission, error)
	GetAdmissionsByStatuses(statuses []string) ([]models.Admission, error)
	GetAdmissionsByIDs(ids []uuid.UUID) ([]models.Admission, error)

	// Labs
	GetLabsByAdmissionID(admissionID uuid.UUID) ([]models.Lab, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{db: db}
}
