package repository

import (
	"time"
	"gorm.io/gorm"
)

// Model Arrangement
type BasicColumn struct {
	CreatedAt           time.Time          `json:"createdAt"`
	UpdatedAt           time.Time          `json:"updatedAt"`
	DeletedAt           gorm.DeletedAt     `json:"deletedAt"`
	
}