package repository

import (
	analytics "VisiGo/internal/analytics/entity"

	"gorm.io/gorm"
)

type VisitRepository interface {
	CreateVisit(visit *analytics.Visit) error
	FindVisitByID(id string) (*analytics.Visit, error)
	FindVisitsByVisitorID(visitorID string) ([]analytics.Visit, error)
}

type VisitDataBaseInteraction struct {
	DB *gorm.DB
}

func (v *VisitDataBaseInteraction) CreateVisit(visit *analytics.Visit) error {
	result := v.DB.Create(&visit)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (v *VisitDataBaseInteraction) FindVisitByID(id string) (*analytics.Visit, error) {
	var visit *analytics.Visit
	if err := v.DB.Where("id = ?", id).First(&visit).Error; err != nil {
		return nil, err
	}
	return visit, nil
}

func (v *VisitDataBaseInteraction) FindVisitsByVisitorID(visitorID string) ([]analytics.Visit, error) {
	var visits []analytics.Visit
	if err := v.DB.Where("visitor_id = ?", visitorID).Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}

func NewVisitRepository(db *gorm.DB) VisitRepository {
	return &VisitDataBaseInteraction{
		DB: db,
	}
}
