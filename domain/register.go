package domain

import (
	"company-micro/util"
	"context"
	"errors"
	"net/http"
	"net/mail"
)

type RegisterRequest struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (rr *RegisterRequest) Bind(r *http.Request) error {
	_, err := mail.ParseAddress(rr.Email)

	if err != nil {
		return errors.New("Bad email.")
	}

	if util.IsStringBlank(rr.Password) || len(rr.Password) <= 5 {
		return errors.New("Bad input password. Password should consist of 5 or more alphanumeric characters")
	}

	if util.IsStringBlank(rr.Username) || len(rr.Username) <= 4 {
		return errors.New("Bad input username. Username should consist of 5 or more alphanumeric characters")
	}

	return nil
}

type RegisterResponse struct {
	Response
}

type RegisterUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
