package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name 		string	`json:"name"`
	Mail		string	`json:"mail"`
	Password	string	`json:"password"`
}

type UserResponce struct {
	Name 		string	`json:"name"`
	Mail		string	`json:"mail"`
}