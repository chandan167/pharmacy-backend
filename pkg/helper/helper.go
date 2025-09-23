package helper

import (
	"errors"
	"strconv"
	"strings"

	"github.com/chandan167/pharmacy-backend/types"
	"github.com/gofiber/fiber/v2"
)

func CalculatePage(totalRecords int64, pageSize int) int {
	return int((totalRecords + int64(pageSize) - 1) / int64(pageSize))
}

func GetPaginationSearchParam(ctx *fiber.Ctx) (*types.PaginationSearchParam, error) {
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		return nil, errors.New("invalid page query")
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size", "50"))
	if err != nil {
		return nil, errors.New("invalid page_size query")
	}
	search := ctx.Query("search", "")
	searchKey := strings.Split(ctx.Query("search_key", ""), ",")
	return &types.PaginationSearchParam{
		Page:      page,
		PageSize:  pageSize,
		Search:    search,
		SearchKey: searchKey,
	}, nil
}
