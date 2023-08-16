package response

import "github.com/danyukod/cadastro-user-go/internal/domain/model"

type UserResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Document  string `json:"document"`
	BirthDate string `json:"birth_date"`
	Password  string `json:"password"`
}

type JWTResponse struct {
	AccessToken string `json:"access_token"`
}

func UserDomainToFindWebResponse(userDomain model.UserDomainInterface) *UserResponse {
	return &UserResponse{
		Id:        userDomain.GetID().String(),
		Name:      userDomain.GetName(),
		LastName:  userDomain.GetLastName(),
		Email:     userDomain.GetEmail(),
		Document:  userDomain.GetDocument(),
		BirthDate: userDomain.GetBirthDate().String(),
		Password:  userDomain.GetPassword(),
	}
}

func UserDomainToRegisterWebResponse(userDomain model.UserDomainInterface) *UserResponse {
	return &UserResponse{
		Id:        userDomain.GetID().String(),
		Name:      userDomain.GetName(),
		LastName:  userDomain.GetLastName(),
		Email:     userDomain.GetEmail(),
		Document:  userDomain.GetDocument(),
		BirthDate: userDomain.GetBirthDate().String(),
		Password:  userDomain.GetPassword(),
	}
}
