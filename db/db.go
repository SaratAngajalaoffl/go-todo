package db

import (
	"context"
	"fmt"

	domain "github.com/SaratAngajalaoffl/go-todo/server/domain/users"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var RedisClient *redis.Client

var Ctx = context.Background()

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=chandrasar dbname=go_todo port=5432 sslmode=disable TimeZone=Asia/Kolkata",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		fmt.Println("Database Error ---  ", err)
	}
	DB.AutoMigrate(&domain.UserModel{})
}