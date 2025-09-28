package controller

import (
	"errors"
	"strconv"

	"github.com/chandan167/pharmacy-backend/internal/service"
	"github.com/chandan167/pharmacy-backend/pkg/helper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
		return helper.BadRequestError(err.Error())
	}
	userResult, err := uc.us.PaginateWithSearchUsers(pageParam)
	if err != nil {
		return helper.NewAppError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(fiber.Map{
		"user": userResult,
	})
}

func (uc *UserController) GetUserHandler(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequestError("invalid param id")
	}
	users, err := uc.us.GetUserById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.NotFoundError("user not found")
		}
		return helper.NewAppError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(fiber.Map{
		"message": "user detail",
		"user":    users,
	})
}

func (uc *UserController) UpdateUserHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "user update",
	})
}

func (uc *UserController) DeleteUserHandler(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequestError("invalid param id")
	}
	uc.us.DeleteUserById(id)
	return ctx.JSON(fiber.Map{
		"message": "user deleted successful",
	})
}
