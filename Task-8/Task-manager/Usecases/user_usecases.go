package usecases

import (
	domain "tskmgr/Domain"
	repositories "tskmgr/Repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type Usecase struct {
	Myrepo *repositories.DataMinipulator
}

func NewUserUsecase(coll *mongo.Collection) *Usecase {
	return &Usecase{
		Myrepo: repositories.NewDataMinipulator(coll),
	}
}

func (u *Usecase) CreateUser(user *domain.User) error {
	err := u.Myrepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}
