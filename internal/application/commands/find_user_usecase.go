package commands

import (
	"github.com/danyukod/cadastro-user-go/internal/application/commands/dto"
	"github.com/danyukod/cadastro-user-go/internal/domain/model"
)

type FindUserUseCaseInterface interface {
	Execute(dto dto.FindUserDTO) (*model.UserDomainInterface, error)
}

func NewFindUserUseCase() FindUserUseCaseInterface {
	return &findUserUseCase{}
}

type findUserUseCase struct {
}

func (p *findUserUseCase) Execute(dto dto.FindUserDTO) (*model.UserDomainInterface, error) {
	panic("implement me")
}
