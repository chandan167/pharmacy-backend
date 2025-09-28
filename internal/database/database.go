package database

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/chandan167/pharmacy-backend/internal/model"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
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

func GetTestDbConnect(ctx context.Context) (testcontainers.Container, *gorm.DB, error) {
	// Create container request
	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.0", // MySQL 8
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "root",
			"MYSQL_DATABASE":      "testdb",
			"MYSQL_USER":          "testuser",
			"MYSQL_PASSWORD":      "testpass",
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
	}

	// Start container
	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, nil, err
	}

	// Get host and port
	host, err := mysqlC.Host(ctx)
	if err != nil {
		return nil, nil, err
	}
	port, err := mysqlC.MappedPort(ctx, "3306")
	if err != nil {
		return nil, nil, err
	}

	// DSN
	dsn := fmt.Sprintf("testuser:testpass@tcp(%s:%s)/testdb?charset=utf8mb4&parseTime=True&loc=Local",
		host, port.Port())

	fmt.Println("dns = ", dsn)

	// Connect to DB
	var db *gorm.DB
	for i := 0; i < 10; i++ { // retry few times until MySQL is ready
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Retrying DB connection (%d/10): %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		_ = mysqlC.Terminate(ctx)
		return nil, nil, err
	}

	if db == nil {
		_ = mysqlC.Terminate(ctx)
		return nil, nil, fmt.Errorf("could not connect to database: %w", err)
	}

	// Run migration
	if err := db.AutoMigrate(&model.UserModel{}); err != nil {
		_ = mysqlC.Terminate(ctx)
		return nil, nil, err
	}

	return mysqlC, db, nil
}
