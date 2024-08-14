package usecases_test

import (
	"errors"
	"testing"
	infrastructure "tskmgr/Infrastructure"
	usecases "tskmgr/Usecases"
	"tskmgr/mocks"

	"github.com/stretchr/testify/suite"
)

// generate test suite for user usecases
type UserUsecaseTestSuite struct {
	suite.Suite
	mockrepo    *mocks.UserRepositoryInterface
	userusecase *usecases.UserUsecase
}

// setup test suite
func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.mockrepo = new(mocks.UserRepositoryInterface)
	suite.userusecase = usecases.NewUserUsecase(suite.mockrepo)
}

// test CreateUser
func (suite *UserUsecaseTestSuite) TestCreateUser() {
	// test case success
	suite.Run("TestCreateUser_Success", func() {
		user := mocks.GetNewUser()
		suite.mockrepo.On("GetByUsername", user.Username).Return(nil, errors.New("not found")).Once()
		suite.mockrepo.On("Create", user).Return(nil)

		err := suite.userusecase.CreateUser(user)
		suite.Nil(err)
	})

	// test case user already exists
	suite.Run("TestCreateUser_UserExists", func() {
		user := mocks.GetNewUser()
		suite.mockrepo.On("GetByUsername", user.Username).Return(user, nil).Once()

		err := suite.userusecase.CreateUser(user)
		suite.NotNil(err)
	})
}

// test LogUser
func (suite *UserUsecaseTestSuite) TestLogUser() {
	// test case success
	suite.Run("TestLogUser_Success", func() {
		user := mocks.GetNewUser()
		suite.mockrepo.On("GetByUsername", user.Username).Return(user, nil).Once()

		infrastructure.CheckPassword = func(existingPassword string, loginPassword string) error {
			return nil
		}
		
		// suite.userusecase.jwtservice.On("GenerateToken", mocks.GetNewClaim() ).Return("", errors.New("error"))
		
		token, err := suite.userusecase.LogUser(user)
		suite.Nil(err)
		suite.NotNil(token)
	})

	// test case user not found
	suite.Run("TestLogUser_UserNotFound", func() {
		user := mocks.GetNewUser()
		suite.mockrepo.On("GetByUsername", user.Username).Return(nil, errors.New("not found")).Once()

		token, err := suite.userusecase.LogUser(user)
		suite.NotNil(err)
		suite.Empty(token)
	})

	// test case invalid password
	suite.Run("TestLogUser_InvalidPassword", func() {
		user := mocks.GetNewUser()
		suite.mockrepo.On("GetByUsername", user.Username).Return(user, nil).Once()

		infrastructure.CheckPassword = func(existingPassword string, loginPassword string) error {
			return errors.New("invalid password")
		}


		token, err := suite.userusecase.LogUser(user)
		suite.NotNil(err)
		suite.Empty(token)
	})
}

// setup test suite
func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
