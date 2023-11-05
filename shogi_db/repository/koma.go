package repository

import (
	"gorm.io/gorm"
)

// Model Koma
type Koma struct {
	Id                 uint8
	MoveId             uint8
	MoveId2            uint8
	Name               string
	Name2              string
	gorm.Model
}

func ExportKoma() []Koma {
	komas := []Koma{
		Koma{Id: 1, MoveId: 1, MoveId2: 1, Name: "ou", Name2: ""},
		Koma{Id: 2, MoveId: 2, MoveId2: 2, Name: "kin", Name2: ""},
		Koma{Id: 3, MoveId: 3, MoveId2: 3, Name: "gin", Name2: "narigin"},
		Koma{Id: 4, MoveId: 4, MoveId2: 4, Name: "kei", Name2: "narikei"},
		Koma{Id: 5, MoveId: 5, MoveId2: 5, Name: "kyou", Name2: "narikyou"},
		Koma{Id: 6, MoveId: 6, MoveId2: 6, Name: "hu", Name2: "to"},
		Koma{Id: 7, MoveId: 7, MoveId2: 7, Name: "hisha", Name2: "ryuu"},
		Koma{Id: 8, MoveId: 8, MoveId2: 8, Name: "kaku", Name2: "uma"},
	}
	return komas
}