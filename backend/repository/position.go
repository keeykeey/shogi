package repository

import (
	"gorm.io/gorm"
)

type Position struct {
	positionBase        PositionBase
	gorm.Model
}

type PositionBase struct {
	Id               uint16
	Number           uint8
	Name             string
}
