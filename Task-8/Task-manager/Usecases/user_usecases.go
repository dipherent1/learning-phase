package usecases

import (
	"time"
	domain "tskmgr/Domain"
	infrastructure "tskmgr/Infrastructure"
	repositories "tskmgr/Repositories"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserUsecase struct {
	MyUserRepo *repositories.UserDataManipulator
}

func NewUserUsecase(coll *mongo.Collection) *UserUsecase {
	return &UserUsecase{
		MyUserRepo: repositories.NewUserDataManipulator(coll),
	}
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	err := u.MyUserRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) LogUser(user *domain.User) (string, error) {
	existinguser, err := u.MyUserRepo.GetByUsername(user.Username)
	if err != nil {
		return "", err
	}

	err = infrastructure.CheckPassword(existinguser.Password, user.Password)
	if err != nil {
		return "", err
	}

	claims := domain.Claims{
		UserId:    existinguser.UserID,
		UserEmail: existinguser.Email,
		Username:  existinguser.Username,
		UserRole:  existinguser.UserRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token, err := infrastructure.GetToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
