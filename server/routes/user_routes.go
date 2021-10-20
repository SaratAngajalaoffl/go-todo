package routes

import (
	domain "github.com/SaratAngajalaoffl/go-todo/server/domain/users"
	handlers "github.com/SaratAngajalaoffl/go-todo/server/handlers"
	service "github.com/SaratAngajalaoffl/go-todo/server/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(apiRouter *gin.RouterGroup, db *gorm.DB) {
	uh := handlers.Userhandlers{Service: service.NewCustomerService(domain.NewUserRepositoryDb(db))}
	route := apiRouter.Group("/users")
	{
		route.POST("/", uh.CreateUser)
	}
}