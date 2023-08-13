package handler

import (
	"github.com/danyukod/cadastro-user-go/internal/application/commands"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler/model/request"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUserHandlerInterface(findUserUseCase commands.FindUserUseCaseInterface, registerUseCase commands.RegisterUserUseCaseInterface) UserHandlerInterface {
	return &handler{
		findUserUseCase: findUserUseCase,
		registerUseCase: registerUseCase,
	}
}

type UserHandlerInterface interface {
	RegisterUser(c *gin.Context)
	FindUserById(c *gin.Context)
}

type handler struct {
	findUserUseCase commands.FindUserUseCaseInterface
	registerUseCase commands.RegisterUserUseCaseInterface
}

func (p *handler) FindUserById(c *gin.Context) {
	var userRequest request.FindUserRequest

	err := c.ShouldBindUri(&userRequest)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	userDomain, err := p.findUserUseCase.Execute(userRequest.ToDTO())
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, response.UserDomainToFindWebResponse(*userDomain))

}

func (p *handler) RegisterUser(c *gin.Context) {
	var userRequest request.RegisterUserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	userDomain, err := p.registerUseCase.Execute(userRequest.ToDTO())
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.UserDomainToRegisterWebResponse(*userDomain))
}
