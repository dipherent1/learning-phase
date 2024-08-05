package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"tskmgr/models"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	Users *mongo.Collection
}

func NewUserData(coll *mongo.Collection) *UserData {
	return &UserData{Users: coll}
}

func (d *UserData) SignupService(user *models.User) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Internal server error")
	}
	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)
	_, err = d.Users.InsertOne(context.TODO(), user)

	if err != nil {
		log.Fatal(err.Error())
	}

}

func (d *UserData) LoginService(user *models.User) string {
	var exisitingUser models.User
	filter := bson.M{"email": user.Email}
	err := d.Users.FindOne(context.TODO(), filter).Decode(&exisitingUser)

	if err != nil {
		fmt.Println("---------")
		log.Fatal(err.Error())
	}

	if bcrypt.CompareHashAndPassword([]byte(exisitingUser.Password), []byte(user.Password)) != nil {
		fmt.Println("++++++++")
		log.Fatal(err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       exisitingUser.ID,
		"email":    exisitingUser.Email,
		"username": exisitingUser.UserName,
	})

	key := os.Getenv("JWT_SECRET")
	jwtSecret := []byte(key)

	if jwtSecret == nil {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		fmt.Println("==========")
		log.Fatal(err.Error())
	}

	return jwtToken
}
