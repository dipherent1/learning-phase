package controllers

import (
	"errors"
	"net/http"
	domain "tskmgr/Domain"
	usecases "tskmgr/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskcontroller struct {
	MyTaskUsecase usecases.TaskUsecase
}

func NewTaskController(coll *mongo.Collection) *taskcontroller {
	return &taskcontroller{
		MyTaskUsecase: *usecases.NewTaskUsecase(coll),
	}
}

func getclaim(c *gin.Context) (*domain.Claims, error) {
	claim, exists := c.Get("claim")
	if !exists {
		return nil, errors.New("claim not set")
	}

	userClaims, ok := claim.(domain.Claims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return &userClaims, nil
}

func (cont *taskcontroller) CreateTask(c *gin.Context) {
	claim, err := getclaim(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	newTask := &domain.Task{}
	if err := c.ShouldBindJSON(newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}

	userid := claim.UserId
	newTask.UserId = userid

	newTask, err = cont.MyTaskUsecase.CreateTask(newTask)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTask)
}

func (cont *taskcontroller) GetTaskByTitle(c *gin.Context) {

	title := c.Param("title")
	task,err := cont.MyTaskUsecase.GetTaskByTitle(title)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusAccepted, task)

}
