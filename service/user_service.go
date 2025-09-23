package service

import (
	"fmt"
	"strings"

	"github.com/chandan167/pharmacy-backend/model"
	"github.com/chandan167/pharmacy-backend/pkg/helper"
	"github.com/chandan167/pharmacy-backend/types"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (us *UserService) PaginateWithSearchUsers(param *types.PaginationSearchParam) (*types.PaginatedResult[model.UserModel], error) {
	var users []model.UserModel
	var total int64
	query := us.db.Model(&model.UserModel{})
	offset := (param.Page - 1) * param.PageSize

	if param.Search != "" && len(param.SearchKey) > 0 {
		like := "%" + param.Search + "%"
		var conditions []string
		var values []interface{}

		for _, f := range param.SearchKey {
			conditions = append(conditions, f+" Like ?")
			values = append(values, like)
		}
		fmt.Println(strings.Join(conditions, " OR "))
		query = query.Where(strings.Join(conditions, " OR "), values...)
	}
	query.Count(&total)
	result := query.Limit(param.PageSize).Offset(offset).Find(&users)
	totalPage := helper.CalculatePage(total, param.PageSize)
	return &types.PaginatedResult[model.UserModel]{
		Data:      users,
		Total:     total,
		Page:      param.Page,
		PageSize:  param.PageSize,
		TotalPage: totalPage,
	}, result.Error
}

func (us *UserService) DeleteUserById(id int) error {
	result := us.db.Delete(&model.UserModel{}, id).Error
	return result
}

func (us *UserService) GetUserById(id int) (*model.UserModel, error) {
	var user model.UserModel
	err := us.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
