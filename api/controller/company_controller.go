package controller

import (
	"company-micro/config"
	"company-micro/domain"
	"company-micro/mongodb"
	"company-micro/util"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	contentTypeJsonMergePatch = "application/merge-patch+json"
	contentyTypeJson          = "application/json"
)

type CompanyController struct {
	CompanyUsecase domain.CompanyUsecase
	Env            *config.Env
}

// Create Handler for Company POST endpoint.
// Creates a company based on the passed parameters
// If no error occurs and company is created successfully,
// Status code is set to 201("Created")
// adds a Location header with the location for the created resource.
func (cc *CompanyController) Create(w http.ResponseWriter, r *http.Request) {
	request := &domain.CompanyRequestCreate{}

	err := render.Bind(r, request)
	if err != nil {
		log.Println(err.Error())
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, "Please check your input"))
		return
	}

	_, err = cc.CompanyUsecase.GetByName(r.Context(), request.Name)
	if err == nil {
		message := fmt.Sprintf("Company with name: %s already exists, please use a different name", request.Name)
		log.Printf(message)
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, message))
		return
	}

	company := getCompanyFromCreateRequest(request)
	err = cc.CompanyUsecase.Create(r.Context(), &company)
	if err != nil {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	companyResponse := domain.CompanyCreateResponse{
		Response: domain.Response{
			Success: true,
			Message: "Company created successfully",
		},
		UUID: company.UUID,
	}

	resourceLocation := fmt.Sprintf("/api/v1/order/%s", company.Name)

	w.Header().Add("Location", resourceLocation)
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, companyResponse)
	return
}

// GetOne Handler for Company Get endpoint.
// Only returns one result as a response
func (cc *CompanyController) GetOne(w http.ResponseWriter, r *http.Request) {
	request := domain.CompanyRequestSingle{
		Name: chi.URLParam(r, "name"),
	}

	err := util.ValidateStruct(request)
	if err != nil {
		log.Println(err.Error())
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, "Please check your input. Maximum name length is 15"))
		return
	}
	log.Println(request)

	company, err := cc.CompanyUsecase.GetByName(r.Context(), request.Name)

	if err != nil {
		if mongodb.DocumentsNotFound(err) {
			// send Not Found status code if company was not found
			render.Render(w, r, domain.NewGeneralErrResponse(http.StatusNotFound, "Company not found."))
		} else {
			log.Println("Error while retrieving company: ", err.Error())
			render.Render(w, r, domain.NewErrResponse(http.StatusBadRequest, err, "Error while retrieve company"))
		}
		return
	}

	companyGetResponse := domain.CompanyGetResponse{
		Company: company,
	}

	render.JSON(w, r, companyGetResponse)
	return
}

// Patch Handler that updates a company
// Accepts a uri slug "name" to validate the company.
// If the company was not found, a 404 is returned instead.
func (cc *CompanyController) Patch(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-type") != contentTypeJsonMergePatch {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, "Please check your input"))
		return
	}

	// TODO: get rid of this and accept "application/merge-patch+json" content type
	hackContentType(r)

	// take name param from the uri
	request := &domain.CompanyRequestPatch{
		Name: chi.URLParam(r, "name"),
	}

	err := render.Bind(r, request)
	if err != nil {
		log.Println(err.Error())
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, "Please check your input"))
		return
	}

	company, err := cc.CompanyUsecase.GetByName(r.Context(), request.Name)
	if err != nil {
		message := fmt.Sprintf("Unable to find specified company", request.Name)
		log.Printf(message)
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusNotFound, message))
		return
	}

	updateResult, err := cc.CompanyUsecase.UpdateByName(r.Context(), request.Name, request.ToMongoPayload())
	if err != nil {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// send kafka event
	if updateResult.ModifiedCount == 1 {

	}

	request.UpdateCompany(&company)

	companyResponse := domain.CompanyGetResponse{
		Company: company,
	}

	render.JSON(w, r, companyResponse)
	return
}

// Delete Handler for Company DELETE endpoint.
// Deletes a company based on Uri param "name"
// If no error occurs and company is deleted successfully,
// Status code is set to 204("No Content")
// If company was not found in the database,
// returns status 404(Not Found).
func (cc *CompanyController) Delete(w http.ResponseWriter, r *http.Request) {
	request := domain.CompanyRequestSingle{
		Name: chi.URLParam(r, "name"),
	}

	err := util.ValidateStruct(request)
	if err != nil {
		log.Println(err.Error())
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusBadRequest, "Please check your input. Maximum name length is 15"))
		return
	}

	deletedCount, err := cc.CompanyUsecase.DeleteByName(r.Context(), request.Name)
	if err != nil {
		log.Println("Error while deleting company: ", err.Error())
		render.Render(w, r, domain.NewErrResponse(http.StatusBadRequest, err, "Error while deleted company"))
		return
	}

	// Send Not Found status code if company was not deleted
	if deletedCount == 0 {
		render.Render(w, r, domain.NewGeneralErrResponse(http.StatusNotFound, "Company not found"))
		return
	}

	// sends successful response with no content
	w.WriteHeader(http.StatusNoContent)
	return
}

func getCompanyFromCreateRequest(request *domain.CompanyRequestCreate) domain.Company {
	return domain.Company{
		Id:             primitive.NewObjectID(),
		UUID:           uuid.New().String(),
		Name:           request.Name,
		Description:    request.Description,
		EmployeesCount: request.EmployeesCount,
		Type:           request.Type,
	}
}

func patchHasEmptyBody(crp domain.CompanyRequestPatch) bool {

	if crp.Description == nil && crp.EmployeesCount == nil && crp.Registered == nil && crp.Type == nil {
		return true
	}

	return false
}

// hackContentType
// TODO: Remove this and accept merge patch json content type
func hackContentType(r *http.Request) {
	r.Header.Set("Content-type", contentyTypeJson)
}
