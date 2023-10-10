package main

import (
	"go-gin-gorm-example/controller"
	"go-gin-gorm-example/db"
	"go-gin-gorm-example/repository"
	"go-gin-gorm-example/router"
	"go-gin-gorm-example/usecase"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DB := db.DBOpen()
	defer db.DBClose(DB)
	ur := repository.NewUserRepository(DB)
	uuc := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uuc)
	router.NewRouter(uc)
}
