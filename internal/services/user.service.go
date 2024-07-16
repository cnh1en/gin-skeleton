package service

import (
	"github.com/cnh1en/lets_go/internal/repos"
	"github.com/cnh1en/lets_go/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUser(ctx *gin.Context) {
	response.ErrorResponse(ctx, 20003, "fuck")
}
