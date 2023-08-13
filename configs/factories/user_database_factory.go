package factories

import (
	"github.com/danyukod/cadastro-user-go/configs"
	"github.com/danyukod/cadastro-user-go/configs/logger"
	"gorm.io/gorm"
)

func NewUserDatabaseFactory(conf *configs.Config) (*gorm.DB, error) {
	if err := configs.MigrateDatabase(*conf); err != nil {
		logger.Error("Database migration failed: ", err)
		return nil, err
	}

	database, err := NewGormDatabaseFactory(*conf)
	if err != nil {
		logger.Error("Failed to create Gorm Database: ", err)
		return nil, err
	}

	return database, nil
}
