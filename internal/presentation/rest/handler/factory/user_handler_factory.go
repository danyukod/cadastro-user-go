package factory

import (
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler"
	"gorm.io/gorm"
)

func NewUserHandlerFactory(gormDB *gorm.DB) handler.UserHandlerInterface {
	return handler.NewUserHandlerInterface()
}
