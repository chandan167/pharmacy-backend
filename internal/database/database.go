package database

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/chandan167/pharmacy-backend/internal/model"
	"gopkg.in/natefinch/lumberjack.v2"
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

func getDbLogger() logger.Interface {
	var logWriter []io.Writer
	logWriter = append(logWriter, &lumberjack.Logger{
		Filename:   "logs/sql/sql-log.log",
		MaxSize:    20, // MB
		MaxBackups: 10,
		MaxAge:     30,   // days
		Compress:   true, // gzip
	})

	if os.Getenv("GO_ENV") == "development" {
		logWriter = append(logWriter, os.Stdout)
	}
	var newLogger = logger.New(
		log.New(io.MultiWriter(logWriter...), "\r\n[QUERY]\t", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Log queries slower than this
			LogLevel:      logger.Info, // Log level: Silent, Error, Warn, Info
			Colorful:      false,       // Color output
		},
	)
	return newLogger
}

func DbConnect(config *DbConfig) *gorm.DB {
	dbLogger := getDbLogger()
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.DbName,
	)
	con, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		time.Sleep(time.Second * 5)
		return DbConnect(config)
	}

	log.Println("Database connected")
	con.AutoMigrate(&model.UserModel{})
	return con
}
