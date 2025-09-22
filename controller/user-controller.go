package controller

import "github.com/gofiber/fiber/v2"

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) CreateUserHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "user create",
	})
}

func (uc *UserController) GetUsersHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "user list",
	})
}

func (uc *UserController) GetUserHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "user detail",
	})
}

func (uc *UserController) UpdateUserHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "user update",
	})
}

func (uc *UserController) DeleteUserHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "user update",
	})
}
