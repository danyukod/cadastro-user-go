package routes

import (
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler"
	"github.com/gin-gonic/gin"
)

func InitPixKeyRoutes(
	r *gin.RouterGroup,
	handler handler.UserHandlerInterface,
) {
	r.GET("/users/:id", handler.FindUserById)
}
