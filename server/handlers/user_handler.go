package users

import (
	domain "github.com/SaratAngajalaoffl/go-todo/server/domain/users"
	service "github.com/SaratAngajalaoffl/go-todo/server/services"
	"github.com/gin-gonic/gin"
)

type Userhandlers struct {
	Service service.UserService
}

func (uh *Userhandlers) CreateUser(c *gin.Context) {
	u := domain.UserModel{}
	if err := c.BindJSON(&u); err != nil {
		c.JSON(406, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, err := uh.Service.CreateUser(u)
	if err != nil {
		c.JSON(406, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": user,
	})
}