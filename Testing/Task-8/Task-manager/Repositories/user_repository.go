package repositories

import (
	"context"
	domain "tskmgr/Domain"
	infrastructure "tskmgr/Infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDataManipulator struct {
	Data domain.Collection
}

func NewUserDataManipulator(coll domain.Collection) *UserDataManipulator {
	return &UserDataManipulator{Data: coll}
}

func (repo *UserDataManipulator) GetByUsername(username string) (*domain.User, error) {
	user := &domain.User{}
	filter := bson.M{"username": username}
	err := repo.Data.FindOne(context.TODO(), filter).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserDataManipulator) Create(user *domain.User) error {

	user.UserID = primitive.NewObjectID()

	hp, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hp

	_, err = repo.Data.InsertOne(context.TODO(), user)

	if err != nil {
		return err
	}

	return nil
}
