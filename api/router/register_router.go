package router

import (
	"company-micro/api/controller"
	"company-micro/config"
	"company-micro/domain"
	"company-micro/mongodb"
	"company-micro/repository"
	"company-micro/usecase"
	"time"

	"github.com/go-chi/chi/v5"
)

func NewRegisterRouter(r chi.Router, env *config.Env, timeout time.Duration, db mongodb.Database) {
	userRepository := repository.NewUserRepository(db, domain.CollectionUser)
	registerController := &controller.RegisterController{
		RegisterUsecase: usecase.NewRegisterUsecase(userRepository, timeout),
		Env:             env,
	}
	r.Post("/register", registerController.Register)
}
