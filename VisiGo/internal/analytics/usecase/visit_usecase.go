package usecase

import (
	analytics "VisiGo/internal/analytics/entity"
	"VisiGo/internal/analytics/repository"
	"context"
	"errors"
)

type VisitUseCase interface {
	CreateVisit(ctx context.Context, visit *analytics.Visit) error
	FindVisitByID(visitID string) (*analytics.Visit, error)
}

type VisitInteraction struct {
	VisitRepository repository.VisitRepository
}

func (v *VisitInteraction) CreateVisit(ctx context.Context, visit *analytics.Visit) error {
	// Burada bazı validasyonlar veya iş mantıkları eklenebilir
	if visit.VisitorID == "" {
		return errors.New("visitor ID cannot be empty")
	}
	err := v.VisitRepository.CreateVisit(visit)
	if err != nil {
		return err
	}
	return nil
}

func (v *VisitInteraction) FindVisitByID(visitID string) (*analytics.Visit, error) {
	visit, err := v.VisitRepository.FindVisitByID(visitID)
	if err != nil {
		return nil, err
	}
	return visit, nil
}

func VisitUseC(repo repository.VisitRepository) VisitUseCase {
	return &VisitInteraction{
		VisitRepository: repo,
	}
}
