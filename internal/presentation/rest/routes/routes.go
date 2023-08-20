package routes

import (
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler"
	"github.com/danyukod/cadastro-user-go/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func InitPrivateUserRoutes(
	r *gin.RouterGroup,
	handler handler.UserHandlerInterface,
) {
	usersGroup := r.Group("/users")
	{
		usersGroup.Use(middleware.TokenAuthMiddleware())
		usersGroup.GET("/:id", handler.FindUserById)
	}
}

func InitPublicUserRoutes(
	r *gin.RouterGroup,
	handler handler.UserHandlerInterface,
) {
	usersGroup := r.Group("/users")
	{
		usersGroup.POST("auth", handler.GetJWT)
		usersGroup.POST("/", handler.RegisterUser)
	}
}
