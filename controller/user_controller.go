package controller

import (
	"github.com/chandan167/pharmacy-backend/pkg/helper"
	"github.com/chandan167/pharmacy-backend/service"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	us *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		us: userService,
	}
}

func (uc *UserController) CreateUserHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "user create",
	})
}

func (uc *UserController) GetUsersHandler(ctx *fiber.Ctx) error {
	pageParam, err := helper.GetPaginationSearchParam(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	userResult, err := uc.us.PaginateWithSearchUsers(pageParam)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(fiber.Map{
		"user": userResult,
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
