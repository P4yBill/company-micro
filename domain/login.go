package domain

import (
	"company-micro/util"
	"context"
	"errors"
	"net/http"
	"net/mail"
)

type LoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

// used by chi.render to bind struct to the request
func (lr *LoginRequest) Bind(r *http.Request) error {
	_, err := mail.ParseAddress(lr.Email)

	if err != nil {
		return errors.New("Bad email.")
	}

	if util.IsStringBlank(lr.Password) && len(lr.Password) > 5 {
		return errors.New("Bad input password. Password should consist of 5 or more alphanumeric characters")
	}

	return nil
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Username     string `json:"username"`
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
