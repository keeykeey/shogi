package repository

import (
	"fmt"
)

type KomaArrangements struct {
	ID         			uint16
	ArrangementID       uint16
	Arrangement         Arrangement
    KomaID              byte
	Koma                Koma
	PositionID          uint16
	Position            Position
	IsFirstMove         bool
	IsFront             bool
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