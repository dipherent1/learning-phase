package usecases

import (
	"errors"
	"time"
	domain "tskmgr/Domain"
	infrastructure "tskmgr/Infrastructure"

	"github.com/golang-jwt/jwt"
)

type UserUsecase struct {
	MyUserRepo domain.UserRepositoryInterface
}

func NewUserUsecase(repo domain.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{
		MyUserRepo: repo,
	}
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	//check if user already exists
	_, err := u.MyUserRepo.GetByUsername(user.Username)
	if err == nil {
		return errors.New("user already exists")
	}
	err = u.MyUserRepo.Create(user)
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
