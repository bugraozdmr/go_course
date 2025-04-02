package analytics

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Visit struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VisitorID string    `gorm:"type:varchar(255);not null"`
	UserAgent string    `gorm:"type:text;not null"`
	Referer   *string   `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	PageViews []PageView `gorm:"foreignKey:VisitID"`
	Events    []Event    `gorm:"foreignKey:VisitID"`
}

func (visit *Visit) BeforeCreate(tx *gorm.DB) (err error) {
	if visit.ID == uuid.Nil {
		visit.ID = uuid.New()
	}
	return
}
