package commands

import (
	"github.com/danyukod/cadastro-user-go/internal/application/commands/dto"
	"github.com/danyukod/cadastro-user-go/internal/domain/model"
	"github.com/danyukod/cadastro-user-go/internal/infrastructure/persistence"
)

type FindUserUseCaseInterface interface {
	Execute(dto dto.FindUserDTO) (model.UserDomainInterface, error)
}

func NewFindUserUseCase(persistenceInterface persistence.UserPersistenceInterface) FindUserUseCaseInterface {
	return &findUserUseCase{
		persistenceInterface,
	}
}

type findUserUseCase struct {
	persistence.UserPersistenceInterface
}

func (p *findUserUseCase) Execute(dto dto.FindUserDTO) (model.UserDomainInterface, error) {
	return p.UserPersistenceInterface.FindUserById(dto.Id)
}
