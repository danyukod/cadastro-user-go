package factory

import (
	"github.com/danyukod/cadastro-user-go/internal/presentation"
	"gorm.io/gorm"
)

func NewUserHandlerFactory(gormDB *gorm.DB) presentation.UserHandlerInterface {
	return presentation.NewUserHandlerInterface()
}
