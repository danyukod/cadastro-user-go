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

const (
	SecretKey     = "secretKey"
	JwtExpiration = "jwtExpiration"
)

func NewUserRouterFactory(database *gorm.DB, config configs.Config) error {

	userHandler := factory.NewUserHandlerFactory(database)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	router.Use(middleware.TimeoutMiddleware(), getJwtConfig(config))
	routes.InitPublicUserRoutes(v1, userHandler)
	routes.InitPrivateUserRoutes(v1, userHandler)

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router.Run(":8081")
}

func getJwtConfig(config configs.Config) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set(SecretKey, config.GetJWTSecret())
		c.Set(JwtExpiration, config.GetJWTExpiration())
	}
}
