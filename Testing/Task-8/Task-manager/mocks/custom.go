package mocks

import (
	domain "tskmgr/Domain"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNewUser() *domain.User {
	return &domain.User{
		UserID:   primitive.NewObjectID(),
		Username: "user1",
		Email:   "test.gmail.com",
		Password: "password1",
		UserRole: "user",
	}
}
