package controller

import (
	"github.com/cnh1en/lets_go/internal/services"
	"github.com/gin-gonic/gin"
)

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

type UserController struct {
	userService *service.UserService
}

func (uc *UserController) GetInfoUserById(c *gin.Context) {

}
