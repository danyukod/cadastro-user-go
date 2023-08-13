package factory

import (
	"github.com/danyukod/cadastro-user-go/internal/application/commands"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler"
	"gorm.io/gorm"
)

func NewUserHandlerFactory(gormDB *gorm.DB) handler.UserHandlerInterface {
	usecase := commands.NewFindUserUseCase()
	return handler.NewUserHandlerInterface(usecase)
}
