package controller

import (
	"company-micro/domain"
	"net/http"

	"github.com/go-chi/render"
)

func Get(w http.ResponseWriter, r *http.Request) {
	healthcheckResponse := domain.HealthcheckResponse{
		Status: "ok",
	}

	render.JSON(w, r, healthcheckResponse)
	return
}
