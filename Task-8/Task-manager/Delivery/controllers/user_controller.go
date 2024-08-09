package controllers

import (
	"net/http"
	domain "tskmgr/Domain"
	usecases "tskmgr/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type usercontroller struct {
	MyuserUsecase usecases.UserUsecase

}

func NewUserController(coll *mongo.Collection) *usercontroller {
	return &usercontroller{
		MyuserUsecase: *usecases.NewUserUsecase(coll),
	}
}

func (cont *usercontroller) SignupController(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	err:=cont.MyuserUsecase.CreateUser(&user)
	if err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"message":"user registered successfully"})

}

func (cont *usercontroller) LoginController(c *gin.Context){
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	token,err := cont.MyuserUsecase.LogUser(&user)
	if err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	c.JSON(200,gin.H{"message":"user logged in successfully", "token": token})

}
