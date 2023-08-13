package factories

import (
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRouterFactory(database *gorm.DB) error {

	userController := handler.NewUserHandlerInterface()

	router := gin.Default()
	routes.InitPixKeyRoutes(&router.RouterGroup, userController)

	return router.Run(":8080")
}
