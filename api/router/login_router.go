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

func NewLoginRouter(r chi.Router, env *config.Env, timeout time.Duration, db mongodb.Database) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	r.Post("/login", lc.Login)
}
