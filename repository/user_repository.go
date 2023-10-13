package repository

import (
	"go-gin-gorm-example/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	AddUser(c *gin.Context, user models.User) error
	FindUser(c *gin.Context, user models.User) (string, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) AddUser(c *gin.Context, user models.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) FindUser(c *gin.Context, user models.User) (string, error) {
	newUser := models.User{}
	if err := ur.db.Select("password").Where("mail = ?", user.Mail).First(&newUser).Error; err != nil {
		return "", err
	}
	return newUser.Password, nil
}
