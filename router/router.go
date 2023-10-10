package router

import (
	"go-gin-gorm-example/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(uc controller.IUserController) {
	r := gin.Default()
	
	r.POST("/signup", uc.SignUp)

	r.Run(":8081")
}