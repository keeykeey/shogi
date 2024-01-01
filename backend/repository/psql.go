package repository

import (
  _ "github.com/lib/pq"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "github.com/joho/godotenv"
  "fmt"
  "os"
) 

var db *gorm.DB

func GetDbConnection() *gorm.DB {
	err := godotenv.Load(".env");
	if err != nil {
		panic("couldn't find env file\n")
	}

	HOST_NAME := os.Getenv("HOST_NAME")
	USER_NAME := os.Getenv("USER_NAME")
	PASSWORD := os.Getenv("PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Tokyo", 
		HOST_NAME,
		USER_NAME, 
		PASSWORD, 
		DB_NAME, 
		DB_PORT)

	var err2 error
	db, err2 = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err2 != nil {
		fmt.Printf("%v",dsn)
		panic("couldn't connect to database\n")
	}

	return db
}
