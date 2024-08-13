package repositories_test

import (
	"errors"
	"testing"
	domain "tskmgr/Domain"
	infrastructure "tskmgr/Infrastructure"
	repositories "tskmgr/Repositories"
	"tskmgr/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

var mockUser = mock.AnythingOfType("*domain.User")

// create user test suite
type UserRepoTestSuite struct {
	suite.Suite
	repo       *repositories.UserDataManipulator
	collection *mocks.Collection
}

// Test the GetByUsername method
func (suite *UserRepoTestSuite) TestGetByUsername() {
	// success case
	suite.Run("GetByUsername success", func() {
		user := mocks.GetNewUser()
		username := user.Username
		user.Password = "hashedpassword123"

		res := new(mocks.SingleResult)
		res.On("Decode", mock.Anything).Return(nil).Once().Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*domain.User)
			*userPtr = *user
		})

		suite.collection.On("FindOne", mock.Anything, mock.Anything).Return(res).Once()

		result, err := suite.repo.GetByUsername(username)

		suite.NoError(err)
		suite.Equal(user, result)
	})

	// failure case
	suite.Run("GetUserByUsername_Failure", func() {
		username := "nonexistent"

		res := new(mocks.SingleResult)
		res.On("Decode", mockUser).Return(mongo.ErrNoDocuments).Once()
		suite.collection.On("FindOne", mock.Anything, mock.Anything).Return(res).Once()

		result, err := suite.repo.GetByUsername(username)
		suite.Error(err)
		suite.Nil(result)
	})

}

// Test the Create method
func (suite *UserRepoTestSuite) TestCreate() {
	// success case
	suite.Run("Create success", func() {
		user := mocks.GetNewUser()
		user.Password = "hashedpassword123"

		suite.collection.On("InsertOne", mock.Anything, user).Return(&mongo.InsertOneResult{}, nil).Once()
		err := suite.repo.Create(user)

		suite.NoError(err)

	})

	// insertion failure case
	suite.Run("AddUser_Failure", func() {
		user := mocks.GetNewUser()
		suite.collection.On("InsertOne", mock.Anything, user).Return(&mongo.InsertOneResult{}, mongo.ErrClientDisconnected).Once()
		
		
		err := suite.repo.Create(user)
		suite.Error(err)
	})
	
	// hashing failure case
	suite.Run("AddUser_Failure_by_hash", func() {
		user := mocks.GetNewUser()
		infrastructure.HashPassword = func(password string) (string, error) {
			return "", errors.New("hash error")
		}

		err := suite.repo.Create(user)
		suite.Error(err)
	})

}

// setup test suite
func (suite *UserRepoTestSuite) SetupTest() {
	suite.collection = new(mocks.Collection)
	suite.repo = repositories.NewUserDataManipulator(suite.collection)

	infrastructure.HashPassword = func(password string) (string, error) {
		return "hashedpassword123", nil
	}
}

// A method that finalizes the test suite.
func (suite *UserRepoTestSuite) TearDownTest() {
	suite.collection.AssertExpectations(suite.T())
}

// A method that runs the test suite.
func TestMongoUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
