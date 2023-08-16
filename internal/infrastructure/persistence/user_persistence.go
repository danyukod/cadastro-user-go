package persistence

import (
	"github.com/danyukod/cadastro-user-go/internal/domain/model"
	entity2 "github.com/danyukod/cadastro-user-go/internal/infrastructure/persistence/entity"
	"gorm.io/gorm"
)

type UserPersistenceInterface interface {
	FindUserById(id string) (model.UserDomainInterface, error)
	CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, error)
	FindByEmail(email string) (model.UserDomainInterface, error)
}

type userPersistence struct {
	*gorm.DB
}

func NewUserPersistence(gormDB *gorm.DB) UserPersistenceInterface {
	return &userPersistence{gormDB}
}

func (u userPersistence) CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, error) {
	entity := entity2.DomainToEntity(user)

	if err := u.DB.Create(&entity).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u userPersistence) FindUserById(id string) (model.UserDomainInterface, error) {
	var entity entity2.UserEntity
	if err := u.DB.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return entity2.EntityToDomain(entity)
}

func (u userPersistence) FindByEmail(email string) (model.UserDomainInterface, error) {
	var entity entity2.UserEntity
	if err := u.DB.First(&entity, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return entity2.EntityToDomain(entity)
}
