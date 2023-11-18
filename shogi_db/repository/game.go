package repository

import (
	"gorm.io/gorm"
)

// Model Game
type Game struct {
	Id                     uint16       `gorm:"primaryKey"`
	ArrangementId          uint16
	Arrangement            Arrangement  `gorm:"foreignKey:ArrangementId"`
	FirstPlayerId          uint16
	FirstPlayer            User         `gorm:"foreignKey:FirstPlayerId"`
	SecondPlayerId         uint16
	SecondPlayer           User         `gorm:"foreignKey:SecondPlayerId"`
	WinnerId               uint16
	Winner                 User         `gorm:"foreignKey:WinnerId"`
	LoserId                uint16
	Loser                  User         `gorm:"foreignKey:LoserId"`
	gorm.Model
}

func ExportGame() []Game {
	users := ExportUser()
	arrangements := ExportArrangement()
	games := []Game{
		Game{Id: 1, ArrangementId: arrangements[0].Id, FirstPlayerId: users[0].Id, SecondPlayerId: users[1].Id, WinnerId: users[0].Id, LoserId: users[1].Id},
		Game{Id: 2, ArrangementId: arrangements[1].Id, FirstPlayerId: users[1].Id, SecondPlayerId: users[2].Id, WinnerId: users[1].Id, LoserId: users[2].Id},
		Game{Id: 3, ArrangementId: arrangements[0].Id, FirstPlayerId: users[2].Id, SecondPlayerId: users[3].Id, WinnerId: users[3].Id, LoserId: users[2].Id},
		Game{Id: 4, ArrangementId: arrangements[0].Id, FirstPlayerId: users[3].Id, SecondPlayerId: users[0].Id, WinnerId: users[0].Id, LoserId: users[3].Id},
	}
	return games
}