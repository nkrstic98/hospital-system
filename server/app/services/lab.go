package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"hospital-system/server/app/constants"
	"hospital-system/server/app/dto"
	"hospital-system/server/models"
	"math"
	"math/rand"
	"time"
)

type labRepo interface {
	GetLabs(ctx context.Context) ([]models.Lab, error)
	GetLab(ctx context.Context, id uuid.UUID) (*models.Lab, error)
	UpdateLab(ctx context.Context, lab *models.Lab) error
}

type LabService struct {
	log  *zap.Logger
	repo labRepo
}

func NewLabService(log *zap.Logger, repo labRepo) *LabService {
	return &LabService{
		log:  log,
		repo: repo,
	}
}

func (s *LabService) GetLabs(ctx context.Context) ([]dto.Lab, error) {
	labs, err := s.repo.GetLabs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get labs: %w", err)
	}

	result := make([]dto.Lab, 0, len(labs))
	for _, lab := range labs {
		var testResults []dto.LabTest
		if lab.TestResults != nil {
			if stdErr := json.Unmarshal(lab.TestResults, &testResults); stdErr != nil {
				return nil, fmt.Errorf("failed to unmarshal lab results for lab %v: %w", lab.ID, stdErr)
			}
		}

		result = append(result, dto.Lab{
			ID:          lab.ID,
			RequestedAt: lab.RequestedAt,
			ProcessedAt: lo.Ternary(lab.ProcessedAt.Valid, &lab.ProcessedAt.Time, nil),
			TestType:    lab.TestType,
			TestResults: &testResults,
			RequestedBy: lab.RequestedBy,
			ProcessedBy: lo.Ternary(lab.ProcessedAt.Valid, lab.ProcessedBy, nil),
		})
	}

	return result, nil
}

func (s *LabService) ProcessLabTest(ctx context.Context, id, userId uuid.UUID) error {
	lab, err := s.repo.GetLab(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get lab: %w", err)
	}
	if lab == nil {
		return fmt.Errorf("lab not found")
	}

	labTest, found := constants.LabTests[lab.TestType]
	if !found {
		return fmt.Errorf("lab test not found")
	}

	results := make([]dto.LabTest, 0, len(labTest))
	for _, lt := range labTest {
		results = append(results, dto.LabTest{
			Name:           lt.Name,
			Unit:           lt.Unit,
			MinValue:       lt.MinValue,
			MaxValue:       lt.MaxValue,
			ReferenceRange: lt.ReferenceRange,
			Result:         generateRandomResult(lt.MinValue, lt.MaxValue),
		})
	}

	marshalledResults, stdErr := json.Marshal(results)
	if stdErr != nil {
		return fmt.Errorf("failed to marshal lab results: %w", stdErr)
	}

	lab.TestResults = marshalledResults
	lab.ProcessedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	lab.ProcessedBy = &userId

	if err = s.repo.UpdateLab(ctx, lab); err != nil {
		return fmt.Errorf("failed to update lab: %w", err)
	}

	return nil
}

func generateRandomResult(min, max float64) *float64 {
	var result float64

	randValue := rand.Float64()
	if randValue < 0.2 {
		// 20% chance to be below the minimum
		result = min * (0.8 + rand.Float64()*0.2)
	} else if randValue < 0.6 {
		// 40% chance to be above the maximum
		result = max + (max-min)*0.2*rand.Float64()
	} else {
		// 40% chance to be within the normal range
		result = min + rand.Float64()*(max-min)
	}

	// Round to 2 decimal places
	result = math.Round(result*100) / 100

	return lo.ToPtr(result)
}
