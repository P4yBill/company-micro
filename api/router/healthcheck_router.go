package router

import (
	"company-micro/api/controller"

	"github.com/go-chi/chi/v5"
)

func NewHealthcheckRouter(r chi.Router) {
	r.Get("/healthcheck", controller.Get)
}
