package repository

// Model Koma
type Koma struct {
	ID                 byte           `gorm:"primaryKey"`
	MoveID             byte
	MoveID2            byte
	Name               string
	Name2              string
	BasicColumn
}

func ExportKoma() []Koma {
	komas := []Koma{
		Koma{ID: 1, MoveID: 1, MoveID2: 1, Name: "ou", Name2: ""},
		Koma{ID: 2, MoveID: 2, MoveID2: 2, Name: "kin", Name2: ""},
		Koma{ID: 3, MoveID: 3, MoveID2: 3, Name: "gin", Name2: "narigin"},
		Koma{ID: 4, MoveID: 4, MoveID2: 4, Name: "kei", Name2: "narikei"},
		Koma{ID: 5, MoveID: 5, MoveID2: 5, Name: "kyou", Name2: "narikyou"},
		Koma{ID: 6, MoveID: 6, MoveID2: 6, Name: "hu", Name2: "to"},
		Koma{ID: 7, MoveID: 7, MoveID2: 7, Name: "hisha", Name2: "ryuu"},
		Koma{ID: 8, MoveID: 8, MoveID2: 8, Name: "kaku", Name2: "uma"},
	}
	return komas
}

func ExportKomaForTest() []Koma {
	komas := []Koma{
		Koma{ID: 1, MoveID: 1, MoveID2: 1, Name: "ou", Name2: ""},
		Koma{ID: 2, MoveID: 2, MoveID2: 2, Name: "kin", Name2: ""},
		Koma{ID: 3, MoveID: 3, MoveID2: 3, Name: "gin", Name2: "narigin"},
	}
	return komas
}