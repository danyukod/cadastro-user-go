package controller

import "github.com/gin-gonic/gin"

func NewUserControllerInterface() UserControllerInterface {
	return &userController{}
}

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
}

type userController struct {
}
