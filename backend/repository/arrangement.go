package repository

import (
	"gorm.io/gorm"
)

type Arrangement struct {
	Id                  uint16
	Name                string
	gorm.Model
}

func ExportArrangement() []Arrangement {
	arrangements := []Arrangement{
		Arrangement{Id: 1, Name: "平手" },
		Arrangement{Id: 2, Name: "二枚落ち" },
	}
	return arrangements
}
