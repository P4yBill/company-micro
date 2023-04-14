package router

import (
	"company-micro/api/middleware"
	"company-micro/config"
	"company-micro/domain"
	"company-micro/mongodb"
	"company-micro/repository"

	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func GetRouter(env *config.Env, timeout time.Duration, db mongodb.Database) *chi.Mux {
	r := chi.NewRouter()

	ur := repository.NewUserRepository(db, domain.CollectionUser)
	authMiddleware := middleware.JwtAuthenticator(env.AccessTokenSecret, domain.UserCtxIdKey, ur)

	r.Use(chiMiddleware.Logger)

	// Public routes
	NewHealthcheckRouter(r)
	NewLoginRouter(r, env, timeout, db)
	NewRegisterRouter(r, env, timeout, db)
	NewRefreshTokenRouter(r, env, timeout, db)

	// /api/v1/ router
	apiRouter := chi.NewRouter()

	NewCompanyRouter(apiRouter, env, timeout, db, authMiddleware)

	r.Mount("/api/v1", apiRouter)
	return r
}
