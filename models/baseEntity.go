package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
