package analytics

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Event struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
    VisitID   uuid.UUID `gorm:"type:uuid;not null;index"`
    EventType string    `gorm:"type:varchar(255);not null"`
    // CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (event *Event) BeforeCreate(tx *gorm.DB) (err error) {
    if event.ID == uuid.Nil {
        event.ID = uuid.New()
    }
    return
}
