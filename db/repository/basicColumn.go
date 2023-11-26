package repository

import (
	"time"
)

// Model Arrangement
type BasicColumn struct {
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           time.Time
}