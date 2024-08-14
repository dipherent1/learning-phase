package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
	"tskmgr/Delivery/controllers"
	"tskmgr/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

// generate suite for task controller
type TaskControllerSuite struct {
	suite.Suite
	Controller *controllers.TaskController
	usecase    *mocks.TaskUsecaseInterface
}

// setup the test suite
func (suite *TaskControllerSuite) SetupTest() {
	suite.usecase = &mocks.TaskUsecaseInterface{}
	suite.Controller = controllers.NewTaskController(suite.usecase)
}

// test create task
func (suite *TaskControllerSuite) TestCreateTask() {
	// test case success
	suite.Run("TestCreateTask_Success", func() {
		task := mocks.GetNewTask()
		claim := *mocks.GetNewClaim()

		newtask := *task

		newtask.UserId = claim.UserId
		// suite.Equal(newtask, *task)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Set("claim", claim)

		body, err := json.Marshal(task)
		suite.Nil(err)

		ctx.Request = httptest.NewRequest("POST", "/task", strings.NewReader(string(body)))

		suite.usecase.On("CreateTask", &newtask).Return(&newtask, nil).Once()

		suite.Controller.CreateTask(ctx)

		expected, err := json.Marshal(newtask)
		suite.Nil(err)

		suite.Equal(201, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
	// test case invalid json
	suite.Run("TestCreateTask_InvalidJSON", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Set("claim", *mocks.GetNewClaim())

		ctx.Request = httptest.NewRequest("POST", "/task", strings.NewReader("invalid json"))

		suite.Controller.CreateTask(ctx)

		expected, err := json.Marshal(gin.H{"error": "Invalid JSON provided"})
		suite.Nil(err)

		suite.Equal(400, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
	// test case invalid claim
	suite.Run("TestCreateTask_InvalidClaim", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("POST", "/task", nil)

		suite.Controller.CreateTask(ctx)

		expected, err := json.Marshal(gin.H{"error": "claim not set"})
		suite.Nil(err)

		suite.Equal(401, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
	// test case failure
	suite.Run("TestCreateTask_Failure", func() {
		task := mocks.GetNewTask()
		claim := *mocks.GetNewClaim()

		newtask := *task

		newtask.UserId = claim.UserId

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Set("claim", claim)

		body, err := json.Marshal(task)
		suite.Nil(err)

		ctx.Request = httptest.NewRequest("POST", "/task", strings.NewReader(string(body)))

		suite.usecase.On("CreateTask", &newtask).Return(nil, errors.New("error")).Once()

		suite.Controller.CreateTask(ctx)

		expected, err := json.Marshal(gin.H{"error": "error"})
		suite.Nil(err)

		suite.Equal(401, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
}

// test get task by title
func (suite *TaskControllerSuite) TestGetTaskByTitle() {
	// test case success
	suite.Run("TestGetTaskByTitle_Success", func() {
		task := mocks.GetNewTask()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Params = gin.Params{
			gin.Param{Key: "title", Value: task.Title},
		}

		ctx.Request = httptest.NewRequest("GET", "/task/"+task.Title, nil)

		suite.usecase.On("GetTaskByTitle", task.Title).Return(task, nil).Once()

		suite.Controller.GetTaskByTitle(ctx)

		expected, err := json.Marshal(task)
		suite.Nil(err)

		suite.Equal(202, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
	// test case failure
	suite.Run("TestGetTaskByTitle_Failure", func() {
		task := mocks.GetNewTask()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		
		ctx.Params = gin.Params{
			gin.Param{Key: "title", Value: task.Title},
		}

		ctx.Request = httptest.NewRequest("GET", "/task/"+task.Title, nil)

		suite.usecase.On("GetTaskByTitle", task.Title).Return(nil, errors.New("error")).Once()

		suite.Controller.GetTaskByTitle(ctx)

		expected, err := json.Marshal(gin.H{"error": "error"})
		suite.Nil(err)

		suite.Equal(417, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
}

// test get all tasks
func (suite *TaskControllerSuite) TestGetAllTasks() {
	// test case success
	suite.Run("TestGetAllTasks_Success", func() {
		tasks := mocks.GetMultipleTask()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("GET", "/tasks", nil)

		suite.usecase.On("GetAllTasks").Return(tasks, nil).Once()

		suite.Controller.GetAllTasks(ctx)

		expected, err := json.Marshal(tasks)
		suite.Nil(err)

		suite.Equal(202, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
	// test case failure
	suite.Run("TestGetAllTasks_Failure", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("GET", "/tasks", nil)

		suite.usecase.On("GetAllTasks").Return(nil, errors.New("error")).Once()

		suite.Controller.GetAllTasks(ctx)

		expected, err := json.Marshal(gin.H{"error": "error"})
		suite.Nil(err)

		suite.Equal(417, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
}

// test get user tasks
func (suite *TaskControllerSuite) TestGetUserTasks() {
	// test case success
	suite.Run("TestGetUserTasks_Success", func() {
		tasks := mocks.GetMultipleTask()
		claim := *mocks.GetNewClaim()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Set("claim", claim)

		ctx.Request = httptest.NewRequest("GET", "/usertasks", nil)

		suite.usecase.On("GetUserTasks", claim.UserId).Return(tasks, nil).Once()

		suite.Controller.GetUserTasks(ctx)

		expected, err := json.Marshal(tasks)
		suite.Nil(err)

		suite.Equal(202, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
	// test case failure
	suite.Run("TestGetUserTasks_Failure", func() {
		claim := *mocks.GetNewClaim()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Set("claim", claim)

		ctx.Request = httptest.NewRequest("GET", "/usertasks", nil)

		suite.usecase.On("GetUserTasks", claim.UserId).Return(nil, errors.New("error")).Once()

		suite.Controller.GetUserTasks(ctx)

		expected, err := json.Marshal(gin.H{"error": "error"})
		suite.Nil(err)

		suite.Equal(417, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
}

// test update task
func (suite *TaskControllerSuite) TestUpdateTask() {
	// test case success
	suite.Run("TestUpdateTask_Success", func() {
		task := mocks.GetNewTask()
		claim := *mocks.GetNewClaim()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Set("claim", claim)

		ctx.Params = gin.Params{
			gin.Param{Key: "title", Value: task.Title},
		}

		body, err := json.Marshal(task)
		suite.Nil(err)

		ctx.Request = httptest.NewRequest("PUT", "/task/"+task.Title, strings.NewReader(string(body)))

		suite.usecase.On("UpdateTask", claim.UserRole, claim.UserId, task.Title, task).Return(task, nil).Once()

		suite.Controller.UpdateTask(ctx)

		expected, err := json.Marshal(task)
		suite.Nil(err)

		suite.Equal(201, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})

	// test case invalid json
	suite.Run("TestUpdateTask_InvalidJSON", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Set("claim", *mocks.GetNewClaim())

		ctx.Request = httptest.NewRequest("PUT", "/task/test", strings.NewReader("invalid json"))

		suite.Controller.UpdateTask(ctx)

		expected, err := json.Marshal(gin.H{"error": "Invalid JSON provided"})
		suite.Nil(err)

		suite.Equal(400, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})

	// test case invalid claim
	suite.Run("TestUpdateTask_InvalidClaim", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("PUT", "/task/test", nil)

		suite.Controller.UpdateTask(ctx)

		expected, err := json.Marshal(gin.H{"error": "claim not set"})
		suite.Nil(err)

		suite.Equal(401, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})

	// test case failure
	suite.Run("TestUpdateTask_Failure", func() {
		task := mocks.GetNewTask()
		claim := *mocks.GetNewClaim()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Set("claim", claim)

		ctx.Params = gin.Params{
			gin.Param{Key: "title", Value: task.Title},
		}

		body, err := json.Marshal(task)
		suite.Nil(err)

		ctx.Request = httptest.NewRequest("PUT", "/task/"+task.Title, strings.NewReader(string(body)))

		suite.usecase.On("UpdateTask", claim.UserRole, claim.UserId, task.Title, task).Return(nil, errors.New("error")).Once()

		suite.Controller.UpdateTask(ctx)

		expected, err := json.Marshal(gin.H{"error": "error"})
		suite.Nil(err)

		suite.Equal(401, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})


}

// test delete task
func (suite *TaskControllerSuite) TestDeleteTask() {
	// test case success
	suite.Run("TestDeleteTask_Success", func() {
		task := mocks.GetNewTask()
		claim := *mocks.GetNewClaim()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Set("claim", claim)

		ctx.Params = gin.Params{
			gin.Param{Key: "title", Value: task.Title},
		}

		ctx.Request = httptest.NewRequest("DELETE", "/task/"+task.Title, nil)

		suite.usecase.On("DeleteTask", claim.UserRole, claim.UserId, task.Title).Return(nil).Once()

		suite.Controller.DeleteTask(ctx)

		suite.Equal(201, w.Code)
	})

	// test case invalid claim
	suite.Run("TestDeleteTask_InvalidClaim", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("DELETE", "/task/test", nil)

		suite.Controller.DeleteTask(ctx)

		expected, err := json.Marshal(gin.H{"error": "claim not set"})
		suite.Nil(err)

		suite.Equal(401, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})

	// test case failure
	suite.Run("TestDeleteTask_Failure", func() {
		task := mocks.GetNewTask()
		claim := *mocks.GetNewClaim()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Set("claim", claim)

		ctx.Params = gin.Params{
			gin.Param{Key: "title", Value: task.Title},
		}

		ctx.Request = httptest.NewRequest("DELETE", "/task/"+task.Title, nil)

		suite.usecase.On("DeleteTask", claim.UserRole, claim.UserId, task.Title).Return(errors.New("error")).Once()

		suite.Controller.DeleteTask(ctx)

		expected, err := json.Marshal(gin.H{"error": "error"})
		suite.Nil(err)

		suite.Equal(401, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})

	
}



// tear down the test suite
func (suite *TaskControllerSuite) TearDownSuite() {
	suite.usecase.AssertExpectations(suite.T())
}

// run the test suite
func TestTaskControllerSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}
