package repository

// Model Game
type Game struct {
	ID                     uint16       `gorm:"primaryKey"`
	ArrangementID          uint16
	Arrangement            Arrangement  `gorm:"foreignKey:ArrangementID"`
	FirstPlayerID          uint16
	FirstPlayer            User         `gorm:"foreignKey:FirstPlayerID"`
	SecondPlayerID         uint16
	SecondPlayer           User         `gorm:"foreignKey:SecondPlayerID"`
	WinnerID               uint16
	Winner                 User         `gorm:"foreignKey:WinnerID"`
	LoserID                uint16
	Loser                  User         `gorm:"foreignKey:LoserID"`
	BasicColumn
}

func ExportGame() []Game {
	users := ExportUser()
	arrangements := ExportArrangement()
	games := []Game{
		Game{ID: 1, ArrangementID: arrangements[0].ID, FirstPlayerID: users[0].ID, SecondPlayerID: users[1].ID, WinnerID: users[0].ID, LoserID: users[1].ID},
		Game{ID: 2, ArrangementID: arrangements[1].ID, FirstPlayerID: users[1].ID, SecondPlayerID: users[2].ID, WinnerID: users[1].ID, LoserID: users[2].ID},
		Game{ID: 3, ArrangementID: arrangements[0].ID, FirstPlayerID: users[2].ID, SecondPlayerID: users[3].ID, WinnerID: users[3].ID, LoserID: users[2].ID},
		Game{ID: 4, ArrangementID: arrangements[0].ID, FirstPlayerID: users[3].ID, SecondPlayerID: users[0].ID, WinnerID: users[0].ID, LoserID: users[3].ID},
	}
	return games
}