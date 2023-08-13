package routes

import (
	"github.com/danyukod/cadastro-user-go/internal/presentation"
	"github.com/gin-gonic/gin"
)

func InitPixKeyRoutes(
	r *gin.RouterGroup,
	handler presentation.UserHandlerInterface,
) {
	r.GET("/users/:id", handler.FindUserById)
}
