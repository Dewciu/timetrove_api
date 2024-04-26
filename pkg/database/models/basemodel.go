package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID        uuid.UUID      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	b.ID = uuid.New()
	return nil
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}

func (b *BaseModel) BeforeDelete(tx *gorm.DB) error {
	b.DeletedAt = gorm.DeletedAt{
		Time:  time.Now(),
		Valid: true,
	}
	return nil
}
