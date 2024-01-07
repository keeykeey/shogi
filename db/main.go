package main

import (
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"db/repository"
	"github.com/joho/godotenv"
	"fmt"
	"os"
)

func ConnectPsql() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic("couldn't find env file\n")
	}

	HOST_NAME := os.Getenv("HOST_NAME")
	USER_NAME := os.Getenv("USER_NAME")
	PASSWORD := os.Getenv("PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	PORT := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Tokyo", HOST_NAME, USER_NAME, PASSWORD, DB_NAME, PORT)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("couldn't connect to database\n")
	}
	return db
}

func delete(db *gorm.DB) {
	db.Exec("DROP TABLE IF EXISTS users CASCADE")
	db.Exec("DROP TABLE IF EXISTS games CASCADE")
	db.Exec("DROP TABLE IF EXISTS komas CASCADE")
	db.Exec("DROP TABLE IF EXISTS boards CASCADE")
	db.Exec("DROP TABLE IF EXISTS koma_arrangements CASCADE")
	db.Exec("DROP TABLE IF EXISTS arrangements CASCADE")
	db.Exec("DROP TABLE IF EXISTS positions CASCADE")
}

func insert(db *gorm.DB) {
	// Migration
	db.AutoMigrate(
		&repository.User{},
		&repository.Board{},
		&repository.Koma{},
		&repository.Position{},
		&repository.Arrangement{},
		&repository.KomaArrangement{},
		&repository.Game{}, 
	)

	// User
	user := repository.ExportUser()
	db.Create(&user)

	// Board
	board := repository.ExportBoard()
	db.Create(&board)

	// Koma
	koma := repository.ExportKoma()
	db.Create(&koma)

	// Position
	position := repository.ExportPosition()
	db.Create(&position)

	// Arrangement
	arrangement := repository.ExportArrangement()
	db.Create(&arrangement)

	// KomaArrangement
	komaArrangement := repository.ExportKomaArrangement()
	db.Create(&komaArrangement)

	// Game
	game := repository.ExportGame()
	db.Create(&game)

} 

func insertForTest(db *gorm.DB) {
	// Migration
	db.AutoMigrate(
		&repository.User{},
		&repository.Koma{},
		&repository.Position{},
		&repository.Arrangement{},
		&repository.KomaArrangement{},
		&repository.Game{}, 
	)

	// User
	user := repository.ExportUserForTest()
	db.Create(&user)

	// Board
	board := repository.ExportBoardForTest()
	db.Create(&board)

	// Koma
	koma := repository.ExportKomaForTest()
	db.Create(&koma)

	// Position
	position := repository.ExportPositionForTest()
	db.Create(&position)

	// Arrangement
	arrangement := repository.ExportArrangementForTest()
	db.Create(&arrangement)

	// KomaArrangement
	komaArrangement := repository.ExportKomaArrangementForTest()
	db.Create(&komaArrangement)

	// Game
	game := repository.ExportGameForTest()
	db.Create(&game)

} 

func main() {
	db := ConnectPsql()
	
	if len(os.Args) > 1 {
		var argv = fmt.Sprintf("%s",os.Args[1])
		var isTest = (argv == "test")
		if isTest {
			delete(db)
			insertForTest(db)
		}
		return
	}

	delete(db)
	insert(db)
}