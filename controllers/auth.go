package controllers

import (
	"github.com/kataras/iris/v12"
	"kong_blog/models"
	"kong_blog/services"
)

type AuthController struct {
	UserService services.UserService
	BaseController
}

func NewAuthController() *AuthController {
	return &AuthController{
		UserService: services.NewUserServices(),
	}
}

//登录请求参数
type LoginParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

//登录返回参数
type LoginResponse struct {
	Token string `json:"token"`
}

//获取当前用户详情
func (c *AuthController) PostLogin(ctx iris.Context) *models.Result {
	var loginParams LoginParams
	err := ctx.ReadJSON(&loginParams)
	if err != nil {
		ctx.Application().Logger().Warn("解析错误:", err)
		return models.NewResult(err, 0)
	}

	ctx.Application().Logger().Debug("登录信息:", loginParams)

	//登录操作
	token, err := c.UserService.Login(loginParams.Name, loginParams.Password)
	if err != nil {
		ctx.Application().Logger().Debug("登录失败:", err)
		return models.NewResult(err, 0)
	}

	return models.NewResult(LoginResponse{Token: token}, 1)
}
