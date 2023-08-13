package main

import (
	"github.com/danyukod/cadastro-user-go/configs"
	"github.com/danyukod/cadastro-user-go/configs/factories"
	"github.com/danyukod/cadastro-user-go/configs/logger"
)

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

	if err := factories.NewUserRouterFactory(database); err != nil {
		logger.Error("Failed to start PixKey Router: ", err)
		return err
	}

	return nil
}
