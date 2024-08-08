package controllers

import (
	"net/http"
	domain "tskmgr/Domain"
	usecases "tskmgr/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type controller struct {
	MyUsecase usecases.Usecase

}

func NewController(coll *mongo.Collection) *controller {
	return &controller{
		MyUsecase: *usecases.NewUserUsecase(coll),
	}
}

func (cont *controller) SignupController(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	err:=cont.MyUsecase.CreateUser(&user)
	if err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"message":"user registered successfully"})

}

func (cont *controller) LoginController(c *gin.Context){
	
}
