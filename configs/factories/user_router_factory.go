package factories

import (
	"github.com/danyukod/cadastro-user-go/internal/app/user/src/ui/adapter/rest/controller"
	"github.com/danyukod/cadastro-user-go/internal/app/user/src/ui/adapter/rest/controller/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRouterFactory(database *gorm.DB) error {

	userController := controller.NewUserControllerInterface()

	router := gin.Default()
	routes.InitPixKeyRoutes(&router.RouterGroup, userController)

	return router.Run(":8080")
}
