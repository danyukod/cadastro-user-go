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
	r.POST("/users/auth", handler.GetJWT)
	r.Use(midleware.TokenAuthMiddleware(secretKey))
	r.POST("/users", handler.RegisterUser)
	r.GET("/users/:id", handler.FindUserById)

}
