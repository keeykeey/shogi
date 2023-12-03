package repository

import (
	"gorm.io/gorm"
)

type Arrangement struct {
	ID                  uint16
	Name                string
	gorm.Model
}
