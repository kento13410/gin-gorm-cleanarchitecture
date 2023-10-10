package usecase

import (
	"go-gin-gorm-example/models"
	"go-gin-gorm-example/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(c *gin.Context, user models.User)
}

type UserUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		ur: ur,
	}
}

func (uu *UserUsecase) SignUp(c *gin.Context, user models.User) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	}
	if err := uu.ur.AddUser(c, models.User{
		Name: user.Name,
		Mail: user.Mail,
		Password: string(hashedPass),
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	}
}