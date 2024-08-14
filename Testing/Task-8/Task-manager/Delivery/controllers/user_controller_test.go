package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"
	"tskmgr/Delivery/controllers"
	"tskmgr/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

// create user test suite
type UserCtrlTestSuite struct {
	suite.Suite
	ctrl        *controllers.Usercontroller
	mockUsecase *mocks.UserUsecaseInterface
}

// setup test suite
func (suite *UserCtrlTestSuite) SetupTest() {
	suite.mockUsecase = new(mocks.UserUsecaseInterface)
	suite.ctrl = controllers.NewUsercontroller(suite.mockUsecase)
}

// tear down test suite
func (suite *UserCtrlTestSuite) TearDownTest() {
	suite.mockUsecase.AssertExpectations(suite.T())
}
// test signup suite
func (suite *UserCtrlTestSuite) TestSignupController() {
	// test case success
	suite.Run("TestSignupController_succes", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		user := mocks.GetNewUser()

		suite.mockUsecase.On("CreateUser", user).Return(nil)

		body, err := json.Marshal(user)
		suite.Nil(err)
		ctx.Request = httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))

		suite.ctrl.SignupController(ctx)
		suite.Nil(err)

		suite.Equal(201, w.Code)
		expectedResponse := `{"message":"user registered successfully"}`
		suite.Equal(expectedResponse, w.Body.String())
	})

	// A testcase for an invalid request.
	suite.Run("RegisterUser_InvalidRequest", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("POST", "/users", nil)

		suite.ctrl.SignupController(ctx)
		expected, err := json.Marshal(gin.H{"error": "Invalid request payload"})
		suite.Nil(err)

		suite.Equal(400, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})

	// test case failure
	suite.Run("TestSignupController_failure", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		user := mocks.GetNewUser()

		suite.mockUsecase.On("CreateUser", user).Return(errors.New("error"))

		body, err := json.Marshal(user)
		suite.Nil(err)
		ctx.Request = httptest.NewRequest("POST", "/users", bytes.NewReader(body))

		suite.ctrl.SignupController(ctx)

		suite.Equal(400, w.Code)
		expected, _ := json.Marshal(gin.H{"error": "error"})
		suite.Equal(string(expected), w.Body.String())
	})

}

// test logincontroller suite
func (suite *UserCtrlTestSuite) TestLoginController() {
	// test case success
	suite.Run("TestLoginController_success", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		user := mocks.GetNewUser()

		suite.mockUsecase.On("LogUser", user).Return("token", nil)

		body, err := json.Marshal(user)
		suite.Nil(err)
		ctx.Request = httptest.NewRequest("POST", "/users/login", bytes.NewBuffer(body))

		suite.ctrl.LoginController(ctx)

		suite.Equal(200, w.Code)
		expectedResponse := `{"message":"user logged in successfully","token":"token"}`
		suite.Equal(expectedResponse, w.Body.String())
	})

	// A testcase for an invalid request.
	suite.Run("LoginUser_InvalidRequest", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("POST", "/users/login", nil)

		suite.ctrl.LoginController(ctx)
		expected, err := json.Marshal(gin.H{"error": "Invalid request payload"})
		suite.Nil(err)

		suite.Equal(400, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})

	// test case failure
	suite.Run("TestLoginController_failure", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		user := mocks.GetNewUser()

		suite.mockUsecase.On("LogUser", user).Return("", errors.New("error"))

		body, err := json.Marshal(user)
		suite.Nil(err)
		ctx.Request = httptest.NewRequest("POST", "/users/login", bytes.NewReader(body))

		suite.ctrl.LoginController(ctx)

		suite.Equal(400, w.Code)
	})

}

// run the test suite
func Test_UserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserCtrlTestSuite))
}
