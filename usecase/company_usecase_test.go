package usecase_test

// import (
// 	"company-micro/domain"
// 	"company-micro/domain/mocks"
// 	"company-micro/usecase"
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func TestFetchByUserID(t *testing.T) {
// 	mockCompanyRepository := new(mocks.CompanyRepository)
// 	companyNameTest := "Test Company"

// 	t.Run("success", func(t *testing.T) {

// 		companyMock := domain.Company{
// 			Id:             primitive.NewObjectID(),
// 			UUID:           uuid.New().String(),
// 			Name:           companyNameTest,
// 			Type:           "",
// 			EmployeesCount: 1,
// 			Registered:     false,
// 		}

// 		mockCompanyRepository.On("FindByName", mock.Anything, companyMock.Name).Return(companyMock, nil).Once()

// 		u := usecase.NewCompanyUsecase(mockCompanyRepository, time.Second*2)

// 		company, err := u.GetByName(context.Background(), companyMock.Name)

// 		assert.NoError(t, err)
// 		assert.NotNil(t, companyMock)
// 		assert.Equal(t, company.Name, companyMock.Name)

// 		mockCompanyRepository.AssertExpectations(t)
// 	})

// 	t.Run("error", func(t *testing.T) {
// 		mockCompanyRepository.On("FindByName", mock.Anything, companyNameTest).Return(nil, errors.New("Unexpected")).Once()

// 		u := usecase.NewCompanyUsecase(mockCompanyRepository, time.Second*2)

// 		company, err := u.GetByName(context.Background(), companyNameTest)

// 		assert.Error(t, err)
// 		assert.Nil(t, company)

// 		mockCompanyRepository.AssertExpectations(t)
// 	})

// }
