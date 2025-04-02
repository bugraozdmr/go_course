package repository

import (
	analytics "VisiGo/internal/analytics/entity"

	"gorm.io/gorm"
)

type OneShot struct {
	VisitorID string
	EventType string
	PageURL   string
	UserAgent string
	Referer   *string
}

type OneShotRepository interface {
	CreateAll(oShot *OneShot) error
}

type OneShotDataBaseInteraction struct {
	DB *gorm.DB
}

// CreateAll method to create Visit, PageView, and Event
func (o *OneShotDataBaseInteraction) CreateAll(oShot *OneShot) error {
	// Start a transaction to ensure all records are created successfully
	tx := o.DB.Begin()

	// Step 1: Create Visit
	visit := &analytics.Visit{
		VisitorID: oShot.VisitorID,
		UserAgent: oShot.UserAgent,
		Referer:   oShot.Referer,
	}

	if err := tx.Create(&visit).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Step 2: Create PageView
	pageView := &analytics.PageView{
		VisitID: visit.ID,
		PageURL: oShot.PageURL,
	}

	if err := tx.Create(&pageView).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Step 3: Create Event
	event := &analytics.Event{
		VisitID:   visit.ID,
		EventType: oShot.EventType,
	}

	if err := tx.Create(&event).Error; err != nil {
		tx.Rollback()
		return err
	}

	// If all are created successfully, commit the transaction
	tx.Commit()
	return nil
}

func NewOneShotRepository(db *gorm.DB) OneShotRepository {
	return &OneShotDataBaseInteraction{
		DB: db,
	}
}