package repository

import (
	"gorm.io/gorm"
)

type Koma struct {
	komaBase KomaBase
	gorm.Model
}

type KomaBase struct {
	Id                 uint8
	MoveId             uint8
	MoveId2            uint8
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
