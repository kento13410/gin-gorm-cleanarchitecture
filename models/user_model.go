package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name 		string	`form:"name"`
	Mail		string	`form:"mail" gorm:"unique"`
	Password	string	`form:"password"`
}
