package repository

// intermediate table between model Arrangement and model Koma
type KomaArrangement struct {
	ID                   uint16           `gorm:"primaryKey"`
	ArrangementID        uint16
	Arrangement          Arrangement
	KomaID               uint8
	Koma                 Koma
	PositionID           uint16
	Position             Position
	IsFirstMove          bool             // 先手か後手か
	IsFront              bool             // 駒が成っているか否か
	BasicColumn
}

func ExportKomaArrangement() []KomaArrangement {
	arrangements := ExportArrangement()
	komas := ExportKoma()
	positions := ExportPosition()
	komaArrangements := []KomaArrangement{
		KomaArrangement{ID:1 ,ArrangementID: arrangements[0].ID, KomaID: komas[0].ID, PositionID: positions[76].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:2 ,ArrangementID: arrangements[0].ID, KomaID: komas[1].ID, PositionID: positions[77].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:3 ,ArrangementID: arrangements[0].ID, KomaID: komas[1].ID, PositionID: positions[75].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:4 ,ArrangementID: arrangements[0].ID, KomaID: komas[2].ID, PositionID: positions[78].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:5 ,ArrangementID: arrangements[0].ID, KomaID: komas[2].ID, PositionID: positions[74].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:6 ,ArrangementID: arrangements[0].ID, KomaID: komas[3].ID, PositionID: positions[79].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:7 ,ArrangementID: arrangements[0].ID, KomaID: komas[3].ID, PositionID: positions[73].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:8 ,ArrangementID: arrangements[0].ID, KomaID: komas[4].ID, PositionID: positions[80].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:9 ,ArrangementID: arrangements[0].ID, KomaID: komas[4].ID, PositionID: positions[72].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:10 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[54].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:11 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[55].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:12 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[56].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:13 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[57].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:14 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[58].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:15 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[59].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:16 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[60].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:17 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[61].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:18 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[62].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:19 ,ArrangementID: arrangements[0].ID, KomaID: komas[6].ID, PositionID: positions[70].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:20 ,ArrangementID: arrangements[0].ID, KomaID: komas[7].ID, PositionID: positions[64].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:21 ,ArrangementID: arrangements[0].ID, KomaID: komas[0].ID, PositionID: positions[4].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:22 ,ArrangementID: arrangements[0].ID, KomaID: komas[1].ID, PositionID: positions[5].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:23 ,ArrangementID: arrangements[0].ID, KomaID: komas[1].ID, PositionID: positions[3].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:24 ,ArrangementID: arrangements[0].ID, KomaID: komas[2].ID, PositionID: positions[6].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:25 ,ArrangementID: arrangements[0].ID, KomaID: komas[2].ID, PositionID: positions[2].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:26 ,ArrangementID: arrangements[0].ID, KomaID: komas[3].ID, PositionID: positions[7].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:27 ,ArrangementID: arrangements[0].ID, KomaID: komas[3].ID, PositionID: positions[1].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:28 ,ArrangementID: arrangements[0].ID, KomaID: komas[4].ID, PositionID: positions[8].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:29 ,ArrangementID: arrangements[0].ID, KomaID: komas[4].ID, PositionID: positions[0].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:30 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[18].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:31 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[19].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:32 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[20].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:33 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[21].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:34 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[22].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:35 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[23].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:36 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[24].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:37 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[25].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:38 ,ArrangementID: arrangements[0].ID, KomaID: komas[5].ID, PositionID: positions[26].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:39 ,ArrangementID: arrangements[0].ID, KomaID: komas[6].ID, PositionID: positions[10].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:40 ,ArrangementID: arrangements[0].ID, KomaID: komas[7].ID, PositionID: positions[16].ID, IsFirstMove: false, IsFront: true },

		KomaArrangement{ID:41 ,ArrangementID: arrangements[1].ID, KomaID: komas[0].ID, PositionID: positions[76].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:42 ,ArrangementID: arrangements[1].ID, KomaID: komas[1].ID, PositionID: positions[77].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:43 ,ArrangementID: arrangements[1].ID, KomaID: komas[1].ID, PositionID: positions[75].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:44 ,ArrangementID: arrangements[1].ID, KomaID: komas[2].ID, PositionID: positions[78].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:45 ,ArrangementID: arrangements[1].ID, KomaID: komas[2].ID, PositionID: positions[74].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:46 ,ArrangementID: arrangements[1].ID, KomaID: komas[3].ID, PositionID: positions[79].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:47 ,ArrangementID: arrangements[1].ID, KomaID: komas[3].ID, PositionID: positions[73].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:48 ,ArrangementID: arrangements[1].ID, KomaID: komas[4].ID, PositionID: positions[80].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:49 ,ArrangementID: arrangements[1].ID, KomaID: komas[4].ID, PositionID: positions[72].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:50 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[55].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:51 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[56].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:52 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[57].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:53 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[58].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:54 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[59].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:55 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[60].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:56 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[61].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:57 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[62].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:58 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[63].ID, IsFirstMove: true, IsFront: true },
		KomaArrangement{ID:59 ,ArrangementID: arrangements[1].ID, KomaID: komas[0].ID, PositionID: positions[4].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:60 ,ArrangementID: arrangements[1].ID, KomaID: komas[1].ID, PositionID: positions[5].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:61 ,ArrangementID: arrangements[1].ID, KomaID: komas[1].ID, PositionID: positions[3].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:62 ,ArrangementID: arrangements[1].ID, KomaID: komas[2].ID, PositionID: positions[6].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:63 ,ArrangementID: arrangements[1].ID, KomaID: komas[2].ID, PositionID: positions[2].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:64 ,ArrangementID: arrangements[1].ID, KomaID: komas[3].ID, PositionID: positions[7].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:65 ,ArrangementID: arrangements[1].ID, KomaID: komas[3].ID, PositionID: positions[1].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:66 ,ArrangementID: arrangements[1].ID, KomaID: komas[4].ID, PositionID: positions[8].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:67 ,ArrangementID: arrangements[1].ID, KomaID: komas[4].ID, PositionID: positions[0].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:68 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[10].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:69 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[11].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:70 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[12].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:71 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[13].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:72 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[14].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:73 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[15].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:74 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[16].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:75 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[17].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:76 ,ArrangementID: arrangements[1].ID, KomaID: komas[5].ID, PositionID: positions[18].ID, IsFirstMove: false, IsFront: true },
	}

	return komaArrangements
}


func ExportKomaArrangementForTest() []KomaArrangement {
	arrangements := ExportArrangementForTest()
	komas := ExportKomaForTest()
	positions := ExportPositionForTest()
	komaArrangements := []KomaArrangement{
		KomaArrangement{ID:1 ,ArrangementID: arrangements[0].ID, KomaID: komas[0].ID, PositionID: positions[0].ID, IsFirstMove: false, IsFront: true },
		KomaArrangement{ID:2 ,ArrangementID: arrangements[0].ID, KomaID: komas[2].ID, PositionID: positions[2].ID, IsFirstMove: false, IsFront: true },

	}

	return komaArrangements
}