package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
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

// test signup suite
func (suite *UserCtrlTestSuite) TestSignupController() {
	// test case success
	suite.Run("TestSignupController_succes", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		user := mocks.GetNewUser()
		
		suite.mockUsecase.On("CreateUser", &user).Return(nil)
		
		body, err := json.Marshal(user)
		suite.Nil(err)
		ctx.Request = httptest.NewRequest("POST", "/users", bytes.NewReader(body))

		suite.ctrl.SignupController(ctx)
		expected, err := json.Marshal(user)
		suite.Nil(err)

		suite.Equal(201, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})

}
