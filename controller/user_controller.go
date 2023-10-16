package controller

import (
	"go-gin-gorm-example/models"
	"go-gin-gorm-example/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	SignUp(c *gin.Context)
	LogIn(c *gin.Context)
	LogOut(c *gin.Context)
}

type UserController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &UserController{
		uu: uu,
	}
}

func (uc *UserController) SignUp(c *gin.Context) {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	}

	uc.uu.SignUp(c, user)
}

func (uc *UserController) LogIn(c *gin.Context) {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	}
	c.SetCookie("user", uc.uu.LogIn(c, user), 24, "/", "localhost", false, true)
}

func (uc *UserController) LogOut(c *gin.Context) {
	c.SetCookie("user", "", -1, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{
		"message":"logout completed!",
	})
}