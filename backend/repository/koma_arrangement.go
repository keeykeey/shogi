package repository

import (
	"fmt"
)

type KomaArrangements struct {
	ID         			uint16           `gorm:"primaryKey" json:"id"`
	ArrangementID       uint16           `json:"arrangementId"`
	Arrangement         arrangement      `json:"arrangement"`
    KomaID              byte             `json:"komaId"`
	Koma                koma             `json:"koma"`
	PositionID          uint16           `json:"positionId"`
	Position            position         `json:"position"`
	IsFirstMove         bool             `json:"isFirstMove"`
	IsFront             bool             `json:"isFront"`
}

type arrangement struct {
	ID                  uint16     `gorm:"primaryKey" json:"id"`
	Name                string     `json:"name"`
}

type koma struct {
	ID                 byte           `gorm:"primaryKey" json:"id"`
	MoveID             byte           `json:"moveId"`
	MoveID2            byte           `json:"moveId2"`
	Name               string         `json:"name"`
	Name2              string         `json:"name2"`
}

type position struct {
	ID               uint16      `gorm:"primaryKey" json:"id"`
	Width            byte        `json:"width"`
	Height           byte        `json:"height"`
	Name             string      `json:"name"`
}

func GetKomaArrangements(arrangementId uint16) ([]KomaArrangements) {
    var komaArrangements []KomaArrangements
	db := GetDbConnection()
	db.Debug().
	  Table("koma_arrangements").
	  Preload("Arrangement").
	  Preload("Koma").
	  Preload("Position").
	  Where("arrangement_id = ?", arrangementId).
	  Joins("join arrangements on koma_arrangements.arrangement_id = arrangements.id").
	  Joins("join komas on koma_arrangements.koma_id = komas.id").
	  Joins("join positions on koma_arrangements.position_id = positions.id").
	  Find(&komaArrangements)
	s := fmt.Sprintf("%v",komaArrangements);
	return komaArrangements;
}