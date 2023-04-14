package controller

import (
	"company-micro/api/middleware"
	"company-micro/config"
	"company-micro/domain"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type RefreshTokenController struct {
	RefreshTokenUsecase domain.RefreshTokenUsecase
	Env                 *config.Env
}

func (rtc *RefreshTokenController) Refresh(w http.ResponseWriter, r *http.Request) {
	request := &domain.RefreshTokenRequest{}

	err := render.Bind(r, request)
	if err != nil {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, "All params are required"))
		return
	}
	fmt.Println(request)

	if request.GrantType != "refresh_token" {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, "Could not refresh token"))
		return
	}

	id, err := rtc.RefreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, rtc.Env.RefreshTokenSecret)
	if err != nil {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, "User not found"))
		return
	}

	user, err := rtc.RefreshTokenUsecase.GetUserByID(r.Context(), id)
	if err != nil {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, "User not found"))
		return
	}

	accessToken, err := rtc.RefreshTokenUsecase.CreateAccessToken(&user, rtc.Env.AccessTokenSecret, rtc.Env.AccessTokenExpiryHour)
	if err != nil {
		render.Render(w, r, domain.NewErrResponse(http.StatusInternalServerError, err, "Could not refresh token"))
		return
	}

	refreshToken, err := rtc.RefreshTokenUsecase.CreateRefreshToken(&user, rtc.Env.RefreshTokenSecret, rtc.Env.RefreshTokenExpiryHour)
	if err != nil {
		render.Render(w, r, domain.NewErrResponse(http.StatusInternalServerError, err, "Could not refresh token"))
		return
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		TokenType:    middleware.TokenType,
		ExpiresIn:    strconv.Itoa(rtc.Env.AccessTokenExpiryHour * 3600),
		RefreshToken: refreshToken,
	}

	render.JSON(w, r, refreshTokenResponse)
	return
}
