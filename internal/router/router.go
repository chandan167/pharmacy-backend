package router

import (
	"github.com/chandan167/pharmacy-backend/internal/container"
	"github.com/chandan167/pharmacy-backend/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoute(app *fiber.App) {
	apiRouter := app.Group("/api/v1")
	ioc := container.BuildContainer()

	ioc.Invoke(func(userController *controller.UserController) {
		userRouter := apiRouter.Group("/user")
		userRouter.Get("/", userController.GetUsersHandler)
		userRouter.Post("/", userController.CreateUserHandler)
		userRouter.Get("/:id", userController.GetUserHandler)
		userRouter.Put("/:id", userController.UpdateUserHandler)
		userRouter.Delete("/:id", userController.DeleteUserHandler)
	})

}
