package domain

import (
	"company-micro/util"
	"context"
	"net/http"
)

type RefreshTokenRequest struct {
	GrantType    string `form:"grant_type" validate:"required"`
	RefreshToken string `form:"refresh_token" validate:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenUsecase interface {
	GetUserByID(c context.Context, id string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (string, error)
}

func (rtr *RefreshTokenRequest) Bind(r *http.Request) error {
	return util.ValidateStruct(rtr)
}
