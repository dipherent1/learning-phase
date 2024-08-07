package data

import (
	"context"
	"errors"
	"log"
	"os"
	"time"
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

func (d *UserData) userexists(email string) (bool, error) {
	var user models.User
	filter := bson.M{"email": email}
	err := d.Users.FindOne(context.TODO(), filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return false, nil // No user found with this email
	} else if err != nil {
		return false, err // Other errors
	}

	return true, nil // User found
}

func (d *UserData) SignupService(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Internal server error")
	}

	exists, err := d.userexists(user.Email)

	if err != nil {

		return err
	}

	if exists {

		return errors.New("user with this email already exists")
	}

	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)
	_, err = d.Users.InsertOne(context.TODO(), user)

	if err != nil {
		log.Fatal(err.Error())
	}

	return nil

}

func (d *UserData) LoginService(user *models.User) (string, error) {
	var exisitingUser models.User
	filter := bson.M{"username": user.UserName}
	err := d.Users.FindOne(context.TODO(), filter).Decode(&exisitingUser)

	if err != nil {
		return "", err
	}

	if bcrypt.CompareHashAndPassword([]byte(exisitingUser.Password), []byte(user.Password)) != nil {
		return "", errors.New("invalid password")

	}

	claims := models.Claims{
		UserId:    exisitingUser.ID,
		UserEmail: exisitingUser.Email,
		Username:  exisitingUser.UserName,
		UserRole:  exisitingUser.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	key := os.Getenv("JWT_SECRET")
	jwtSecret := []byte(key)

	if jwtSecret == nil {
		return "", errors.New("JWT_SECRET environment variable not set")
	}

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (d *UserData) GetAllUsers() ([]models.User, error) {
	cursor, err := d.Users.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil

}
func (uc *UserData) ChangeRoleUser(username string) {
	// filter := bson

}
