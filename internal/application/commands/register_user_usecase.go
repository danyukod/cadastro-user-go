package commands

import (
	"github.com/danyukod/cadastro-user-go/internal/application/commands/dto"
	"github.com/danyukod/cadastro-user-go/internal/domain/model"
	"github.com/danyukod/cadastro-user-go/internal/infrastructure/persistence"
)

type RegisterUserUseCaseInterface interface {
	Execute(dto dto.RegisterUserDTO) (model.UserDomainInterface, error)
}

func NewRegisterUserUseCase(persistenceInterface persistence.UserPersistenceInterface) RegisterUserUseCaseInterface {
	return &registerUserUseCase{
		persistenceInterface,
	}
}

type registerUserUseCase struct {
	persistence.UserPersistenceInterface
}

func (p *registerUserUseCase) Execute(dto dto.RegisterUserDTO) (model.UserDomainInterface, error) {
	userDomain, err := model.NewUserDomain(dto.Name, dto.LastName, dto.Email, dto.Document, dto.BirthDate, dto.Password)
	if err != nil {
		return nil, err
	}

	return p.UserPersistenceInterface.CreateUser(userDomain)

}
