package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Actor struct {
	FirstName	string `gorm:"column:first_name" json:"first_name"`
	LastName	string `gorm:"column:last_name" json:"last_name"`
}

func (Actor) TableName() string {
	return "actor"
}

func main() {
	// GET, POST, QueryParameter
    r := gin.Default()
    
	db := gormConnect()
	defer db.Close()

    r.GET("actor/:name", func(ctx *gin.Context){
		query := ctx.Param("name")
		actor := searchActor(ctx, db, query)
		ctx.JSON(http.StatusOK, actor)
	})

	r.GET("actor", func(ctx *gin.Context){
		actor := []Actor{}
		ctx.JSON(http.StatusOK, db.Limit(3).Find(&actor))
	})
	
	r.Run(":8081")
}

func searchActor(ctx *gin.Context, db *gorm.DB, actorName string) Actor {
	fullnameActor := Actor{}
	db.Where("first_name = ?", actorName).First(&fullnameActor)
	return fullnameActor
}

func gormConnect() *gorm.DB {
	USER := "root"
	PASS := "#kK36238708#"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "sakila"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		panic(err)
	}

	return db
}
