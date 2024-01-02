package repository

type Position struct {
	ID               uint16      `gorm:"primaryKey" json:"id"`
	Width            byte        `json:"width"`
	Height           byte        `json:"height"`
	Name             string      `json:"name"`
	BasicColumn
}

type PositionBase struct {
	ID               uint16      `gorm"primaryKey" json:"id"`
	Width            byte        `json:"width"`
	Height           byte        `json:"height"`
	Name             string      `json:"name"`
}
