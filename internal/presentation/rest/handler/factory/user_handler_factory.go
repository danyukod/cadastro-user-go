package factory

import (
	"github.com/danyukod/cadastro-user-go/internal/application/commands"
	"github.com/danyukod/cadastro-user-go/internal/infrastructure/persistence"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler"
	"gorm.io/gorm"
)

func NewUserHandlerFactory(gormDB *gorm.DB) handler.UserHandlerInterface {
	userPersistence := persistence.NewUserPersistence(gormDB)
	findUseCase := commands.NewFindUserUseCase(userPersistence)
	registerUseCase := commands.NewRegisterUserUseCase(userPersistence)
	return handler.NewUserHandlerInterface(findUseCase, registerUseCase)
}
