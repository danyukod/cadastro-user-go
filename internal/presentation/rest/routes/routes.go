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
	r.Use(middleware.TokenAuthMiddleware())
	r.GET("/:id", handler.FindUserById)
}

func InitPublicUserRoutes(
	r *gin.RouterGroup,
	handler handler.UserHandlerInterface,
) {
	r.POST("auth", handler.GetJWT)
	r.POST("/", handler.RegisterUser)
}
