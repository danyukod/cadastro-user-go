package factory

import (
	"github.com/danyukod/cadastro-user-go/internal/app/user/src/ui/adapter/rest/controller"
	"gorm.io/gorm"
)

func NewUserControllerFactory(gormDB *gorm.DB) controller.UserControllerInterface {
	return controller.NewUserControllerInterface()
}
