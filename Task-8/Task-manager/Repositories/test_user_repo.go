package repositories_test

import (
	repositories "tskmgr/Repositories"

	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	UserRepo *repositories.UserDataManipulator
	TaskRepo *repositories.TaskDataManipulator
}
