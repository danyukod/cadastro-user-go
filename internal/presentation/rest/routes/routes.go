package routes

import (
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler"
	"github.com/danyukod/cadastro-user-go/pkg/midleware"
	"github.com/gin-gonic/gin"
)

func InitPrivateUserRoutes(
	r *gin.RouterGroup,
	handler handler.UserHandlerInterface,
	secretKey string,
) {
	r.Use(midleware.TokenAuthMiddleware(secretKey))
	r.GET("/:id", handler.FindUserById)
}

func InitPublicUserRoutes(
	r *gin.RouterGroup,
	handler handler.UserHandlerInterface,
) {
	r.POST("auth", handler.GetJWT)
	r.POST("/", handler.RegisterUser)
}
