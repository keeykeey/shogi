package repository

// Model User
type User struct {
	ID                 uint16    `gorm:"primaryKey"`
	Name               string
	BasicColumn
}

func ExportUser() []User {
	users := []User{
		User{ID: 1, Name: "羽生"},
		User{ID: 2, Name: "藤井"},
		User{ID: 3, Name: "森内"},
		User{ID: 4, Name: "大山"},
		User{ID: 5, Name: "渡辺"},
	}
	return users
}
