package repository

type Arrangement struct {
	ID                  uint16     `gorm:"primaryKey" json:"id"`
	Name                string     `json:"name"`
	BasicColumn
}
