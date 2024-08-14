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
		UserEmail: "test@gmail.com",
		Username:  "test",
		UserRole:  "user",
	}
}

// get sample task data
func GetNewTask() *domain.Task {
	return &domain.Task{
		Id:          primitive.NewObjectID(),
		UserId:      primitive.NewObjectID(),
		Title:       "test",
		Description: "test description",
		Priority:    "high",
		Status:      "pending",
	}
}

// get multiple sample task data
func GetMultipleTask() []domain.Task {
	return []domain.Task{
		{
			Id:          primitive.NewObjectID(),
			UserId:      primitive.NewObjectID(),
			Title:       "test1",
			Description: "test description",
			Priority:    "high",
			Status:      "pending",
		},
		{
			Id:          primitive.NewObjectID(),
			UserId:      primitive.NewObjectID(),
			Title:       "test2",
			Description: "test description",
			Priority:    "high",
			Status:      "pending",
		},
	}
}
