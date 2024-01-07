package repository

type Board struct {
	ID                  byte         `gorm:"primaryKey"`
	BoardTop            uint16       `json:"boardTop"`
	BoardLeft           uint16       `json:"boardLeft"`
	SquareHeightCount   byte         `json:"squareHeightCount"`
	SquareWidthCount    byte         `json:"squareWidthCount"`
	SquareHeightLen     uint16       `json:"squareHeightLen"`
	SquareWidthLen      uint16       `json:"squareWidthLen"`
	LineWidth           byte         `json:"lineWidth"`
	BasicColumn
}

type BoardBase struct {
	ID                  byte         `gorm:"primaryKey"`
	BoardTop            uint16       `json:"boardTop"`
	BoardLeft           uint16       `json:"boardLeft"`
	SquareHeightCount   byte         `json:"squareHeightCount"`
	SquareWidthCount    byte         `json:"squareWidthCount"`
	SquareHeightLen     uint16       `json:"squareHeightLen"`
	SquareWidthLen      uint16       `json:"squareWidthLen"`
	LineWidth           byte         `json:"lineWidth"`
}
func GetBoard(boardId byte) BoardBase {
	var board BoardBase
	db := GetDbConnection()
	db.Table("boards").Where("id = ?", boardId).Find(&board)

	return board
}