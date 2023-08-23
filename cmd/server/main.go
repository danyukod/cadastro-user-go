package main

import (
	"github.com/danyukod/cadastro-user-go/configs"
	"github.com/danyukod/cadastro-user-go/configs/factories"
	"github.com/danyukod/cadastro-user-go/configs/logger"
)

// @title Cadastro de Usuário API
// @version v1
// @description API para cadastro de usuário
// @termsOfService http://swagger.io/terms/

// @contact.name Danilo Kodavara
// @contact.url  https://www.linkedin.com/in/danilo-kodavara/
// @contact.email danilo.kodavara@gmail.com

// @license.name Danilo  Kodavara License
// @license.url https://www.linkedin.com/in/danilo-kodavara/

// @host localhost:8081
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := startPixKeyAPI()
	if err != nil {
		return
	}
}

func startPixKeyAPI() error {
	logger.Info("About to start User API...")

	conf, err := configs.LoadConfig("cmd/server")
	if err != nil {
		logger.Error("Failed to load config: ", err)
		return err
	}

	database, err := factories.NewUserDatabaseFactory(conf)
	if err != nil {
		logger.Error("Failed to create PixKey Database: ", err)
		return err
	}

	if err := factories.NewUserRouterFactory(database, *conf); err != nil {
		logger.Error("Failed to start PixKey Router: ", err)
		return err
	}

	return nil
}
