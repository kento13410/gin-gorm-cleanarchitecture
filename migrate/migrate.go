package main

import (
	"go-gin-gorm-example/db"
	"go-gin-gorm-example/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DB := db.DBOpen()
	defer db.DBClose(DB)
	DB.AutoMigrate(models.User{})
}