package models

import (
	"time"

	"gorm.io/gorm"
)

// Base model dengan timestamp dan soft delete
type BaseModel struct {
	CreatedAt time.Time      `gorm:"autoCreateTime"` // Timestamp saat insert
	UpdatedAt time.Time      `gorm:"autoUpdateTime"` // Timestamp saat update
	DeletedAt gorm.DeletedAt `gorm:"index"`          // Soft Delete
}
