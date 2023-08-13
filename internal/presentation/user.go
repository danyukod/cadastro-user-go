package presentation

import "github.com/gin-gonic/gin"

func NewUserHandlerInterface() UserHandlerInterface {
	return &userHandler{}
}

type UserHandlerInterface interface {
	FindUserById(c *gin.Context)
}

type userHandler struct {
}

func (p *userHandler) FindUserById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
