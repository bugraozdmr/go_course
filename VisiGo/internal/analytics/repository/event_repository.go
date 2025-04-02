package repository

import (
	analytics "VisiGo/internal/analytics/entity"

	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(event *analytics.Event) error
	FindEventByID(id string) (*analytics.Event, error)
	FindEventsByVisitID(visitID string) ([]analytics.Event, error)
}

type EventDataBaseInteraction struct {
	DB *gorm.DB
}

func (e *EventDataBaseInteraction) CreateEvent(event *analytics.Event) error {
	result := e.DB.Create(&event)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (e *EventDataBaseInteraction) FindEventByID(id string) (*analytics.Event, error) {
	var event *analytics.Event
	if err := e.DB.Where("id = ?", id).First(&event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (e *EventDataBaseInteraction) FindEventsByVisitID(visitID string) ([]analytics.Event, error) {
	var events []analytics.Event
	if err := e.DB.Where("visit_id = ?", visitID).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &EventDataBaseInteraction{
		DB: db,
	}
}
