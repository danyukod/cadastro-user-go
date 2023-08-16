package commands

import (
	"github.com/danyukod/cadastro-user-go/internal/application/commands/dto"
	"github.com/danyukod/cadastro-user-go/internal/infrastructure/persistence"
	"github.com/go-chi/jwtauth"
	"github.com/pkg/errors"
	"time"
)

type GenerateTokenUseCaseInterface interface {
	Execute(dto dto.GenerateTokenDTO) (string, error)
}

type generateTokenUseCase struct {
	persistence.UserPersistenceInterface
	*jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewGenerateTokenUseCaseInterface(
	userPersistence persistence.UserPersistenceInterface,
	jwtAuth *jwtauth.JWTAuth,
	jwtExpiresIn int,
) GenerateTokenUseCaseInterface {
	return &generateTokenUseCase{
		userPersistence,
		jwtAuth,
		jwtExpiresIn,
	}
}

func (g generateTokenUseCase) Execute(dto dto.GenerateTokenDTO) (string, error) {

	u, err := g.UserPersistenceInterface.FindByEmail(dto.Email)
	if err != nil {
		return "", err
	}

	if !u.ValidatePassword(dto.Password) {
		return "", errors.New("invalid password")
	}

	_, accessToken, err := g.JWTAuth.Encode(map[string]interface{}{
		"sub": u.GetID().String(),
		"exp": time.Now().Add(time.Second * time.Duration(g.JwtExpiresIn)).Unix(),
	})

	return accessToken, nil
}
