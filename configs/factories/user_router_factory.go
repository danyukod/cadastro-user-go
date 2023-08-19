package factories

import (
	"github.com/danyukod/cadastro-user-go/configs"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler/factory"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRouterFactory(database *gorm.DB, config configs.Config) error {

	userHandler := factory.NewUserHandlerFactory(database, config.GetTokenAuth(), config.GetJWTExpiration())

	router := gin.Default()
	routes.InitPublicUserRoutes(&router.RouterGroup, userHandler)
	routes.InitPrivateUserRoutes(&router.RouterGroup, userHandler, config.GetJWTSecret())

	return router.Run(":8081")
}
