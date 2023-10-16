package usecase

import (
	"go-gin-gorm-example/models"
	"go-gin-gorm-example/repository"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(c *gin.Context, user models.User)
	LogIn(c *gin.Context, user models.User) string
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

func (uu *UserUsecase) LogIn(c *gin.Context, user models.User) string {
	newUser, err := uu.ur.FindUser(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":"login completed!",
		})
	}

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	claims := jwt.MapClaims{
		"user_id": newUser.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(os.Getenv("JWT_KEY"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	}

	return tokenString
}