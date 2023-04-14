package usecase

import (
	"company-micro/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type companyUsecase struct {
	companyRepository domain.CompanyRepository
	contextTimeout    time.Duration
}

func NewCompanyUsecase(companyRepository domain.CompanyRepository, timeout time.Duration) domain.CompanyUsecase {
	return &companyUsecase{
		companyRepository: companyRepository,
		contextTimeout:    timeout,
	}
}

func (cu *companyUsecase) Create(c context.Context, company *domain.Company) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.companyRepository.Create(ctx, company)
}

func (cu *companyUsecase) GetByUUID(c context.Context, companyUUID string) (domain.Company, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.companyRepository.FindByUUID(ctx, companyUUID)
}

func (cu *companyUsecase) GetByName(c context.Context, name string) (domain.Company, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.companyRepository.FindByName(ctx, name)
}

func (cu *companyUsecase) DeleteByName(c context.Context, name string) (int64, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.companyRepository.DeleteByName(ctx, name)
}

func (cu *companyUsecase) UpdateByName(c context.Context, name string, p *domain.CompanyPatchMongoPayload) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.companyRepository.UpdateByName(ctx, name, p)
}
