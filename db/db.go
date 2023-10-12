package db

import "github.com/jinzhu/gorm"

func DBOpen() *gorm.DB {
	USER := "root"
	PASS := "#kK36238708#"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "goapp"
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