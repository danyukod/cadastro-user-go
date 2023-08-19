package factories

import (
	"github.com/danyukod/cadastro-user-go/configs"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler/factory"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/routes"
	"github.com/danyukod/cadastro-user-go/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	SecretKey     = "secretKey"
	JwtExpiration = "jwtExpiration"
)

func NewUserRouterFactory(database *gorm.DB, config configs.Config) error {

	userHandler := factory.NewUserHandlerFactory(database)

	router := gin.Default()
	router.Use(middleware.TimeoutMiddleware())
	rGroup := router.Group("/users")
	rGroup.Use(func(c *gin.Context) {
		c.Set(SecretKey, config.GetJWTSecret())
		c.Set(JwtExpiration, config.GetJWTExpiration())
	})
	routes.InitPublicUserRoutes(rGroup, userHandler)
	routes.InitPrivateUserRoutes(rGroup, userHandler)

	return router.Run(":8081")
}
