package services

import (
	"net/http"

	"github.com/tahsinature/future-proof-gin/pkg/error"
	"github.com/tahsinature/future-proof-gin/pkg/forms"
)

type AuthService struct{}

func (AuthService) HandleLogin(payload forms.Login) (err *error.Response, data interface{}) {
	err = new(error.Response).New(http.StatusUnauthorized, error.Flags.Get("INVALID_LOGIN"), "Invalid credentials")

	return err, map[string]string{
		"refreshToken": "refresh_token", "accessToken": "access_token",
	}
}
