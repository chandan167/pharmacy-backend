package container

import (
	"github.com/chandan167/pharmacy-backend/controller"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(controller.NewUserController)
	return container
}
