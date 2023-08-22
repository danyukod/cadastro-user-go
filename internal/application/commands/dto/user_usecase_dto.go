package dto

type FindUserDTO struct {
	Id string
}

type RegisterUserDTO struct {
	Name      string
	LastName  string
	Email     string
	Document  string
	BirthDate string
	Password  string
}

type GenerateTokenDTO struct {
	Email        string
	Password     string
	JwtSecret    string
	JwtExpiresIn int
}
