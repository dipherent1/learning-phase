package controllers

import (
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

	uc.userdata.SignupService(&user)
	
	c.JSON(200, gin.H{"message": "User registered successfully"})

}

func (uc *UserController) LoginController(c *gin.Context) {
	var user models.User
	
	if err := c.ShouldBindJSON(&user); err != nil {
	  c.JSON(400, gin.H{"error": "Invalid request payload"})
	  return
	}

	jwtToken := uc.userdata.LoginService(&user)
	
	c.JSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})

}
