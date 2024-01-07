package repository

type Board struct {
	ID                  byte         `gorm:"primaryKey"`
	BoardTop            uint16
	BoardLeft           uint16
	SquareHeightCount   byte
	SquareWidthCount    byte
	SquareHeightLen     uint16
	SquareWidthLen      uint16
	LineWidth           byte
	BasicColumn
}

func ExportBoard() []Board {
	boards := []Board{
		Board{
			ID: 1,
			BoardTop: 65,
			BoardLeft: 41,
			SquareHeightCount: 9,
			SquareWidthCount: 9,
			SquareHeightLen: 30,
			SquareWidthLen: 30,
			LineWidth: 2,
		},
	}
	return boards
}

func ExportBoardForTest() []Board {
	boards := []Board{
		Board{
			ID: 1,
			BoardTop: 65,
			BoardLeft: 41,
			SquareHeightCount: 9,
			SquareWidthCount: 9,
			SquareHeightLen: 30,
			SquareWidthLen: 30,
			LineWidth: 2,
		},
	}
	return boards
}
