package controller

import (
	"company-micro/config"
	"company-micro/domain"
	"company-micro/util"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterController struct {
	RegisterUsecase domain.RegisterUsecase
	Env             *config.Env
}

func (rc *RegisterController) Register(w http.ResponseWriter, r *http.Request) {
	request := &domain.RegisterRequest{}

	if err := render.Bind(r, request); err != nil {
		log.Println("Error while binding request: " + err.Error())
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, err.Error()))
		return
	}
	fmt.Println(request)

	_, err := rc.RegisterUsecase.GetUserByEmail(r.Context(), request.Email)
	if err == nil {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusConflict, "User already exists with the given email."))
		return
	}

	encryptedPassword, err := util.HashPassword(request.Password)
	if err != nil {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusInternalServerError, err.Error()))

		return
	}

	user := domain.User{
		Id:       primitive.NewObjectID(),
		Username: request.Username,
		Email:    request.Email,
		Password: encryptedPassword,
		Roles:    []string{"USER"},
	}

	err = rc.RegisterUsecase.Create(r.Context(), &user)
	if err != nil {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	registerResponse := domain.RegisterResponse{
		Response: domain.Response{
			Success: true,
			Message: "Register successful",
		},
	}

	render.JSON(w, r, registerResponse)
	return
}
