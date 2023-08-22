package request

import (
	"github.com/danyukod/cadastro-user-go/internal/application/commands/dto"
)

type RegisterUserRequest struct {
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Document  string `json:"document"`
	BirthDate string `json:"birthdate"`
	Password  string `json:"password"`
}

type FindUserRequest struct {
	Id string `uri:"id" binding:"required"`
}

type GetJWTRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r GetJWTRequest) ToDTO(secret string, expire int) dto.GenerateTokenDTO {
	return dto.GenerateTokenDTO{
		Email:        r.Email,
		Password:     r.Password,
		JwtSecret:    secret,
		JwtExpiresIn: expire,
	}
}

func (p *RegisterUserRequest) ToDTO() dto.RegisterUserDTO {
	return dto.RegisterUserDTO{
		Name:      p.Name,
		LastName:  p.LastName,
		Email:     p.Email,
		Document:  p.Document,
		BirthDate: p.BirthDate,
		Password:  p.Password,
	}
}

func (p *FindUserRequest) ToDTO() dto.FindUserDTO {
	return dto.FindUserDTO{
		Id: p.Id,
	}
}
