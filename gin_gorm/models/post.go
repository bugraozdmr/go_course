package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// gorm.Model -> createdAt deletedAt updatedAt ID(int autoincrement)

type Post struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"-"` // json:"-" hidden id
	Title     string
	Body      string
	CreatedAt time.Time // Automatically handled by GORM
	UpdatedAt time.Time // Automatically handled by GORM
}

// BeforeCreate, bir post oluşturulmadan önce çalışacak fonksiyon
func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	if post.ID == uuid.Nil {
		post.ID = uuid.New()
	}
	return
}
