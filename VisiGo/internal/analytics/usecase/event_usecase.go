package usecase

import (
	analytics "VisiGo/internal/analytics/entity"
	"VisiGo/internal/analytics/repository"
	"context"
	"errors"
)

type EventUseCase interface {
	CreateEvent(ctx context.Context, event *analytics.Event) error
	FindEventByID(eventID string) (*analytics.Event, error)
}

type EventInteraction struct {
	EventRepository repository.EventRepository
}

func (e *EventInteraction) CreateEvent(ctx context.Context, event *analytics.Event) error {
	// Burada bazı validasyonlar veya iş mantıkları eklenebilir
	if event.EventType == "" {
		return errors.New("event type cannot be empty")
	}
	err := e.EventRepository.CreateEvent(event)
	if err != nil {
		return err
	}
	return nil
}

func (e *EventInteraction) FindEventByID(eventID string) (*analytics.Event, error) {
	event, err := e.EventRepository.FindEventByID(eventID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func EventUseC(repo repository.EventRepository) EventUseCase {
	return &EventInteraction{
		EventRepository: repo,
	}
}
