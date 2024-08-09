package domain

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID   primitive.ObjectID `json:"userID" bson:"_id" `
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	UserRole string             `json:"userrole"`
}

type Claims struct {
	UserId    primitive.ObjectID `json:"userid"`
	UserEmail string             `json:"useremail"`
	Username  string             `json:"username"`
	UserRole  string             `json:"userrole"`
	jwt.StandardClaims
}

type Task struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	UserId      primitive.ObjectID `json:"userid" bson:"userid"`
	Title       string             `json:"title" bson:"title"`             // Title of the task
	Description string             `json:"description" bson:"description"` // Description of the task
	Priority    string             `json:"priority" bson:"priority"`       // Priority level of the task
	Status      string             `json:"status" bson:"status"`           // Current status of the task (e.g., "pending", "completed")
}
