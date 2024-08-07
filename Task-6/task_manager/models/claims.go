package models

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	UserId    primitive.ObjectID `json:"userid"`
	UserEmail string             `json:"useremail"`
	Username  string             `json:"username"`
	UserRole  string             `json:"userrole"`
	jwt.StandardClaims
}
