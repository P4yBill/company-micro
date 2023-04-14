package router

import (
	"company-micro/api/controller"
	"company-micro/api/middleware"
	"company-micro/config"
	"company-micro/domain"
	"company-micro/mongodb"
	"company-micro/repository"
	"company-micro/usecase"
	"time"

	"github.com/go-chi/chi/v5"
)

func NewCompanyRouter(r chi.Router, env *config.Env, timeout time.Duration, db mongodb.Database, authMiddleware middleware.Middleware) {
	companyRepository := repository.NewCompanyRepository(db, domain.CollectionCompany)
	companyController := &controller.CompanyController{
		CompanyUsecase: usecase.NewCompanyUsecase(companyRepository, timeout),
		Env:            env,
	}
	companyRouter := chi.NewRouter()
	//public routes
	companyRouter.Group(func(r chi.Router) {
		r.Get("/{name}", companyController.GetOne)
	})

	//require authentication
	companyRouter.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Post("/", companyController.Create)
		r.Delete("/{name}", companyController.Delete)
		r.Patch("/{name}", companyController.Patch)
	})

	r.Mount("/company", companyRouter)
}
