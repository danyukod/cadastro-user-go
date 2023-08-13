package commands

import (
	"github.com/danyukod/cadastro-user-go/internal/application/commands/dto"
	"github.com/danyukod/cadastro-user-go/internal/domain/model"
)

type RegisterUserUseCaseInterface interface {
	Execute(dto dto.RegisterUserDTO) (*model.UserDomainInterface, error)
}

func NewRegisterUserUseCase() RegisterUserUseCaseInterface {
	return &registerUserUseCase{}
}

type registerUserUseCase struct {
}

func (p *registerUserUseCase) Execute(dto dto.RegisterUserDTO) (*model.UserDomainInterface, error) {
	panic("implement me")
}
