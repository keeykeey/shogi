package repository

import (
	"time"
	"gorm.io/gorm"
)

// Model Arrangement
type BasicColumn struct {
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}