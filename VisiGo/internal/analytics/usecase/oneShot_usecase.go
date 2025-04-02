package usecase

import (
	"VisiGo/internal/analytics/repository"
	"context"
	"errors"
)

type OneShotUseCase interface {
	CreateAll(ctx context.Context, OneShot *repository.OneShot) error
}

type OneShotInteraction struct {
	OneShotRepository repository.OneShotRepository
}

func (v *OneShotInteraction) CreateAll(ctx context.Context, OneShot *repository.OneShot) error {
	if OneShot.PageURL == "" {
		return errors.New("Page Url can not be empty")
	}
	if OneShot.VisitorID == "" {
		return errors.New("Visitor ID cannot be empty")
	}
	if OneShot.EventType == "" {
		return errors.New("Event Type cannot be empty")
	}
	if OneShot.UserAgent == "" {
		return errors.New("User Agent cannot be empty")
	}


	err := v.OneShotRepository.CreateAll(OneShot)
	if err != nil {
		return err
	}
	return nil
}

func OneShotUseC(repo repository.OneShotRepository) OneShotUseCase {
	return &OneShotInteraction{
		OneShotRepository: repo,
	}
}
