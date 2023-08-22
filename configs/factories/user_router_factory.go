package factories

import (
	"github.com/danyukod/cadastro-user-go/api/docs"
	"github.com/danyukod/cadastro-user-go/configs"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler/factory"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/routes"
	"github.com/danyukod/cadastro-user-go/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func NewUserRouterFactory(database *gorm.DB, config configs.Config) error {

	userHandler := factory.NewUserHandlerFactory(database)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	v1.Use(middleware.TimeoutMiddleware())
	v1.Use(middleware.SetJwtConfig(config.GetJWTSecret(), config.GetJWTExpiration()))
	routes.InitPublicUserRoutes(v1, userHandler)
	routes.InitPrivateUserRoutes(v1, userHandler)

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router.Run(":8081")
}
