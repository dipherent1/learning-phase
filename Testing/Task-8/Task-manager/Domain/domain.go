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

//generate all the interfaces for user repository

type UserRepositoryInterface interface {
	GetByUsername(username string) (*User, error)
	Create(user *User) error
}

// generate all the interfaces for user usecase
type UserUsecaseInterface interface {
	CreateUser(user *User) error
	LogUser(user *User) (string, error)
}

// generate all the interfaces for task repository
type TaskRepositoryInterface interface {
	Create(task *Task) (*Task, error)
	GetByTitle(title string) (*Task, error)
	GetAllTasks() ([]Task, error)
	GetUserTasks(userid primitive.ObjectID) ([]Task, error)
	UpdateTask(title string, task *Task) (*Task, error)
	DeleteTask(title string) error
}

// gerate all the interfaces for task usecase
type TaskUsecaseInterface interface {
	CreateTask(task *Task) (*Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByTitle(title string) (*Task, error)
	GetUserTasks(userid primitive.ObjectID) ([]Task, error)
	UpdateTask(userrole string, userid primitive.ObjectID, title string, newtask *Task) (*Task, error)
	DeleteTask(userrole string, userid primitive.ObjectID, title string) error
}

// passowrd service interface
type PasswordServiceInterface interface {
	HashPassword(password string) (string, error)
	CheckPassword(existingPassword string, loginPassword string) error
}

// jwt service interface
type JwtServiceInterface interface {
	GetToken(claims *Claims) (string, error)
}
