package repository

type Koma struct {
	ID                 byte           `gorm:"primaryKey" json:"id"`
	MoveID             byte           `json:"moveId"`
	MoveID2            byte           `json:"moveId2"`
	Name               string         `json:"name"`
	Name2              string         `json:"name2"`
	BasicColumn
}

type KomaBase struct {
	ID                 byte           `gorm:"primaryKey" json:"id"`
	MoveID             byte           `json:"moveId"`
	MoveID2            byte           `json:"moveId2"`
	Name               string         `json:"name"`
	Name2              string         `json:"name2"`
}

func GetKomas() []KomaBase {
    var komas []KomaBase
	db := GetDbConnection()
	db.Find(&komas)
	db.Table("komas").Select("id", "move_id", "move_id2", "name", "name2").Scan(&komas);
	return komas;
}
