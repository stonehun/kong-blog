package controllers

import (
	"github.com/kataras/iris/v12"
	"kong_blog/models"
	"kong_blog/repositories"
	"kong_blog/services"
)

type UserController struct {
	UserService    services.UserService
	UserRepository repositories.UserRepository
	BaseController
}

func NewUserController() *UserController {
	return &UserController{
		UserService:    services.NewUserServices(),
		UserRepository: repositories.NewUserRepository(),
	}
}

//获取当前用户详情
func (c *UserController) GetDetail(ctx iris.Context) *models.Result {
	return models.NewResult(c.UserRepository.GetUserById(c.AuthUser(ctx).Id), 1)
}
