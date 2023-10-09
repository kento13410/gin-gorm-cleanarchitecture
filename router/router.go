package router

import (
	"github.com/gin-gonic/gin"
	
	"github.com/kento13410/gin-gorm-cleanarchitecture/controller"
)

func NewRouter() {
	r := gin.Default()

	r.POST("/signup", controller.Signup)
}