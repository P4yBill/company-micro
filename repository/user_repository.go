package repository

import (
	"company-micro/domain"
	"company-micro/mongodb"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongodb.Database
	collection string
}

func NewUserRepository(db mongodb.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []domain.User

	err = cursor.All(c, &users)
	if users == nil {
		return []domain.User{}, err
	}

	return users, err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	// user := domain.User{
	// 	Id:       primitive.NewObjectID(),
	// 	Name:     "vasiliss",
	// 	Password: "$2a$10$/cy/Alh7GOLre1uK.7faGO7Xm8MyKz1l49iHTG0w3VZV38LauTraK",
	// 	Roles:    []string{},
	// 	Email:    "v@g.com",
	// }
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	fmt.Println("INside getByEmail")
	fmt.Println(user)
	fmt.Println(err)

	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}
