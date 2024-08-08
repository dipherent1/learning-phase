package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct{
	UserID   primitive.ObjectID `json:"userID" bson:"_id" `
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	UserRole string             `json:"userrole"`
}
