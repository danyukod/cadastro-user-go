package request

type CreateUserRequest struct {
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Document  string `json:"document"`
	BirthDate string `json:"birth_date"`
	Password  string `json:"password"`
}
