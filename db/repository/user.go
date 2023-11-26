package repository

// Model User
type User struct {
	Id                 uint16    `gorm:"primaryKey"`
	Name               string
	BasicColumn
}

func ExportUser() []User {
	users := []User{
		User{Id: 1, Name: "羽生"},
		User{Id: 2, Name: "藤井"},
		User{Id: 3, Name: "森内"},
		User{Id: 4, Name: "大山"},
		User{Id: 5, Name: "渡辺"},
	}
	return users
}
