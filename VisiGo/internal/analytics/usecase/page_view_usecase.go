package usecase

import (
	analytics "VisiGo/internal/analytics/entity"
	"VisiGo/internal/analytics/repository"
	"context"
	"errors"
)

type PageViewUseCase interface {
	CreatePageView(ctx context.Context, pageView *analytics.PageView) error
	FindPageViewByID(pageViewID string) (*analytics.PageView, error)
}

type PageViewInteraction struct {
	PageViewRepository repository.PageViewRepository
}

func (p *PageViewInteraction) CreatePageView(ctx context.Context, pageView *analytics.PageView) error {
	// Burada bazı validasyonlar veya iş mantıkları eklenebilir
	if pageView.PageURL == "" {
		return errors.New("page URL cannot be empty")
	}
	err := p.PageViewRepository.CreatePageView(pageView)
	if err != nil {
		return err
	}
	return nil
}

func (p *PageViewInteraction) FindPageViewByID(pageViewID string) (*analytics.PageView, error) {
	pageView, err := p.PageViewRepository.FindPageViewByID(pageViewID)
	if err != nil {
		return nil, err
	}
	return pageView, nil
}

func PageViewUseC(repo repository.PageViewRepository) PageViewUseCase {
	return &PageViewInteraction{
		PageViewRepository: repo,
	}
}
