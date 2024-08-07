package controllers

import (
	"net/http"
	"tskmgr/data"
	"tskmgr/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userdata *data.UserData
}

func NewUserController(coll *data.UserData) *UserController {
	return &UserController{userdata: coll}
}

func (uc *UserController) SignupController(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	err := uc.userdata.SignupService(&user)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})

}

func (uc *UserController) LoginController(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})

	}

	UserToken, err := uc.userdata.LoginService(&user)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": UserToken})

}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.userdata.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)

}

func (uc *UserController) ChangeRole(c *gin.Context){
	username := c.Param("username")	
	err := uc.userdata.ChangeRoleUser(username)
	
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message": "role toggled successfuly"})
	
}
