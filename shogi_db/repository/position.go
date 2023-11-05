package repository

import (
	"gorm.io/gorm"
)

type Position struct {
	Id               uint16   `gorm:"primaryKey"`
	Number           uint8
	Name             string
	gorm.Model
}


func ExportPosition() []Position {
	positions := []Position{
		Position{Id: 1, Number: 91, Name: "九一"},
		Position{Id: 2, Number: 81, Name: "八一"},
		Position{Id: 3, Number: 71, Name: "七一"},
		Position{Id: 4, Number: 61, Name: "六一"},
		Position{Id: 5, Number: 51, Name: "五一"},
		Position{Id: 6, Number: 41, Name: "四一"},
		Position{Id: 7, Number: 31, Name: "三一"},
		Position{Id: 8, Number: 21, Name: "二一"},
		Position{Id: 9, Number: 11, Name: "一一"},
		Position{Id: 10, Number: 92, Name: "九二"},
		Position{Id: 11, Number: 82, Name: "八二"},
		Position{Id: 12, Number: 72, Name: "七二"},
		Position{Id: 13, Number: 62, Name: "六二"},
		Position{Id: 14, Number: 52, Name: "五二"},
		Position{Id: 15, Number: 42, Name: "四二"},
		Position{Id: 16, Number: 32, Name: "三二"},
		Position{Id: 17, Number: 22, Name: "二二"},
		Position{Id: 18, Number: 12, Name: "一二"},
		Position{Id: 19, Number: 93, Name: "九三"},
		Position{Id: 20, Number: 83, Name: "八三"},
		Position{Id: 21, Number: 73, Name: "七三"},
		Position{Id: 22, Number: 63, Name: "六三"},
		Position{Id: 23, Number: 53, Name: "五三"},
		Position{Id: 24, Number: 43, Name: "四三"},
		Position{Id: 25, Number: 33, Name: "三三"},
		Position{Id: 26, Number: 23, Name: "二三"},
		Position{Id: 27, Number: 13, Name: "一三"},
		Position{Id: 28, Number: 94, Name: "九四"},
		Position{Id: 29, Number: 84, Name: "八四"},
		Position{Id: 30, Number: 74, Name: "七四"},
		Position{Id: 31, Number: 64, Name: "六四"},
		Position{Id: 32, Number: 54, Name: "五四"},
		Position{Id: 33, Number: 44, Name: "四四"},
		Position{Id: 34, Number: 34, Name: "三四"},
		Position{Id: 35, Number: 24, Name: "二四"},
		Position{Id: 36, Number: 14, Name: "一四"},
		Position{Id: 37, Number: 95, Name: "九五"},
		Position{Id: 38, Number: 85, Name: "八五"},
		Position{Id: 39, Number: 75, Name: "七五"},
		Position{Id: 40, Number: 65, Name: "六五"},
		Position{Id: 41, Number: 55, Name: "五五"},
		Position{Id: 42, Number: 45, Name: "四五"},
		Position{Id: 43, Number: 35, Name: "三五"},
		Position{Id: 44, Number: 25, Name: "二五"},
		Position{Id: 45, Number: 15, Name: "一五"},
		Position{Id: 46, Number: 96, Name: "九六"},
		Position{Id: 47, Number: 86, Name: "八六"},
		Position{Id: 48, Number: 76, Name: "七六"},
		Position{Id: 49, Number: 66, Name: "六六"},
		Position{Id: 50, Number: 56, Name: "五六"},
		Position{Id: 51, Number: 46, Name: "四六"},
		Position{Id: 52, Number: 36, Name: "三六"},
		Position{Id: 53, Number: 26, Name: "二六"},
		Position{Id: 54, Number: 16, Name: "一六"},
		Position{Id: 55, Number: 97, Name: "九七"},
		Position{Id: 56, Number: 87, Name: "八七"},
		Position{Id: 57, Number: 77, Name: "七七"},
		Position{Id: 58, Number: 67, Name: "六七"},
		Position{Id: 59, Number: 57, Name: "五七"},
		Position{Id: 60, Number: 47, Name: "四七"},
		Position{Id: 61, Number: 37, Name: "三七"},
		Position{Id: 62, Number: 27, Name: "二七"},
		Position{Id: 63, Number: 17, Name: "一七"},
		Position{Id: 64, Number: 98 , Name: "九八"},
		Position{Id: 65, Number: 88 , Name: "八八"},
		Position{Id: 66, Number: 78 , Name: "七八"},
		Position{Id: 67, Number: 68 , Name: "六八"},
		Position{Id: 68, Number: 58 , Name: "五八"},
		Position{Id: 69, Number: 48 , Name: "四八"},
		Position{Id: 70, Number: 38 , Name: "三八"},
		Position{Id: 71, Number: 28 , Name: "二八"},
		Position{Id: 72, Number: 18 , Name: "一八"},
		Position{Id: 73, Number: 99 , Name: "九九"},
		Position{Id: 74, Number: 89 , Name: "八九"},
		Position{Id: 75, Number: 79 , Name: "七九"},
		Position{Id: 76, Number: 69 , Name: "六九"},
		Position{Id: 77, Number: 59 , Name: "五九"},
		Position{Id: 78, Number: 49 , Name: "四九"},
		Position{Id: 79, Number: 39 , Name: "三九"},
		Position{Id: 80, Number: 29 , Name: "二九"},
		Position{Id: 81, Number: 19 , Name: "一九"},
	}
	return positions
}