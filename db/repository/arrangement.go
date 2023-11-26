package repository

// Model Arrangement
type Arrangement struct {
	Id                  uint16    `gorm:"primaryKey"`
	Name                string
	BasicColumn
}

func ExportArrangement() []Arrangement {
	arrangements := []Arrangement{
		Arrangement{Id: 1, Name: "平手" },
		Arrangement{Id: 2, Name: "二枚落ち" },
	}
	return arrangements
}