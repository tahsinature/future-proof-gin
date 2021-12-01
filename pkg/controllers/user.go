package controllers

import (
	"net/http"

	"github.com/tahsinature/future-proof-gin/pkg/db/repositories"
	"github.com/tahsinature/future-proof-gin/pkg/forms"
	"github.com/tahsinature/future-proof-gin/pkg/services"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var (
	userForm       = new(forms.UserForm)
	userRepository = new(repositories.UserRepository)
	authService    = new(services.AuthService)
)

func (ctrl UserController) Login(c *gin.Context) {
	var loginForm forms.LoginForm

	if validationErr := c.ShouldBindJSON(&loginForm); validationErr != nil {
		message := userForm.Login(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, err := userRepository.GetUserByEmail(loginForm.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

	bytePassword := []byte(loginForm.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid login details"})
		return
	}

	token, err := authService.CreateToken(int64(user.ID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "user": user, "token": token})
}

func (ctrl UserController) Register(c *gin.Context) {
	var registerForm forms.RegisterForm

	if validationErr := c.ShouldBindJSON(&registerForm); validationErr != nil {
		message := userForm.Register(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, err := userRepository.Register(registerForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered", "user": user})
}
