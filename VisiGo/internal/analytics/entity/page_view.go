package analytics

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type PageView struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
    VisitID   uuid.UUID `gorm:"type:uuid;not null;index"`
    PageURL   string    `gorm:"type:text;not null"`
    // CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (pageView *PageView) BeforeCreate(tx *gorm.DB) (err error) {
    if pageView.ID == uuid.Nil {
        pageView.ID = uuid.New()
    }
    return
}
