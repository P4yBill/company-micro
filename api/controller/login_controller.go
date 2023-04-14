package controller

import (
	"company-micro/api/middleware"
	"company-micro/config"
	"company-micro/domain"
	"company-micro/util"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *config.Env
}

func (lc *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	request := &domain.LoginRequest{}

	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, "Bad Credentials given"))
		return
	}
	fmt.Println(request)

	user, err := lc.LoginUsecase.GetUserByEmail(r.Context(), request.Email)
	if err != nil {
		log.Println("Error while retrieving user: ", err.Error())
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, "Wrong email or password. Please use different credentials"))
		return
	}
	fmt.Println(user)

	if util.ComparePassword(user.Password, request.Password) != nil {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, "Wrong email or password. Please use different credentials"))
		return
	}
	fmt.Println("user2")

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusInternalServerError, "Something went wrong while creating token"))
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusInternalServerError, "Something went wrong while creating token"))
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		TokenType:    middleware.TokenType,
		ExpiresIn:    strconv.Itoa(lc.Env.AccessTokenExpiryHour * 3600),
		RefreshToken: refreshToken,
		Username:     user.Username,
	}

	render.JSON(w, r, loginResponse)
	return
}
