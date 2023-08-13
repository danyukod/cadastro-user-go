package request

import "github.com/danyukod/cadastro-user-go/internal/application/commands/dto"

type RegisterUserRequest struct {
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Document  string `json:"document"`
	BirthDate string `json:"birth_date"`
	Password  string `json:"password"`
}

type FindUserRequest struct {
	Id string `uri:"id" binding:"required"`
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
