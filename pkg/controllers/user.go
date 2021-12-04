package controllers

import (
	"fmt"

	"github.com/tahsinature/future-proof-gin/pkg/forms"
	"github.com/tahsinature/future-proof-gin/pkg/services"

	"github.com/gin-gonic/gin"
)

type User struct{}

var authService = new(services.AuthService)

func (ctrl User) Login(c *gin.Context) {
	var body forms.Login
	if isValid := ValidateBody(&body, c); !isValid {
		return
	}

	err, data := authService.HandleLogin(body)

	if err != nil {
		Response.FromError(c, *err)
		return
	}

	Response.Success(c, data)
}

func (ctrl User) Register(c *gin.Context) {
	fmt.Println("registering")
}
