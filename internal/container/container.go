package container

import (
	"github.com/chandan167/pharmacy-backend/internal/controller"
	"github.com/chandan167/pharmacy-backend/internal/database"
	"github.com/chandan167/pharmacy-backend/internal/service"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	// Provide Database instance
	container.Provide(database.NewDatabaseConnection)

	//Register Services
	container.Provide(service.NewUserService)

	container.Provide(controller.NewUserController)
	return container
}
