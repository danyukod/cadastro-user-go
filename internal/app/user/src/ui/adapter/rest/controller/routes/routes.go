package routes

import (
	"github.com/danyukod/cadastro-user-go/internal/app/user/src/ui/adapter/rest/controller"
	"github.com/gin-gonic/gin"
)

func InitPixKeyRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
	r.GET("/users/:id", userController.FindUserById)
}
