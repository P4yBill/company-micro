package repository

import (
	"company-micro/domain"
	"company-micro/mongodb"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"context"
)

type companyRepository struct {
	database   mongodb.Database
	collection string
}

func NewCompanyRepository(db mongodb.Database, collection string) domain.CompanyRepository {
	return &companyRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *companyRepository) Create(c context.Context, company *domain.Company) error {
	collection := cr.database.Collection(cr.collection)

	_, err := collection.InsertOne(c, company)

	return err
}

func (cr *companyRepository) FindByUUID(c context.Context, companyUUID string) (domain.Company, error) {
	collection := cr.database.Collection(cr.collection)

	var company domain.Company
	_, err := uuid.Parse(companyUUID)
	if err != nil {
		return company, err
	}

	err = collection.FindOne(c, bson.M{"uuid": companyUUID}).Decode(&company)
	return company, err
}

func (cr *companyRepository) FindByName(c context.Context, name string) (domain.Company, error) {
	collection := cr.database.Collection(cr.collection)

	var company domain.Company
	err := collection.FindOne(c, bson.M{"name": name}).Decode(&company)

	return company, err
}

func (cr *companyRepository) DeleteByName(c context.Context, name string) (int64, error) {
	collection := cr.database.Collection(cr.collection)

	deletedCount, err := collection.DeleteOne(c, bson.M{"name": name})

	return deletedCount, err
}

func (cr *companyRepository) UpdateByName(c context.Context, name string, p *domain.CompanyPatchMongoPayload) (*mongo.UpdateResult, error) {
	collection := cr.database.Collection(cr.collection)
	filter := bson.D{{Key: "name", Value: name}}

	update := bson.D{{Key: "$set", Value: getUpdateMongoPayload(p)}}

	updateResult, err := collection.UpdateOne(c, filter, update)
	return updateResult, err
}

// Make generic and use reflect to iterate
func getUpdateMongoPayload(p *domain.CompanyPatchMongoPayload) bson.D {
	payloadSlice := bson.D{}

	if p.Description != nil {
		payloadSlice = append(payloadSlice, bson.E{Key: "description", Value: *p.Description})
	}

	if p.EmployeesCount != nil {
		payloadSlice = append(payloadSlice, bson.E{Key: "employees_count", Value: *p.EmployeesCount})
	}

	if p.Registered != nil {
		payloadSlice = append(payloadSlice, bson.E{Key: "registered", Value: *p.Registered})
	}

	if p.Type != nil {
		payloadSlice = append(payloadSlice, bson.E{Key: "type", Value: *p.Type})
	}

	fmt.Println(payloadSlice)

	return payloadSlice
}
