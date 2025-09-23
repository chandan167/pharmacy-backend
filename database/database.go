package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/chandan167/pharmacy-backend/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConfig struct {
	DbName   string
	Host     string
	Username string
	Port     int16
	Password string
}

func NewDatabaseConnection() *gorm.DB {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	config := &DbConfig{
		DbName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     int16(port),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	return DbConnect(config)
}

func DbConnect(config *DbConfig) *gorm.DB {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.DbName,
	)
	con, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		time.Sleep(time.Second * 5)
		return DbConnect(config)
	}

	log.Println("Database connected")
	con.AutoMigrate(&model.UserModel{})
	return con
}
