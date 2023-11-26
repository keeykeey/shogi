package repository

// intermediate table between model Arrangement and model Koma
type KomaArrangement struct {
	Id                   uint16           `gorm:"primaryKey"`
	ArrangementId        uint16
	Arrangement          Arrangement      `gorm:"foreignKey:ArrangementId"`
	KomaId               uint8
	Koma                 Koma             `gorm:"foreignKey:KomaId"`
	PositionId           uint16
	Position             Position         `gorm:"foreignKey:PositionId"`
	IsFirstMove          bool  // 先手か後手か
	IsFront              bool  // 駒が成っているか否か
	BasicColumn
}

func ExportKomaArrangement() []KomaArrangement {
	arrangements := ExportArrangement()
	komas := ExportKoma()
	positions := ExportPosition()
	komaArrangements := []KomaArrangement{
		KomaArrangement{Id:1 ,ArrangementId: arrangements[0].Id, KomaId: komas[0].Id, PositionId: positions[76].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:2 ,ArrangementId: arrangements[0].Id, KomaId: komas[1].Id, PositionId: positions[77].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:3 ,ArrangementId: arrangements[0].Id, KomaId: komas[1].Id, PositionId: positions[75].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:4 ,ArrangementId: arrangements[0].Id, KomaId: komas[2].Id, PositionId: positions[78].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:5 ,ArrangementId: arrangements[0].Id, KomaId: komas[2].Id, PositionId: positions[74].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:6 ,ArrangementId: arrangements[0].Id, KomaId: komas[3].Id, PositionId: positions[79].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:7 ,ArrangementId: arrangements[0].Id, KomaId: komas[3].Id, PositionId: positions[73].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:8 ,ArrangementId: arrangements[0].Id, KomaId: komas[4].Id, PositionId: positions[80].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:9 ,ArrangementId: arrangements[0].Id, KomaId: komas[4].Id, PositionId: positions[72].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:10 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[55].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:11 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[56].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:12 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[57].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:13 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[58].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:14 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[59].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:15 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[60].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:16 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[61].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:17 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[62].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:18 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[63].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:19 ,ArrangementId: arrangements[0].Id, KomaId: komas[6].Id, PositionId: positions[70].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:20 ,ArrangementId: arrangements[0].Id, KomaId: komas[7].Id, PositionId: positions[64].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:21 ,ArrangementId: arrangements[0].Id, KomaId: komas[0].Id, PositionId: positions[4].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:22 ,ArrangementId: arrangements[0].Id, KomaId: komas[1].Id, PositionId: positions[5].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:23 ,ArrangementId: arrangements[0].Id, KomaId: komas[1].Id, PositionId: positions[3].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:24 ,ArrangementId: arrangements[0].Id, KomaId: komas[2].Id, PositionId: positions[6].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:25 ,ArrangementId: arrangements[0].Id, KomaId: komas[2].Id, PositionId: positions[2].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:26 ,ArrangementId: arrangements[0].Id, KomaId: komas[3].Id, PositionId: positions[7].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:27 ,ArrangementId: arrangements[0].Id, KomaId: komas[3].Id, PositionId: positions[1].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:28 ,ArrangementId: arrangements[0].Id, KomaId: komas[4].Id, PositionId: positions[8].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:29 ,ArrangementId: arrangements[0].Id, KomaId: komas[4].Id, PositionId: positions[0].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:30 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[10].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:31 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[11].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:32 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[12].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:33 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[13].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:34 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[14].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:35 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[15].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:36 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[16].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:37 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[17].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:38 ,ArrangementId: arrangements[0].Id, KomaId: komas[5].Id, PositionId: positions[18].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:39 ,ArrangementId: arrangements[0].Id, KomaId: komas[6].Id, PositionId: positions[19].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:40 ,ArrangementId: arrangements[0].Id, KomaId: komas[7].Id, PositionId: positions[20].Id, IsFirstMove: false, IsFront: true },

		KomaArrangement{Id:41 ,ArrangementId: arrangements[1].Id, KomaId: komas[0].Id, PositionId: positions[76].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:42 ,ArrangementId: arrangements[1].Id, KomaId: komas[1].Id, PositionId: positions[77].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:43 ,ArrangementId: arrangements[1].Id, KomaId: komas[1].Id, PositionId: positions[75].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:44 ,ArrangementId: arrangements[1].Id, KomaId: komas[2].Id, PositionId: positions[78].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:45 ,ArrangementId: arrangements[1].Id, KomaId: komas[2].Id, PositionId: positions[74].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:46 ,ArrangementId: arrangements[1].Id, KomaId: komas[3].Id, PositionId: positions[79].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:47 ,ArrangementId: arrangements[1].Id, KomaId: komas[3].Id, PositionId: positions[73].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:48 ,ArrangementId: arrangements[1].Id, KomaId: komas[4].Id, PositionId: positions[80].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:49 ,ArrangementId: arrangements[1].Id, KomaId: komas[4].Id, PositionId: positions[72].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:50 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[55].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:51 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[56].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:52 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[57].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:53 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[58].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:54 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[59].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:55 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[60].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:56 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[61].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:57 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[62].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:58 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[63].Id, IsFirstMove: true, IsFront: true },
		KomaArrangement{Id:59 ,ArrangementId: arrangements[1].Id, KomaId: komas[0].Id, PositionId: positions[4].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:60 ,ArrangementId: arrangements[1].Id, KomaId: komas[1].Id, PositionId: positions[5].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:61 ,ArrangementId: arrangements[1].Id, KomaId: komas[1].Id, PositionId: positions[3].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:62 ,ArrangementId: arrangements[1].Id, KomaId: komas[2].Id, PositionId: positions[6].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:63 ,ArrangementId: arrangements[1].Id, KomaId: komas[2].Id, PositionId: positions[2].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:64 ,ArrangementId: arrangements[1].Id, KomaId: komas[3].Id, PositionId: positions[7].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:65 ,ArrangementId: arrangements[1].Id, KomaId: komas[3].Id, PositionId: positions[1].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:66 ,ArrangementId: arrangements[1].Id, KomaId: komas[4].Id, PositionId: positions[8].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:67 ,ArrangementId: arrangements[1].Id, KomaId: komas[4].Id, PositionId: positions[0].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:68 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[10].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:69 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[11].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:70 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[12].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:71 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[13].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:72 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[14].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:73 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[15].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:74 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[16].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:75 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[17].Id, IsFirstMove: false, IsFront: true },
		KomaArrangement{Id:76 ,ArrangementId: arrangements[1].Id, KomaId: komas[5].Id, PositionId: positions[18].Id, IsFirstMove: false, IsFront: true },
	}

	return komaArrangements
}