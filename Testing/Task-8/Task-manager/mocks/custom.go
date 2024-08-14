package mocks

import (
	domain "tskmgr/Domain"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNewUser() *domain.User {
	return &domain.User{
		UserID:   primitive.NewObjectID(),
		Username: "user1",
		Email:    "test.gmail.com",
		Password: "password1",
		UserRole: "user",
	}
}

// sample claim data
func GetNewClaim() *domain.Claims {
	return &domain.Claims{
		UserId:    primitive.NewObjectID(),
		UserEmail: "test.gmail.com",
		Username:  "user1",
		UserRole:  "user",
	}
}
