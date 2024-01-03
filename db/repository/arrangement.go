package repository

// Model Arrangement
type Arrangement struct {
	ID                  uint16    `gorm:"primaryKey"`
	Name                string
	BasicColumn
}

func ExportArrangement() []Arrangement {
	arrangements := []Arrangement{
		Arrangement{ID: 1, Name: "平手" },
		Arrangement{ID: 2, Name: "二枚落ち" },
	}
	return arrangements
}

func ExportArrangementForTest() []Arrangement {
	arrangements := []Arrangement{
		Arrangement{ID: 1, Name: "平手" },
		Arrangement{ID: 2, Name: "二枚落ち" },
	}
	return arrangements
}