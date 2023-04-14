package util

import (
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
)

const (
	companyValidator = "is-company-type"
)

var companyTypesArray = []string{"Corporations", "NonProfit", "Cooperative", "Sole Proprietorship"}

var lock = &sync.Mutex{}
var validateInstance *validator.Validate

// GetValidateInstance Singleton that returns a validator.Validate instance
// validator.Validate is thread safe https://github.com/go-playground/validator/issues/315
// TODO: Refactor to inject this
func GetValidateInstance() *validator.Validate {
	if validateInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if validateInstance == nil {
			validateInstance = validator.New()
			validateInstance.RegisterValidation(companyValidator, companyTypeValidator)
		}
	}

	return validateInstance
}

// ValidateStruct validates the specified struct
// based on the validate tags
func ValidateStruct(s interface{}) error {
	validate := GetValidateInstance()
	err := validate.Struct(s)

	return err
}

// Validates Company Types implements validator.Func
func companyTypeValidator(fl validator.FieldLevel) bool {
	return slices.Contains(companyTypesArray, fl.Field().String())
}

func IsStringBlank(str string) bool {
	if len(strings.Trim(str, " ")) == 0 {
		return true
	}

	return false
}
