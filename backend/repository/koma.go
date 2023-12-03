package repository

type Koma struct {
	ID                 byte           `gorm:"primaryKey"`
	MoveID             byte
	MoveID2            byte
	Name               string
	Name2              string
	BasicColumn
}

type KomaBase struct {
	ID                 byte           `gorm:"primaryKey"`
	MoveID             byte
	MoveID2            byte
	Name               string
	Name2              string
}

func GetKomas() []KomaBase {
    var komas []KomaBase
	db := GetDbConnection()
	db.Find(&komas)
	db.Table("komas").Select("id", "move_id", "move_id2", "name", "name2").Scan(&komas);
	return komas;
}
