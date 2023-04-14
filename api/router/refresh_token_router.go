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

func NewRefreshTokenRouter(r chi.Router, env *config.Env, timeout time.Duration, db mongodb.Database) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	r.Post("/refresh", rtc.Refresh)
}
