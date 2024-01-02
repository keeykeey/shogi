package repository

type Position struct {
	ID               uint16   `gorm:"primaryKey"`
	Width            byte
	Height           byte
	Name             string
	BasicColumn
}


func ExportPosition() []Position {
	positions := []Position{
		Position{ID: 1, Width: 9, Height: 1, Name: "九一"},
		Position{ID: 2, Width: 8, Height: 1, Name: "八一"},
		Position{ID: 3, Width: 7, Height: 1, Name: "七一"},
		Position{ID: 4, Width: 6, Height: 1, Name: "六一"},
		Position{ID: 5, Width: 5, Height: 1, Name: "五一"},
		Position{ID: 6, Width: 4, Height: 1, Name: "四一"},
		Position{ID: 7, Width: 3, Height: 1, Name: "三一"},
		Position{ID: 8, Width: 2, Height: 1, Name: "二一"},
		Position{ID: 9, Width: 1, Height: 1, Name: "一一"},
		Position{ID: 10, Width: 9, Height: 2, Name: "九二"},
		Position{ID: 11, Width: 8, Height: 2, Name: "八二"},
		Position{ID: 12, Width: 7, Height: 2, Name: "七二"},
		Position{ID: 13, Width: 6, Height: 2, Name: "六二"},
		Position{ID: 14, Width: 5, Height: 2, Name: "五二"},
		Position{ID: 15, Width: 4, Height: 2, Name: "四二"},
		Position{ID: 16, Width: 3, Height: 2, Name: "三二"},
		Position{ID: 17, Width: 2, Height: 2, Name: "二二"},
		Position{ID: 18, Width: 1, Height: 2, Name: "一二"},
		Position{ID: 19, Width: 9, Height: 3, Name: "九三"},
		Position{ID: 20, Width: 8, Height: 3, Name: "八三"},
		Position{ID: 21, Width: 7, Height: 3, Name: "七三"},
		Position{ID: 22, Width: 6, Height: 3, Name: "六三"},
		Position{ID: 23, Width: 5, Height: 3, Name: "五三"},
		Position{ID: 24, Width: 4, Height: 3, Name: "四三"},
		Position{ID: 25, Width: 3, Height: 3, Name: "三三"},
		Position{ID: 26, Width: 2, Height: 3, Name: "二三"},
		Position{ID: 27, Width: 1, Height: 3, Name: "一三"},
		Position{ID: 28, Width: 9, Height: 4, Name: "九四"},
		Position{ID: 29, Width: 8, Height: 4, Name: "八四"},
		Position{ID: 30, Width: 7, Height: 4, Name: "七四"},
		Position{ID: 31, Width: 6, Height: 4, Name: "六四"},
		Position{ID: 32, Width: 5, Height: 4, Name: "五四"},
		Position{ID: 33, Width: 4, Height: 4, Name: "四四"},
		Position{ID: 34, Width: 3, Height: 4, Name: "三四"},
		Position{ID: 35, Width: 2, Height: 4, Name: "二四"},
		Position{ID: 36, Width: 1, Height: 4, Name: "一四"},
		Position{ID: 37, Width: 9, Height: 5, Name: "九五"},
		Position{ID: 38, Width: 8, Height: 5, Name: "八五"},
		Position{ID: 39, Width: 7, Height: 5, Name: "七五"},
		Position{ID: 40, Width: 6, Height: 5, Name: "六五"},
		Position{ID: 41, Width: 5, Height: 5, Name: "五五"},
		Position{ID: 42, Width: 4, Height: 5, Name: "四五"},
		Position{ID: 43, Width: 3, Height: 5, Name: "三五"},
		Position{ID: 44, Width: 2, Height: 5, Name: "二五"},
		Position{ID: 45, Width: 1, Height: 5, Name: "一五"},
		Position{ID: 46, Width: 9, Height: 6, Name: "九六"},
		Position{ID: 47, Width: 8, Height: 6, Name: "八六"},
		Position{ID: 48, Width: 7, Height: 6, Name: "七六"},
		Position{ID: 49, Width: 6, Height: 6, Name: "六六"},
		Position{ID: 50, Width: 5, Height: 6, Name: "五六"},
		Position{ID: 51, Width: 4, Height: 6, Name: "四六"},
		Position{ID: 52, Width: 3, Height: 6, Name: "三六"},
		Position{ID: 53, Width: 2, Height: 6, Name: "二六"},
		Position{ID: 54, Width: 1, Height: 6, Name: "一六"},
		Position{ID: 55, Width: 9, Height: 7, Name: "九七"},
		Position{ID: 56, Width: 8, Height: 7, Name: "八七"},
		Position{ID: 57, Width: 7, Height: 7, Name: "七七"},
		Position{ID: 58, Width: 6, Height: 7, Name: "六七"},
		Position{ID: 59, Width: 5, Height: 7, Name: "五七"},
		Position{ID: 60, Width: 4, Height: 7, Name: "四七"},
		Position{ID: 61, Width: 3, Height: 7, Name: "三七"},
		Position{ID: 62, Width: 2, Height: 7, Name: "二七"},
		Position{ID: 63, Width: 1, Height: 7, Name: "一七"},
		Position{ID: 64, Width: 9, Height: 8 , Name: "九八"},
		Position{ID: 65, Width: 8, Height: 8 , Name: "八八"},
		Position{ID: 66, Width: 7, Height: 8 , Name: "七八"},
		Position{ID: 67, Width: 6, Height: 8 , Name: "六八"},
		Position{ID: 68, Width: 5, Height: 8 , Name: "五八"},
		Position{ID: 69, Width: 4, Height: 8 , Name: "四八"},
		Position{ID: 70, Width: 3, Height: 8 , Name: "三八"},
		Position{ID: 71, Width: 2, Height: 8 , Name: "二八"},
		Position{ID: 72, Width: 1, Height: 8 , Name: "一八"},
		Position{ID: 73, Width: 9, Height: 9 , Name: "九九"},
		Position{ID: 74, Width: 8, Height: 9 , Name: "八九"},
		Position{ID: 75, Width: 7, Height: 9 , Name: "七九"},
		Position{ID: 76, Width: 6, Height: 9 , Name: "六九"},
		Position{ID: 77, Width: 5, Height: 9 , Name: "五九"},
		Position{ID: 78, Width: 4, Height: 9 , Name: "四九"},
		Position{ID: 79, Width: 3, Height: 9 , Name: "三九"},
		Position{ID: 80, Width: 2, Height: 9 , Name: "二九"},
		Position{ID: 81, Width: 1, Height: 9 , Name: "一九"},
	}
	return positions
}