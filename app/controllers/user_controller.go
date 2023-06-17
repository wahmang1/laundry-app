package userController

import (
	"fmt"
	"laundryApp/app/models"
	"laundryApp/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Index(c *gin.Context) {
	var users []models.User

	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})

}

func Create(c *gin.Context) {
	var input models.User
	fmt.Println(true)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := []byte(input.Password)

	hashedPassword, error := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if error != nil {
		panic(error)
	}

	// Create Product
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	fmt.Println(true)

	config.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})

}
