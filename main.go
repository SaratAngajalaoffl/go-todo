package main

import (
	"github.com/SaratAngajalaoffl/go-todo/db"
	Routes "github.com/SaratAngajalaoffl/go-todo/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	router := gin.Default()
	Routes.UserRoutes(&router.RouterGroup, db.DB)
	router.Run()
}