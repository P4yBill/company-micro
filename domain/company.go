package domain

import (
	"company-micro/util"
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionCompany = "company"
)

// should not be changed
var CompanyTypesArray = [4]string{"Corporations", "NonProfit", "Cooperative", "Sole Proprietorship"}

// Struct used for decoded mongodb related data
type Company struct {
	Id             primitive.ObjectID `bson:"_id" json:"-"`
	UUID           string             `bson:"uuid" json:"uuid"`
	Name           string             `bson:"name" json:"name"`
	Description    string             `bson:"description" json:"description"`
	EmployeesCount int                `bson:"employees_count" json:"employees_count"`
	Registered     bool               `bson:"registered" json:"registered"`
	Type           string             `bson:"type" json:"type"`
}

type CompanyRepository interface {
	Create(c context.Context, company *Company) error
	FindByUUID(c context.Context, companyUUID string) (Company, error)
	FindByName(c context.Context, name string) (Company, error)
	DeleteByName(c context.Context, name string) (int64, error)
	UpdateByName(c context.Context, name string, p *CompanyPatchMongoPayload) (*mongo.UpdateResult, error)
}

type CompanyUsecase interface {
	Create(c context.Context, company *Company) error
	GetByUUID(c context.Context, companyUUID string) (Company, error)
	GetByName(c context.Context, name string) (Company, error)
	DeleteByName(c context.Context, name string) (int64, error)
	UpdateByName(c context.Context, name string, p *CompanyPatchMongoPayload) (*mongo.UpdateResult, error)
}

type CompanyCreateResponse struct {
	Response
	UUID string `json:"uuid"`
}

type CompanyGetResponse struct {
	Company
}

// Used for validating the POST Create Company parameters
type CompanyRequestCreate struct {
	Name           string `form:"name" validate:"required,max=15"`
	Description    string `form:"description" validate:"omitempty,max=3000"`
	EmployeesCount int    `form:"employees_count" validate:"required,number"`
	Registered     *bool  `form:"registered" validate:"required,boolean"`
	Type           string `form:"type" validate:"required,is-company-type"`
}

// Used for validating the post parameters
type CompanyRequestPatch struct {
	Name           string  `json:"name" validate:"required,max=15"`
	Description    *string `json:"description" validate:"omitempty,required,max=3000"`
	EmployeesCount *int    `json:"employees_count" validate:"omitempty,required,number,gt=0"`
	Registered     *bool   `json:"registered" validate:"omitempty,required,boolean"`
	Type           *string `json:"type" validate:"omitempty,required,is-company-type"`
}

type CompanyPatchMongoPayload struct {
	Description    *string
	EmployeesCount *int
	Registered     *bool
	Type           *string
}

type CompanyRequestSingle struct {
	Name string `validate:"required,max=15"`
}

func (crc *CompanyRequestCreate) Bind(r *http.Request) error {
	return util.ValidateStruct(crc)
}

func (crp *CompanyRequestPatch) Bind(r *http.Request) error {
	return util.ValidateStruct(crp)
}

func (crp *CompanyRequestPatch) ToMongoPayload() *CompanyPatchMongoPayload {
	return &CompanyPatchMongoPayload{
		Description:    crp.Description,
		EmployeesCount: crp.EmployeesCount,
		Registered:     crp.Registered,
		Type:           crp.Type,
	}
}

// TODO: Build generic function to update only non-nil fields
func (crp *CompanyRequestPatch) UpdateCompany(company *Company) {
	if crp.Description != nil {
		company.Description = *crp.Description
	}

	if crp.EmployeesCount != nil {
		company.EmployeesCount = *crp.EmployeesCount
	}

	if crp.Registered != nil {
		company.Registered = *crp.Registered
	}

	if crp.Type != nil {
		company.Type = *crp.Type
	}
}
