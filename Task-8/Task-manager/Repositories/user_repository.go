package repositories

import (
	"context"
	"errors"
	domain "tskmgr/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataMinipulator struct{
	Data *mongo.Collection
}

func NewDataMinipulator (coll *mongo.Collection) *DataMinipulator{
	return &DataMinipulator{Data : coll}
}

func (repo *DataMinipulator) GetByUsername(username string) (*domain.User,error){
	user:= &domain.User{}
	filter := bson.M{"username":username}
	err := repo.Data.FindOne(context.TODO(), filter).Decode(user)

	if err!=nil{
		return nil,err
	}

	return user,nil
}

func (repo *DataMinipulator) Create(user *domain.User) error{
	existinguser,err :=repo.GetByUsername(user.Username)
	user.UserID = primitive.NewObjectID()

	if err!=nil {
		return err
	}

	if existinguser != nil {
		return errors.New("username already exists")
	}
	
	_,err = repo.Data.InsertOne(context.TODO(),user)
	
	if err!=nil {
		return err
	}

	return nil
}