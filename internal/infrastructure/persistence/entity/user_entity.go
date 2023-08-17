package entity

import (
	"github.com/danyukod/cadastro-user-go/internal/domain/model"
	"github.com/danyukod/cadastro-user-go/pkg/entity"
	"time"
)

type UserEntity struct {
	ID         entity.ID `gorm:"primaryKey"`
	Name       string    `gorm:"not null"`
	LastName   string    `gorm:"not null"`
	BirthDate  time.Time `gorm:"type:timestamp;not null"`
	Document   string    `gorm:"not null"`
	Email      string    `gorm:"not null"`
	Password   string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"type:timestamp;not null"`
	ModifiedAt time.Time `gorm:"type:timestamp;not null"`
}

func DomainToEntity(domain model.UserDomainInterface) *UserEntity {
	return &UserEntity{
		ID:         domain.GetID(),
		Name:       domain.GetName(),
		LastName:   domain.GetLastName(),
		BirthDate:  domain.GetBirthDate(),
		Document:   domain.GetDocument(),
		Email:      domain.GetEmail(),
		Password:   domain.GetPassword(),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}
}

func EntityToDomain(entity UserEntity) (model.UserDomainInterface, error) {
	return model.NewUserDomainFromEntity(
		entity.ID,
		entity.BirthDate,
		entity.Name,
		entity.LastName,
		entity.Email,
		entity.Document,
		entity.Password,
	)
}

func (UserEntity) TableName() string {
	return "cadastro_user"
}
