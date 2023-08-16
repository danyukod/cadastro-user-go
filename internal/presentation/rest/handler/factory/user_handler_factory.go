package factory

import (
	"github.com/danyukod/cadastro-user-go/internal/application/commands"
	"github.com/danyukod/cadastro-user-go/internal/infrastructure/persistence"
	"github.com/danyukod/cadastro-user-go/internal/presentation/rest/handler"
	"github.com/go-chi/jwtauth"
	"gorm.io/gorm"
)

func NewUserHandlerFactory(gormDB *gorm.DB, jwtTokenAuth *jwtauth.JWTAuth, jwtExpiration int) handler.UserHandlerInterface {
	userPersistence := persistence.NewUserPersistence(gormDB)
	findUseCase := commands.NewFindUserUseCase(userPersistence)
	registerUseCase := commands.NewRegisterUserUseCase(userPersistence)
	generateTokenUseCase := commands.NewGenerateTokenUseCaseInterface(userPersistence, jwtTokenAuth, jwtExpiration)
	return handler.NewUserHandlerInterface(findUseCase, registerUseCase, generateTokenUseCase)
}
