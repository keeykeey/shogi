package repository

import (
	"fmt"
)

type KomaArrangements struct {
	ID         			uint16           `gorm:"primaryKey" json:"id"`
	ArrangementID       uint16           `json:"arrangementId"`
	Arrangement         Arrangement      `json:"arrangement"`
    KomaID              byte             `json:"komaId"`
	Koma                Koma             `json:"koma"`
	PositionID          uint16           `json:"positionId"`
	Position            Position         `json:"position"`
	IsFirstMove         bool             `json:"isFirstMove"`
	IsFront             bool             `json:"isFront"`
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
	fmt.Printf("\n\nlog: %s",s)
	return komaArrangements;
}