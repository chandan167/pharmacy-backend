package service_test

import (
	"context"
	"os"
	"testing"

	"github.com/chandan167/pharmacy-backend/internal/database"
	"github.com/chandan167/pharmacy-backend/internal/model"
	"github.com/chandan167/pharmacy-backend/internal/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	userService *service.UserService
	goDb        *gorm.DB
)

func TestMain(m *testing.M) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mysqlC, db, _ := database.GetTestDbConnect(ctx)
	defer func() {
		if mysqlC != nil {
			_ = mysqlC.Terminate(ctx)
		}
	}()
	userService = service.NewUserService(db)
	goDb = db

	code := m.Run()

	os.Exit(code)
}

func TestGetUserById(t *testing.T) {
	user, err := userService.GetUserById(1)
	assert.EqualError(t, err, "record not found", "expected user not found")
	assert.Nil(t, user, "expected user to be NULL")
	newUser := model.UserModel{
		Name:     "chandan",
		Email:    "chandan@gmail.com",
		Password: "password",
		IsActive: true,
	}
	goDb.Create(&newUser)

	user, err = userService.GetUserById(int(newUser.ID))
	assert.Nil(t, err, "expected error to be NULL")
	assert.Equal(t, newUser.Email, user.Email)
	assert.Equal(t, newUser.ID, user.ID)

}
