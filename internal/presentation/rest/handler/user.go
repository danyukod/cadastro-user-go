package handler

import (
	"github.com/danyukod/cadastro-user-go/internal/application/commands"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler/model/request"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SecretKey     = "secretKey"
	JwtExpiration = "jwtExpiration"
)

func NewUserHandlerInterface(
	findUserUseCase commands.FindUserUseCaseInterface,
	registerUseCase commands.RegisterUserUseCaseInterface,
	generateTokenUseCase commands.GenerateTokenUseCaseInterface) UserHandlerInterface {
	return &handler{
		findUserUseCase:      findUserUseCase,
		registerUseCase:      registerUseCase,
		generateTokenUseCase: generateTokenUseCase,
	}
}

type UserHandlerInterface interface {
	RegisterUser(c *gin.Context)
	FindUserById(c *gin.Context)
	GetJWT(c *gin.Context)
}

type handler struct {
	findUserUseCase      commands.FindUserUseCaseInterface
	registerUseCase      commands.RegisterUserUseCaseInterface
	generateTokenUseCase commands.GenerateTokenUseCaseInterface
}

// FindUserById godoc
// @Summary Find By UserId
// @Schemes
// @Description Find By UserId
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.UserResponse
// @Failure 500 {object} ErrorsResponse
// @Failure 400 {object} ErrorsResponse
// @Router /users/{id} [get]
// @Security ApiKeyAuth
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

	c.JSON(http.StatusOK, response.UserDomainToFindWebResponse(userDomain))

}

// RegisterUser godoc
// @Summary Register User
// @Schemes
// @Description Register User
// @Tags users
// @Accept json
// @Produce json
// @Param request body request.RegisterUserRequest true "register-user request"
// @Success 201 {object} response.UserResponse
// @Failure 500 {object} ErrorsResponse
// @Failure 400 {object} ErrorsResponse
// @Router /users [post]
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

	c.JSON(http.StatusCreated, response.UserDomainToRegisterWebResponse(userDomain))
}

// GetJWT godoc
// @Summary Get JWT
// @Schemes
// @Description Get JWT
// @Tags jwt
// @Accept json
// @Produce json
// @Param request body request.GetJWTRequest true "user-credentials request"
// @Success 200 {object} response.JWTResponse
// @Failure 500 {object} ErrorsResponse
// @Failure 400 {object} ErrorsResponse
// @Failure 404 {object} ErrorsResponse
// @Router /users/auth [post]
func (p *handler) GetJWT(c *gin.Context) {
	var userRequest request.GetJWTRequest
	secretKey := c.GetString(SecretKey)
	jwtExpiresIn := c.GetInt(JwtExpiration)

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	token, err := p.generateTokenUseCase.Execute(userRequest.ToDTO(secretKey, jwtExpiresIn))
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, response.JWTResponse{AccessToken: token})
}
