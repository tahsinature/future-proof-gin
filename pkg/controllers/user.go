package controllers

import (
	"fmt"

	"github.com/tahsinature/future-proof-gin/pkg/db/repositories"
	"github.com/tahsinature/future-proof-gin/pkg/forms"
	"github.com/tahsinature/future-proof-gin/pkg/services"

	"github.com/gin-gonic/gin"
)

type User struct{}

var (
	userForm       = new(forms.UserForm)
	userRepository = new(repositories.UserRepository)
	authService    = new(services.AuthService)
)

func (ctrl User) Login(c *gin.Context) {
	fmt.Println("logging in")
}

func (ctrl User) Register(c *gin.Context) {
	fmt.Println("registering")
}
