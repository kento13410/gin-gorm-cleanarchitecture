package db

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func DBOpen() *gorm.DB {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	USER := os.Getenv("USER")
	PASS := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("DBNAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		panic(err)
	}

	return db
}

func DBClose(db *gorm.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}