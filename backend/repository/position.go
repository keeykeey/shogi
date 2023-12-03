package repository

type Position struct {
	ID               uint16
	Number           uint8
	Name             string
	BasicColumn
}

type PositionBase struct {
	ID               uint16
	Number           uint8
	Name             string
}
