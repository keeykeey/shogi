package repository

type KomaArrangements struct {
	Id         			uint16
	ArrangementId       uint16
	Arrangement         Arrangement
    KomaId              byte
	Koma                Koma
	PositionId          uint16
	Position            Position
	IsFirstMove         bool
	IsFront             bool
}


func GetKomaArrangements(arrangementId uint16) []KomaArrangements {
    var komaArrangements []KomaArrangements
	db := GetDbConnection()
	db.Table("koma_arrangements").Where("arrangement_id = ?", arrangementId).Scan(&komaArrangements)
	return komaArrangements;
}