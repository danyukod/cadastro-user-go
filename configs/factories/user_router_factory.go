package factories

import (
	"github.com/danyukod/cadastro-user-go/internal/presentation"
	"github.com/danyukod/cadastro-user-go/internal/presentation/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRouterFactory(database *gorm.DB) error {

	userController := presentation.NewUserHandlerInterface()

	router := gin.Default()
	routes.InitPixKeyRoutes(&router.RouterGroup, userController)

	return router.Run(":8080")
}
