package repository

type Board struct {
	ID                  byte         `gorm:"primaryKey"`
	BoardTop            uint16
	BoardLeft           uint16
	SquareHeightCount   byte
	SquarewidthCount    byte
	SquareHeightLen     uint16
	SquareWidthLen      uint16
	lineWidth           byte
}

func ExportBoard() []Board {
	boards := []Board{
		Board{
			ID: 1,
			BoardTop: 65,
			BoardLeft: 41,
			SquareHeightCount: 9,
			SquarewidthCount: 9,
			SquareHeightLen: 30,
			SquareWidthLen: 30,
			lineWidth: 2,
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
			SquarewidthCount: 9,
			SquareHeightLen: 30,
			SquareWidthLen: 30,
			lineWidth: 2,
		},
	}
	return boards
}
