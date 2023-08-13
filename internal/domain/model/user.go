package model

import (
	"github.com/danyukod/cadastro-user-go/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserDomainInterface interface {
	GetID() entity.ID
	SetID(id entity.ID)
	GetName() string
	GetLastName() string
	GetEmail() string
	GetDocument() string
	GetBirthDate() string
	GetPassword() string
	ValidatePassword(string) bool
	Validate() error
}

func NewUserDomain(name, lastName, email, document, birthDate, password string) (UserDomainInterface, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userDomain := userDomain{
		id:        entity.NewID(),
		name:      name,
		lastName:  lastName,
		email:     email,
		document:  document,
		birthDate: birthDate,
		password:  string(hash),
	}
	if err := userDomain.Validate(); err != nil {
		return nil, err
	}
	return &userDomain, nil
}

type userDomain struct {
	id        entity.ID
	name      string
	lastName  string
	email     string
	document  string
	birthDate string
	password  string
}

func (u userDomain) GetPassword() string {
	return u.password
}

func (u userDomain) GetID() entity.ID {
	return u.id
}

func (u userDomain) SetID(id entity.ID) {
	u.id = id
}

func (u userDomain) GetName() string {
	return u.name
}

func (u userDomain) GetLastName() string {
	return u.lastName
}

func (u userDomain) GetEmail() string {
	return u.email
}

func (u userDomain) GetBirthDate() string {
	return u.birthDate
}

func (u userDomain) GetDocument() string {
	return u.document
}

func (u userDomain) Validate() error {
	return nil
}

func (u userDomain) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}
