package commands

import (
	"github.com/danyukod/cadastro-user-go/internal/application/commands/dto"
	"github.com/danyukod/cadastro-user-go/internal/infrastructure/persistence"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"time"
)

type GenerateTokenUseCaseInterface interface {
	Execute(dto dto.GenerateTokenDTO) (string, error)
}

type generateTokenUseCase struct {
	persistence.UserPersistenceInterface
}

func NewGenerateTokenUseCaseInterface(
	userPersistence persistence.UserPersistenceInterface,
) GenerateTokenUseCaseInterface {
	return &generateTokenUseCase{
		userPersistence,
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

	claims := jwt.StandardClaims{
		Subject:   u.GetID().String(),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(dto.JwtExpiresIn)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	accessToken, err := token.SignedString([]byte(dto.JwtSecret))

	return accessToken, nil
}
