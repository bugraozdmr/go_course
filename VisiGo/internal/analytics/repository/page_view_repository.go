package repository

import (
	analytics "VisiGo/internal/analytics/entity"

	"gorm.io/gorm"
)

type PageViewRepository interface {
	CreatePageView(pageView *analytics.PageView) error
	FindPageViewByID(id string) (*analytics.PageView, error)
	FindPageViewsByVisitID(visitID string) ([]analytics.PageView, error)
}

type PageViewDataBaseInteraction struct {
	DB *gorm.DB
}

func (p *PageViewDataBaseInteraction) CreatePageView(pageView *analytics.PageView) error {
	result := p.DB.Create(&pageView)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *PageViewDataBaseInteraction) FindPageViewByID(id string) (*analytics.PageView, error) {
	var pageView *analytics.PageView
	if err := p.DB.Where("id = ?", id).First(&pageView).Error; err != nil {
		return nil, err
	}
	return pageView, nil
}

func (p *PageViewDataBaseInteraction) FindPageViewsByVisitID(visitID string) ([]analytics.PageView, error) {
	var pageViews []analytics.PageView
	if err := p.DB.Where("visit_id = ?", visitID).Find(&pageViews).Error; err != nil {
		return nil, err
	}
	return pageViews, nil
}

func NewPageViewRepository(db *gorm.DB) PageViewRepository {
	return &PageViewDataBaseInteraction{
		DB: db,
	}
}
